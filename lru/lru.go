package lru

import (
	"errors"
	dll "lru-cache/lru/linkedlist"
	"sync"
	"time"
)

type ValuePair struct {
	val interface{}
	ttl int64
}

type LRU struct {
	kvPair          map[int]ValuePair
	recencyRanking  *dll.Dll
	mu              *sync.Mutex
	currentCapacity int
	maxCapacity     int
}

func NewLru(capacity int) *LRU {
	return &LRU{
		kvPair:          make(map[int]ValuePair),
		mu:              &sync.Mutex{},
		recencyRanking:  dll.NewList(),
		currentCapacity: 0,
		maxCapacity:     capacity,
	}
}

//Integrate dll into lru

func (l *LRU) Get(key int) (*ValuePair, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	val, ok := l.kvPair[key]
	if !ok {
		return &ValuePair{}, errors.New("key does not exist")
	}
	if time.Now().Unix() > val.ttl {
		return &ValuePair{}, errors.New("key expired")
	}
	newVal := val
	l.kvPair[key] = newVal
	l.recencyRanking.MoveToTop(key)
	return &val, nil

}

func (l *LRU) Put(key int, val interface{}, ttl time.Time) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	limit := l.currentCapacity >= l.maxCapacity
	l.recencyRanking.AddNode(key, limit)
	l.kvPair[key] = ValuePair{val: val, ttl: ttl.Unix()}
	if !limit {
		l.currentCapacity++
	}
	return nil
}

func (l *LRU) Remove(key int) {
	l.mu.Lock()
	delete(l.kvPair, key)
	l.recencyRanking.Delete(key)
	l.mu.Unlock()

}

func (l *LRU) Clear() {
	l.mu.Lock()
	l.kvPair = make(map[int]ValuePair)
	l.mu.Unlock()
}

func (l *LRU) cleanup() {
	l.mu.Lock()
	for key, val := range l.kvPair {
		if time.Now().Unix() > val.ttl {
			delete(l.kvPair, key)
			l.recencyRanking.Delete(key)
		}
	}

	l.currentCapacity = len(l.kvPair)
	l.mu.Unlock()
}

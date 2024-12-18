package dll

/*
* * DLL: Doubly linked list
 */

type Dll struct {
	head *node
	tail *node
}

func NewList() *Dll {
	return &Dll{head: nil, tail: nil}
}

func (d *Dll) AddNode(key int, limit bool) {
	node := newNode(key)

	if d.head == nil && d.tail == nil {
		d.head = node
		d.tail = node
		return
	}

	node.next = d.head
	d.head.prev = node
	d.head = node

	if limit {
		if d.tail.prev != nil {
			d.tail = d.tail.prev
			d.tail.next = nil
		} else {
			d.tail = node
			d.head = node
			node.next = nil
		}
	}
}

func (d *Dll) MoveToTop(key int) {
	if d.head == nil || d.head.key == key {
		return
	}

	curr := d.head
	for curr != nil {
		if curr.key == key {
			// Remove node from current position
			if curr == d.tail {
				d.tail = curr.prev
				d.tail.next = nil
			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
			}

			// Move to head
			curr.next = d.head
			curr.prev = nil
			d.head.prev = curr
			d.head = curr
			return
		}
		curr = curr.next
	}
}

func (d *Dll) Delete(key int) {
	if d.head == nil {
		return
	}

	if d.head.key == key {
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		} else {
			d.tail = nil
		}
		return
	}

	curr := d.head
	for curr != nil {
		if curr.key == key {
			if curr == d.tail {
				d.tail = curr.prev
				d.tail.next = nil
			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
			}
			return
		}
		curr = curr.next
	}
}

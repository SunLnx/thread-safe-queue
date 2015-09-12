package queue

import "sync"

func NewQueue() *Queue {
	head := &Node{}
	return &Queue{
		head:   head,
		tail:   head,
		length: 0,
		lock:   &sync.Mutex{},
	}
}

type Queue struct {
	head   *Node
	tail   *Node
	length int

	lock *sync.Mutex
}

type Node struct {
	value interface{}
	forth *Node
	back  *Node
}

func (q *Queue) Enqueue(value interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	node := &Node{
		value: value,
	}
	q.tail.back = node
	node.forth = q.tail
	q.tail = node
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.head == q.tail {
		return nil
	}
	value := q.head.back.value

	q.head.back = q.head.back.back
	if q.head.back == nil {
		q.tail = q.head
	} else {
		q.head.back.forth = q.head
	}
	q.length--
	return value
}

func (q *Queue) Empty() bool {
	return q.head == q.tail
}

func (q *Queue) Clear() {
	var node, tmpNode *Node
	q.lock.Lock()
	defer q.lock.Unlock()
	for node = q.head.back; node != nil; {
		tmpNode = node.back
		node.forth = nil
		node.back = nil
		node = nil
		node = tmpNode
	}
	q.head.back = nil
	q.tail = q.head
	q.length = 0
}

func (q *Queue) Len() int {
	return q.length
}

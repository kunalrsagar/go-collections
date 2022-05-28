package collections

type Queue struct {
	firstNode *Node
	lastNode  *Node
	length    int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val interface{}) {
	newNode := NewNode(val)
	if q.length == 0 {
		q.firstNode = newNode
		q.lastNode = newNode
	} else {
		q.lastNode.Next = newNode
		q.lastNode = newNode
	}
	q.length++
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if q.length == 0 {
		return nil, false
	} else {
		first := q.firstNode

		if q.firstNode.Next != nil {
			next := q.firstNode.Next
			q.firstNode = next
		} else {
			q.firstNode = nil
		}

		q.length--
		return first.Value, true
	}
}

func (q *Queue) Peek() (interface{}, bool) {
	if q.length == 0 {
		return nil, false
	} else {
		return q.firstNode.Value, true
	}
}

package list

type node[T any] struct {
	value T
	prev *node[T]
	next *node[T]
}

type List[T any] struct {
	head *node[T]
	last *node[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Push(item T) {
	n := &node[T]{value: item}
	
	if l.head == nil {
		l.head = n
		l.last = n
		return
	}
	l.last.next = n
	n.prev = l.last
	l.last = n
}

func (l *List[T]) Each(fn func (T)) {
	for i := l.head; i != nil; i = i.next {
		fn(i.value)
	}
}

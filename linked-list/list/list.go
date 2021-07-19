package list

type List struct {
	Head *Node
}

func (l List) Size() int {
	size := 1
	current := l.Head
	for current.Next != nil {
		current = current.Next
		size++
	}
	return size
}

func (l List) Empty() bool {
	return l.Head == nil
}

func (l *List) ValueAt(index int) int {
	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l List) Back() int {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Value
}

func (l List) Front() int {
	return l.Head.Value
}

func (l List) Insert(index int, value int) {
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	next := current.Next
	current.Next =
		&Node{Value: value,
			Next: next,
		}
}

func (l List) Erase(index int) {
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
}

func (l *List) Reverses() {
	current := l.Head
	var prev *Node
	for current.Next != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	current.Next = prev
	l.Head = current
}

func (l *List) PopBack() {
	current := l.Head
	prev := &Node{}
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	prev.Next = nil
}

func (l *List) PushBack(value int) {
	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{
		Value: value,
		Next:  nil,
	}
}

func (l *List) PopFront() {
	if !l.Empty() {
		next := l.Head.Next
		l.Head = next
	}
}

func (l *List) RemoveValue(value int) {
	current := l.Head
	prev := &Node{}
	for current.Next != nil {
		if current.Value == value {
			break
		}
		prev = current
		current = current.Next
	}
	prev.Next = current.Next
}

func (l *List) PushFront(value int) {
	head := l.Head
	node := Node{
		Value: value,
		Next:  head,
	}
	l.Head = &node
}

func (l List) PrintAll() {
	current := l.Head
	println("//------------//")
	for current.Next != nil {
		println(current.Value)
		current = current.Next
	}
	println(current.Value)
}

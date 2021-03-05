package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	var el = ListItem{
		Value: v,
		prev:  nil,
	}

	if l.front != nil {
		l.front.prev = &el
		el.Next = l.front
	} else {
		l.back = &el
	}

	l.len++
	l.front = &el
	return &el
}

func (l *list) PushBack(v interface{}) *ListItem {
	var elem = ListItem{
		Value: v,
		Next:  nil,
	}

	if l.back != nil {
		l.back.Next = &elem
		elem.prev = l.back
	} else {
		l.front = &elem
	}

	l.len++
	l.back = &elem
	return &elem
}

func (l *list) Remove(i *ListItem) {
	l.len--
	if i.prev != nil {
		i.prev.Next = i.Next
	} else {
		l.front = i.Next
	}

	if i.Next != nil {
		i.Next.prev = i.prev
	} else {
		l.back = i.prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	if i.Next != nil {
		i.Next.prev = i.prev
	} else {
		l.back = i.prev
	}

	if i.prev != nil {
		i.prev.Next = i.Next
		i.Next = l.front
		l.front.prev = i
	}

	i.prev = nil
	l.front = i
}

func NewList() List {
	return &list{}
}

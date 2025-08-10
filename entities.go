package main

import "fmt"

type Snake struct {
	length int
	head   *Segment
}

func NewSnake(x int, y int, dirX int, dirY int) Snake {
	return Snake{
		length: 1,
		head:   NewSegment(x, y, dirX, dirY),
	}
}

func (s *Snake) Move() {
	if s == nil || s.head == nil {
		panic("snake/snake head is uninitialized")
	}

	for current := s.head; current != nil; current = current.next {
		current.x += current.dirX
		current.y += current.dirY
		if current.next != nil {
			current.dirX = current.next.dirX
			current.dirY = current.next.dirY
		}
	}
}

func (s *Snake) Grow() {
	tail := s.head.GetTail()
	newTail := NewSegment(
		tail.x-(tail.dirX*SegmentSize),
		tail.y-(tail.dirY*SegmentSize),
		tail.dirX,
		tail.dirY,
	)
	tail.Append(newTail)
	s.length += 1
}

type Entity interface {
	GetX() int
	GetY() int
}

type Segment struct {
	x      int
	y      int
	dirX   int
	dirY   int
	before *Segment
	next   *Segment
}

func NewSegment(x int, y int, dirX int, dirY int) *Segment {
	return &Segment{x, y, dirX, dirY, nil, nil}
}

func (s *Segment) Append(new *Segment) {
	if s == nil || new == nil {
		panic(fmt.Sprintf("unexpected nil pointer: appending %v to %v", new, s))
	}

	tail := s.GetTail()
	tail.next = new
	new.before = tail
}

func (s *Segment) GetTail() *Segment {
	current := s
	for current.next != nil {
		current = current.next
	}
	return current
}

func (snake *Snake) Occupies(x, y int) bool {
	for current := snake.head; current != nil; current = current.next {
		if current.x == x && current.y == y {
			return true
		}
	}
	return false
}

func (s *Segment) GetX() int {
	return s.x
}

func (s *Segment) GetY() int {
	return s.y
}

type Food struct {
	x int
	y int
}

func (f *Food) GetX() int {
	return f.x
}

func (f *Food) GetY() int {
	return f.y
}

type Border struct {
	x1, y1 int
	x2, y2 int
}

func NewBorder(x1, y1, x2, y2 int) *Border {
	return &Border{x1, y1, x2, y2}
}

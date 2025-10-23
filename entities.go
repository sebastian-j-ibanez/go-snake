package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Color of a segment entity.
var SegmentColor = rl.NewColor(0, 173, 216, 255)

// Color of a food entity.
var FoodColor = rl.NewColor(255, 117, 20, 255)

// Type of entity that can be drawn.
type Drawable interface {
	Draw()
}

type Entity interface {
	GetX() int
	GetY() int
}

// Snake
type Snake struct {
	length int
	head   *Segment
}

// Create a new snake.
func NewSnake(x int, y int, dirX int, dirY int) Snake {
	return Snake{
		length: 1,
		head:   NewSegment(x, y, dirX, dirY),
	}
}

// Draw a snake.
func (s *Snake) Draw() {
	for current := s.head; current != nil; current = current.next {
		current.Draw()
	}
}

// Increment the snakes position.
func (s *Snake) Move() {
	if s == nil || s.head == nil {
		panic("snake/snake head is uninitialized")
	}

	type direction struct{ x, y int }
	var oldDirections []direction
	for current := s.head; current != nil; current = current.next {
		oldDirections = append(oldDirections, direction{current.dirX, current.dirY})
	}

	i := 0
	for current := s.head; current != nil; current = current.next {
		current.x += current.dirX * SegmentSize
		current.y += current.dirY * SegmentSize

		if i > 0 {
			current.dirX = oldDirections[i-1].x
			current.dirY = oldDirections[i-1].y
		}
		i++
	}
}

// Append a new segment to the back of the snake.
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

// Segment
type Segment struct {
	x      int
	y      int
	dirX   int
	dirY   int
	before *Segment
	next   *Segment
}

func (s *Segment) Draw() {
	rl.DrawRectangle(
		int32(s.x),
		int32(s.y),
		SegmentSize,
		SegmentSize,
		SegmentColor,
	)
}

// Create new segment.
func NewSegment(x int, y int, dirX int, dirY int) *Segment {
	return &Segment{x, y, dirX, dirY, nil, nil}
}

// Append new segment.
func (s *Segment) Append(new *Segment) {
	if s == nil || new == nil {
		panic(fmt.Sprintf("unexpected nil pointer: appending %v to %v", new, s))
	}

	tail := s.GetTail()
	tail.next = new
	new.before = tail
}

// Get reference to tail segment.
func (s *Segment) GetTail() *Segment {
	current := s
	for current.next != nil {
		current = current.next
	}
	return current
}

// Check if snake is on a coordinate.
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

// Food
type Food struct {
	x int
	y int
}

// Draw a food entity.
func (f *Food) Draw() {
	if f != nil {
		rl.DrawRectangle(int32(f.x), int32(f.y), SegmentSize, SegmentSize, FoodColor)
	}
}

func (f *Food) GetX() int {
	return f.x
}

func (f *Food) GetY() int {
	return f.y
}

// Border
type Border struct {
	x1, y1 int
	x2, y2 int
}

// Create a new border.
func NewBorder(x1, y1, x2, y2 int) *Border {
	return &Border{x1, y1, x2, y2}
}

// Draw a border.
func (b *Border) Draw() {
	width := b.x2 - b.x1
	height := b.y2 - b.y1
	rectangle := rl.NewRectangle(float32(b.x1), float32(b.y1), float32(width), float32(height))
	rl.DrawRectangleLinesEx(rectangle, 1, rl.White)
}

// Represents the coordinates used to draw the score board.
type ScoreBoard struct {
	originX, originY int
	width, height int
	scoreX, scoreY int
}


// Draw the score board given a score value.
func (s *ScoreBoard) Draw(score int) {
	rectangle := rl.NewRectangle(float32(s.originX), float32(s.originY), float32(s.width), float32(s.height))
	rl.DrawRectangleLinesEx(rectangle, 3, rl.White)
	scoreText := fmt.Sprintf("Score: %d", score)
	textWidth := int(rl.MeasureText(scoreText, 40))
	centeredX := s.scoreX - textWidth/2
	rl.DrawText(scoreText, int32(centeredX), int32(s.scoreY), 40, SegmentColor);
}

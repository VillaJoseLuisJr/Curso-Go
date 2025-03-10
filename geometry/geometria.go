package geometry

type Triangle struct {
	size int
}

func (t *Triangle) doubleSize() {
	t.size *= 2
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s *Square) Area() float64 {
	return s.size * s.size
}

func (s *Square) Perimeter() float64 {
	return s.size * 4
}

func NewSquare(size float64) *Square {
	return &Square{size: size}
}

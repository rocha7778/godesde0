package interfaces

type Square struct {
	Width  float64
	Height float64
}

func (s *Square) SetWidth(with float64) {
	s.Width = with
}

func (s *Square) SetHeight(height float64) {
	s.Height = height
}

func (s *Square) GetArea() float64 {
	return s.Width * s.Height
}

func (s *Square) MuestraNombre() string {
	return "Cuadrado"
}

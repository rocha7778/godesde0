package interfaces

type Trinagulo struct {
	Base   float64
	Altura float64
}

func (t *Trinagulo) SetBase(base float64) {
	t.Base = base
}

func (t *Trinagulo) SetAltura(altura float64) {
	t.Altura = altura
}

func (t *Trinagulo) GetArea() float64 {
	return t.Base * t.Altura / 2
}
func (t *Trinagulo) MuestraNombre() string {
	return "Trinagulo"
}

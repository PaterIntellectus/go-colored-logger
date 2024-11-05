package ansi

type Ideogram Code

const (
	IUnderlined      Ideogram = 60
	IDoubleUnderline Ideogram = 61
	IOverline        Ideogram = 62
	IDoubleOverline  Ideogram = 63
	IStressMarking   Ideogram = 64
)

func (i Ideogram) String() string {
	return Code(i).String()
}

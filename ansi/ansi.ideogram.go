package ansi

type Ideogram Code

const (
	IdeogramUnderlined      Ideogram = 60
	IdeogramDoubleUnderline Ideogram = 61
	IdeogramOverline        Ideogram = 62
	IdeogramDoubleOverline  Ideogram = 63
	IdeogramStressMarking   Ideogram = 64
)

func (i Ideogram) String() string {
	return Code(i).String()
}

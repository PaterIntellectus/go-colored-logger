package ansi

type Reset Code

const (
	ResetAll Reset = 0

	ResetColor           Reset = 22
	ResetIntensity       Reset = 22
	ResetItalic          Reset = 23
	ResetFracture        Reset = 23
	ResetUnderline       Reset = 24
	ResetBlink           Reset = 25
	ResetInverse         Reset = 27
	ReseatConceal        Reset = 28
	ResetCrossedOut      Reset = 29
	ResetForegroundColor Reset = 39
	ResetBackgroundColor Reset = 49

	ResetFramed    Reset = 54
	ResetEncircled Reset = 54
	ResetOverlined Reset = 55

	ResetIdeogram Reset = 65
)

func (r Reset) String() string {
	return Code(r).String()
}

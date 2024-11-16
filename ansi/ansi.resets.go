package ansi

const (
	ResetAll Code = 0

	ResetColor           Code = 22
	ResetIntensity       Code = 22
	ResetItalic          Code = 23
	ResetFracture        Code = 23
	ResetUnderline       Code = 24
	ResetBlink           Code = 25
	ResetInverse         Code = 27
	ReseatConceal        Code = 28
	ResetCrossedOut      Code = 29
	ResetForegroundColor Code = 39
	ResetBackgroundColor Code = 49

	ResetFramed    Code = 54
	ResetEncircled Code = 54
	ResetOverlined Code = 55

	ResetIdeogram Code = 65
)

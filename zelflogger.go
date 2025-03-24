package zelflogger

import (
	"fmt"
)

type Style struct {}

func (*Style) Bold(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Bold, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) Dim(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Dim, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) Italic(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Italic, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) Underline(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Underline, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) BlinkSlow(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", BlinkSlow, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) BlinkFast(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", BlinkFast, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) Inverse(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Inverse, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) Hidden(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Hidden, fmt.Sprint(i...), Reset)
	return zxc
}

func (Style) Strikethrough(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Strikethrough, fmt.Sprint(i...), Reset)
	return zxc
}

type Color struct {}

func (Color) Black(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBlack, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Yellow(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgYellow, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Blue(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBlue, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Green(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgGreen, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Red(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgRed, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Magenta(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgMagenta, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Cyan(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgCyan, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) White(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgWhite, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightBlack(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightBlack, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightYellow(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightYellow, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightBlue(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightBlue, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightGreen(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightGreen, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightRed(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightRed, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightMagenta(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightMagenta, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightCyan(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightCyan, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) BrightWhite(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", FgBrightWhite, fmt.Sprint(i...), Reset)
	return zxc
}

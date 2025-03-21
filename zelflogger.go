package zelflogger

import (
	"fmt"
)

type Style struct

func (Style) Bold(i ...any) string {
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

type Color struct

func (Color) Black(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Black, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Yellow(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Yellow, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Blue(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Blue, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Green(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Green, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Red(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Red, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Magenta(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Magenta, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) Cyan(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Cyan, fmt.Sprint(i...), Reset)
	return zxc
}

func (Color) White(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", White, fmt.Sprint(i...), Reset)
	return zxc
}

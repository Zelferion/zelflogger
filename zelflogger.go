package zelflogger

type Style interface {
	Bold(...any) string
	Dim(...any) string
	Italic(...any) string
	Underline(...any) string
	BlinkSlow(...any) string
	BlinkFast(...any) string
	Inverse(...any) string
	Hidden(...any) string
	Strikethrough(...any) string
}

func Bold(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Bold, fmt.Sprint(i...), Reset)
	return zxc
}

func Dim(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Dim, fmt.Sprint(i...), Reset)
	return zxc
}

func Italic(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Italic, fmt.Sprint(i...), Reset)
	return zxc
}

func Underline(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Underline, fmt.Sprint(i...), Reset)
	return zxc
}

func BlinkSlow(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", BlinkSlow, fmt.Sprint(i...), Reset)
	return zxc
}

func BlinkFast(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", BlinkFast, fmt.Sprint(i...), Reset)
	return zxc
}

func Inverse(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Inverse, fmt.Sprint(i...), Reset)
	return zxc
}

func Hidden(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Hidden, fmt.Sprint(i...), Reset)
	return zxc
}

func Strikethrough(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Strikethrough, fmt.Sprint(i...), Reset)
	return zxc
}

type Color interface {
	Black(...any) string
	Yellow(...any) string
	Blue(...any) string
	Green(...any) string
	Red(...any) string
	Magenta(...any) string
	Cyan(...any) string
	White(...any) string
}

func ColorBlack(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Black, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorYellow(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Yellow, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorBlue(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Blue, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorGreen(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Green, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorRed(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Red, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorMagenta(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Magenta, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorCyan(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", Cyan, fmt.Sprint(i...), Reset)
	return zxc
}

func ColorWhite(i ...any) string {
	zxc := fmt.Sprintf("%s%v%s", White, fmt.Sprint(i...), Reset)
	return zxc
}

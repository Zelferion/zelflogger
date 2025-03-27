package zelflogger

import (
  "fmt"
)

const (
	// Foreground Colors
	FgBlack         = "\033[30m"
	FgRed           = "\033[31m"
	FgGreen         = "\033[32m"
	FgYellow        = "\033[33m"
	FgBlue          = "\033[34m"
	FgMagenta       = "\033[35m"
	FgCyan          = "\033[36m"
	FgWhite         = "\033[37m"

  FgBrightBlack   = "\033[38;5;244m"
  FgFatal         = "\033[38;5;52m" 

  FgBrightRed     = "\033[91m"
  FgBrightGreen   = "\033[92m"
  FgBrightYellow  = "\033[93m"
  FgBrightBlue    = "\033[94m"
  FgBrightMagenta = "\033[95m"
  FgBrightCyan    = "\033[96m"
  FgBrightWhite   = "\033[97m"

	// Background Colors
	BgBlack         = "\033[40m"
	BgRed           = "\033[41m"
	BgGreen         = "\033[42m"
	BgYellow        = "\033[43m"
	BgBlue          = "\033[44m"
	BgMagenta       = "\033[45m"
	BgCyan          = "\033[46m"
	BgWhite         = "\033[47m"

	// Text Styles
	Bold          = "\033[1m"
	Dim           = "\033[2m"
	Italic        = "\033[3m"
	Underline     = "\033[4m"
	BlinkSlow     = "\033[5m"
	BlinkFast     = "\033[6m"
	Inverse       = "\033[7m"
	Hidden        = "\033[8m"
	Strikethrough = "\033[9m"

	Reset         = "\033[0m"
)

var ColDebug string = fmt.Sprintf("\033[38;2;%d;%d;%dm", 68, 76, 120)

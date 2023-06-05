package progress_bar

import (
	"fmt"
	"strings"
)

type StyleList []Style

type Style int

// Base attributes
const (
	styleReset Style = iota
	StyleBold
	StyleFaint
	StyleItalic
	StyleUnderline
	StyleBlinkSlow
	StyleBlinkRapid
	StyleReverseVideo
	StyleConcealed
	StyleCrossedOut
)

// Foreground text colors
const (
	StyleFgBlack Style = iota + 30
	StyleFgRed
	StyleFgGreen
	StyleFgYellow
	StyleFgBlue
	StyleFgMagenta
	StyleFgCyan
	StyleFgWhite
)

// Foreground Hi-Intensity text colors
const (
	StyleFgHiBlack Style = iota + 90
	StyleFgHiRed
	StyleFgHiGreen
	StyleFgHiYellow
	StyleFgHiBlue
	StyleFgHiMagenta
	StyleFgHiCyan
	StyleFgHiWhite
)

// Background text colors
const (
	StyleBgBlack Style = iota + 40
	StyleBgRed
	StyleBgGreen
	StyleBgYellow
	StyleBgBlue
	StyleBgMagenta
	StyleBgCyan
	StyleBgWhite
)

// Background Hi-Intensity text colors
const (
	StyleBgHiBlack Style = iota + 100
	StyleBgHiRed
	StyleBgHiGreen
	StyleStyleBgHiYellow
	StyleBgHiBlue
	StyleBgHiMagenta
	StyleBgHiCyan
	StyleBgHiWhite
)

func (sl StyleList) String() string {
	var str []string
	for _, s := range sl {
		str = append(str, fmt.Sprintf("%s", s))
	}
	return strings.Join(str, "")
}

func (s Style) String() string {
	return fmt.Sprintf("\033[%dm", int(s))
}

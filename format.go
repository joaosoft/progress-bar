package progress_bar

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type format struct {
	*Format
	pb *ProgressBar
}

type Format struct {
	text      string
	styleList StyleList
	align     align
}

func (f *Format) Left() *Format {
	f.align = alignLeft
	return f
}

func (f *Format) Center() *Format {
	f.align = alignCenter
	return f
}

func (f *Format) Right() *Format {
	f.align = alignRight
	return f
}

func (f *Format) Style(style ...Style) *Format {
	f.styleList = append(f.styleList, style...)
	return f
}

// String print a format
func (f *format) String() string {
	var alignSpace float32

	switch f.align {
	case alignLeft:
	case alignCenter:
		alignSpace = ((f.pb.weight.max / 2) + float32(progressBarLeftSpace)) - ((float32(utf8.RuneCountInString(f.pb.title.text)) * f.pb.weight.unit) / 2)
	case alignRight:
		alignSpace = (f.pb.weight.max + float32(progressBarLeftSpace)) - float32(utf8.RuneCountInString(f.pb.title.text))*f.pb.weight.unit
	}

	return fmt.Sprintf("%s%s%s", strings.Repeat(" ", int(alignSpace)), f.styleList, f.text)
}

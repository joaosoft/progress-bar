package progress_bar

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Format
type format struct {
	text  string
	align Align
	pb    *ProgressBar
}

// String print a format
func (f *format) String() string {
	var alignSpace float32

	switch f.align {
	case AlignLeft:
	case AlignCenter:
		alignSpace = ((f.pb.weight.max / 2) + float32(progressBarLeftSpace)) - ((float32(utf8.RuneCountInString(f.pb.title.text)) * f.pb.weight.unit) / 2)
	case AlignRight:
		alignSpace = (f.pb.weight.max + float32(progressBarLeftSpace)) - float32(utf8.RuneCountInString(f.pb.title.text))*f.pb.weight.unit
	}

	return fmt.Sprintf("%s%s", strings.Repeat(" ", int(alignSpace)), f.text)
}

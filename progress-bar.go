package progress_bar

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

type ProgressBar struct {
	symbol rune
	writer io.Writer
	weight weight
	delay  time.Duration
	mux    *sync.Mutex

	progress     float32
	progressText string
	cleared      bool
	title        *format
	comment      *format
}

type weight struct {
	min  float32
	max  float32
	unit float32
}

// Sizes
const (
	minProgressBar = 0.0
	maxProgressBar = 100.0
)

// Progress bar corners
const (
	progressBarLeftText  = "[ "
	progressBarRightText = "]"
)

// Progress bar Corners size
var (
	progressBarLeftSpace = utf8.RuneCountInString(progressBarLeftText) + utf8.RuneCountInString(progressBarRightText)
)

// New creates a new progress bar
func New(options ...Option) *ProgressBar {
	progressBar := &ProgressBar{
		mux:      &sync.Mutex{},
		symbol:   '█',
		writer:   os.Stdout,
		progress: 0,
		weight: weight{
			min:  minProgressBar,
			max:  (100 / maxProgressBar) * maxProgressBar,
			unit: 100 / maxProgressBar,
		},
		delay: 200 * time.Millisecond,
	}

	progressBar.title = &format{
		align: AlignCenter,
		pb:    progressBar,
	}
	progressBar.comment = &format{
		align: AlignLeft,
		pb:    progressBar,
	}

	progressBar.Reconfigure(options...)

	return progressBar
}

// Title sets the title
func (pb *ProgressBar) Title(text string, align ...Align) {
	pb.title.text = text
	if len(align) > 0 {
		pb.title.align = align[0]
	}
}

// Comment sets a comment
func (pb *ProgressBar) Comment(text string, align ...Align) {
	pb.comment.text = text
	if len(align) > 0 {
		pb.comment.align = align[0]
	}
}

// Add adds a percentage
func (pb *ProgressBar) Add(value int) bool {
	pb.mux.Lock()
	defer pb.mux.Unlock()

	if value == 0 {
		return true
	}

	calcValue := float32(value) * pb.weight.unit
	pb.cleared = false

	switch calcValue > 0 {
	case true:
		if pb.progress < pb.weight.max {
			if calcValue > pb.weight.max {
				pb.progress = pb.weight.max
				pb.progressText += strings.Repeat(string(pb.symbol), int(pb.weight.max-pb.progress))
			} else {
				pb.progress += calcValue
				pb.progressText += strings.Repeat(string(pb.symbol), int(calcValue))
			}
		}
	case false:
		if pb.progress > pb.weight.min {
			if calcValue < pb.weight.min {
				pb.progress = pb.weight.min
				pb.progressText = ""
			} else {
				pb.progress = pb.progress - calcValue
				pb.progressText = pb.progressText[:utf8.RuneCountInString(pb.progressText)-int(calcValue)]
			}
		}
	}

	pb.reprint()

	return pb.progress == pb.weight.max
}

// clean cleans the console
func (pb *ProgressBar) clean() bool {
	if !pb.cleared {
		pb.cleared = true

		// :: clean
		// from bottom
		_, _ = pb.writer.Write([]byte("\x1bc"))

		// from top
		_, _ = pb.writer.Write([]byte("\x1b[2J"))
	}

	return pb.cleared
}

// print prints to the console
func (pb *ProgressBar) print() {
	// :: print title
	if pb.title.text != "" {
		_, _ = pb.writer.Write([]byte(fmt.Sprintf("%s%s%s%s\n", styleModeBold, styleColorGreen, pb.title, styleReset)))
	}

	// :: print progress bar
	_, _ = pb.writer.Write([]byte(fmt.Sprintf("%s%s %s%s", progressBarLeftText, pb.progressText, strings.Repeat(" ", int(pb.weight.max-pb.progress)), progressBarRightText)))

	// :: print comment
	if pb.comment.text != "" {
		_, _ = pb.writer.Write([]byte(fmt.Sprintf("\n%s\n", pb.comment)))
	}
}

// wait waits the defined delay
func (pb *ProgressBar) wait() {
	<-time.After(pb.delay)
}

// reprint reprints the console
func (pb *ProgressBar) reprint() {
	pb.clean()
	pb.print()
	pb.wait()
}

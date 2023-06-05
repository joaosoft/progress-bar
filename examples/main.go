package main

import (
	pb "github.com/joaosoft/progress-bar"
	"time"
)

func main() {
	progressBar := pb.New(pb.WithDelay(time.Millisecond * 500))
	progressBar.Title("Installation").Style(pb.StyleBold, pb.StyleBgHiBlue)

	progressBar.Comment("text 1").Right().Style(pb.StyleFgHiWhite)
	progressBar.Add(20)

	progressBar.Comment("text 2").Center()
	progressBar.Add(20)

	progressBar.Comment("text 3").Left()
	progressBar.Add(20)

	progressBar.Comment("text 4")
	progressBar.Add(20)
	progressBar.Add(-20)
	progressBar.Add(20)

	progressBar.Comment("text 5")
	progressBar.Add(20)
}

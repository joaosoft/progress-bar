package main

import (
	pb "github.com/joaosoft/progress-bar"
)

func main() {
	progressBar := pb.New()
	progressBar.Title("Installation")

	progressBar.Comment("text 1", pb.AlignRight)
	progressBar.Add(20)

	progressBar.Comment("text 2", pb.AlignCenter)
	progressBar.Add(20)

	progressBar.Comment("text 3", pb.AlignLeft)
	progressBar.Add(20)

	progressBar.Comment("text 4")
	progressBar.Add(20)

	progressBar.Comment("text 5")
	progressBar.Add(20)
}

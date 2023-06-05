progress-bar
================

[![Build Status](https://travis-ci.org/joaosoft/progress-bar.svg?branch=master)](https://travis-ci.org/joaosoft/progress-bar) | [![codecov](https://codecov.io/gh/joaosoft/progress-bar/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/progress-bar) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/progress-bar)](https://goreportcard.com/report/github.com/joaosoft/progress-bar) | [![GoDoc](https://godoc.org/github.com/joaosoft/progress-bar?status.svg)](https://godoc.org/github.com/joaosoft/progress-bar)

A simple progress bar implementation.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* title
* comment
* symbol
* writer
* delay

>### Go
```
go get github.com/joaosoft/progress-bar
```

## Usage 
This examples are available in the project at [progress-bar/examples](https://github.com/joaosoft/progress-bar/tree/master/examples)

### Code
```go
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
```

> ##### Result:
```
                                               Installation
[ ████████████████████████████████████████████████████████████████████████████████████████████████████ ]
text 5
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com

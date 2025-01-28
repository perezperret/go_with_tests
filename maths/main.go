package main

import (
	"os"
	"time"

	"maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}

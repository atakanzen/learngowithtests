package main

import (
	"learngowithtests/maths/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}

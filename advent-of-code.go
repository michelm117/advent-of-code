package main

import (
	"flag"
	"fmt"

	"github.com/michelm117/advent-of-code/day_1"
)

func main() {
	filePath := flag.String("filePath", "foo", "path to the input file")
	flag.Parse()

	calibrationSum := day_1.Calibrate(*filePath)
	fmt.Println("Calibration sum: " + fmt.Sprint(calibrationSum))
}

package main

import (
	"flag"
	"fmt"

	"github.com/michelm117/advent-of-code/day_1"
	"github.com/michelm117/advent-of-code/day_2"
	"github.com/michelm117/advent-of-code/day_3"
	"github.com/michelm117/advent-of-code/day_4"
)

func main() {
	day := flag.String("day", "1", "the day to run")
	filePath := flag.String("filePath", "./day_1/input.txt", "path to the input file")

	flag.Parse()

	switch *day {
	case "1":
		calibrationSum := day_1.Calibrate(*filePath)
		fmt.Println("Calibration sum: " + fmt.Sprint(calibrationSum))

	case "2":
		sumOfPossibleGames, sumOfPower := day_2.Solve(*filePath)
		fmt.Println("Sum of possible games: " + fmt.Sprint(sumOfPossibleGames))
		fmt.Println("Sum of power: " + fmt.Sprint(sumOfPower))

	case "3":
		notPart, gearRatio := day_3.Solve(*filePath)
		fmt.Println("Sum of parts:", fmt.Sprint(notPart))
		fmt.Println("gearRatio: ", fmt.Sprint(gearRatio))

	case "4":
		solOne, solTwo := day_4.GetScratchpadPoints(*filePath)
		fmt.Println("Solution One:", fmt.Sprint(solOne))
		fmt.Println("Solution Two:", fmt.Sprint(solTwo))

	default:
		fmt.Println("Day not found")
	}

}

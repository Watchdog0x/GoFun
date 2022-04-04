package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		sum          float64
		product      float64 = 1
		wrongagument []string
	)

	if args := os.Args[1:]; len(args) <= 10 && len(args) >= 2 {

		for i, v := range args {
			args[i] = strings.Replace(v, ",", ".", 1)
		}

		for _, v := range args {
			if arg, err := strconv.ParseFloat(v, 64); err != nil {
				wrongagument = append(wrongagument, v)
				continue

			} else {
				sum += arg
				product *= arg
			}
		}
	} else {
		fmt.Println("The program needs between 2 and 10 arguments")
	}

	if sum != 0 {
		fmt.Printf("Sum: %.2f, Product: %.2f\n", sum, product)
	}
	if wrongagument != nil {
		fmt.Printf("This arguments was skipped %v\n", wrongagument)
	}
}

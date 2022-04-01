package cli

import (
	"flag"
	"fmt"
)

func Execute() error {
	countFlag := flag.Int("count", 10, "maximum number of items to output")
	inputFlag := flag.String("input", "words", "input [words,lines,item]")
	outputFlag := flag.String("output", "words", "output [words,lines]")
	persistenceFlag := flag.Int("persistence", 5, "how hard to try to produce acceptable output")
	similarityFlag := flag.Int("similarity", 3, "degree of similarity between input and output items")

	fmt.Println("input", *inputFlag)
	fmt.Println("count", *countFlag)
	fmt.Println("similarity", *similarityFlag)
	fmt.Println("persistence", *persistenceFlag)
	fmt.Println("output", *outputFlag)
	flag.Parse()
	return nil
}

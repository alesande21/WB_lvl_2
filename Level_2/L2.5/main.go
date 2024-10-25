package main

import (
	"dictionary/internal/dictionary"
	"fmt"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	dict := dictionary.NewDictionary()

	m := dict.CreateDictionary(words)

	for key, value := range *m {
		fmt.Printf("{Key: %v Value: ", key)
		for _, word := range value {
			fmt.Printf("%s ", word)
		}
		fmt.Printf("}\n")
	}
}

package dictionary

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Dictionary struct {
	dic map[[33]int32][]string
}

func NewDictionary() *Dictionary {
	return &Dictionary{dic: make(map[[33]int32][]string)}
}

func (d *Dictionary) CreateDictionary(str []string) *map[[33]int32][]string {

	for _, word := range str {
		key, lowerCase, err := d.getKeyForWord(word)
		if err != nil {
			fmt.Printf("CreateDictionary-> d.getKeyForWord%s", err)
			continue
		}

		multitude, ok := d.dic[key]

		if !ok {
			multitude = append(multitude, lowerCase)
			d.dic[key] = multitude
		} else if ok = d.CheckMultitude(multitude, word); ok {
			multitude = append(multitude, lowerCase)
			d.dic[key] = multitude
		}

	}
	return &d.dic
}

func (d *Dictionary) CheckMultitude(multitude []string, word string) bool {

}

func (d *Dictionary) getKeyForWord(w string) ([33]int32, string, error) {
	var key [33]int32
	var b strings.Builder
	for i := 0; i < len(w); {

		runeValue, width := utf8.DecodeRuneInString(w[i:])

		if runeValue >= 127 && runeValue < 160 {
			key[runeValue-128] += 1
			b.WriteRune(runeValue + 32)
		} else if runeValue == 240 {
			key[runeValue-208] += 1
			b.WriteRune(runeValue + 1)
		} else if runeValue >= 160 && runeValue < 176 {
			key[runeValue-160] += 1
			b.WriteRune(runeValue)
		} else if runeValue >= 224 && runeValue < 240 {
			key[runeValue-208] += 1
			b.WriteRune(runeValue)
		} else if runeValue == 241 {
			key[runeValue-208] += 1
			b.WriteRune(runeValue)
		} else {
			return key, "", fmt.Errorf(": в строке присуствую не русские символы: %s", w)
		}

		i += width
	}

	return key, "", nil
}

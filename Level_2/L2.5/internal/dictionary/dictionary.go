package dictionary

import (
	"fmt"
	"sort"
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
		if utf8.RuneCountInString(word) < 2 {
			continue
		}
		key, lowerCase, err := d.getKeyForWord(word)
		if err != nil {
			fmt.Printf("CreateDictionary-> d.getKeyForWord%s", err)
			continue
		}

		multitude, ok := d.dic[key]
		if !ok || !d.checkMultitude(multitude, word) {
			multitude = append(multitude, lowerCase)
			d.dic[key] = multitude
		}

	}

	for key, words := range d.dic {
		if len(words) <= 1 {
			delete(d.dic, key)
		} else {
			sort.Strings(words)
		}
	}
	return &d.dic
}

func (d *Dictionary) checkMultitude(multitude []string, word string) bool {
	for _, m := range multitude {
		if strings.Compare(m, word) == 0 {
			return true
		}
	}
	return false
}

func (d *Dictionary) getKeyForWord(w string) ([33]int32, string, error) {
	var key [33]int32
	var b strings.Builder
	for i := 0; i < len(w); {

		runeValue, width := utf8.DecodeRuneInString(w[i:])
		if runeValue >= 1040 && runeValue < 1072 {
			key[runeValue-1040] += 1
			b.WriteRune(runeValue + 32)
		} else if runeValue == 1025 || runeValue == 1105 {
			key[32] += 1
			b.WriteRune(1105)
		} else if runeValue >= 1072 && runeValue < 1104 {
			key[runeValue-1072] += 1
			b.WriteRune(runeValue)
		} else {
			return key, "", fmt.Errorf(": в строке присуствую не русские символы/код руны: %s/%d\n", w, runeValue)
		}

		i += width
	}

	return key, b.String(), nil
}

/*
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
			return key, "", fmt.Errorf(": в строке присуствую не русские символы/код руны: %s/%d\n", w, runeValue)
		}

		i += width
	}

	return key, "", nil
}
*/

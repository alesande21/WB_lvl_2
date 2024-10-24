package main

import (
	"fmt"
)

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}
func main() {

	//str := "пятак"
	//str2 := "пятка"

	var i int64

	i |= 1 << 11
	i |= 1 << 0
	fmt.Printf("%b\n", i)
	//fmt.Printf("%v\n", binary(str2))

}

package words

import (
	"fmt"
	"strings"
)

//func main() {
//	split("xx")
//}

func Split(t string) map[string]int {
	//Jowแก
	m := make(map[string]int)
	someString := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	someString = strings.Replace(someString, ",", "", -1)
	someString = strings.Replace(someString, ".", "", -1)
	words := strings.Fields(someString)

	for _, v := range words {
		//fmt.Println(i, v)
		if m[v] == 0 {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}

	for i, v := range m {
		fmt.Println(i, v)
	}

	return m
}

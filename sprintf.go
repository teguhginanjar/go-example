package main

import "fmt"

func main() {
	fmt.Printf("%.09d\n", 45)
	str := fmt.Sprintf("%.09d", 452)
	fmt.Println(str)
	b := make([]byte, 9)

	for i := 0; i < len(str); i++ {
		b[i] = str[i]
	}
	fmt.Printf("%s\n", b)
}

package main

import "fmt"

func CommonString(a, b string) string {

	solution := ""
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			for j := 0; j < len(b); j++ {
				if a[i+j] != b[j] {
					count = 0
					break
				} else {
					count++
					solution += string(a[i+j])
				}
			}
			if count == len(b) {
				break
			} else {
				solution = ""
			}
		}
	}
	return solution
}

func main() {

	fmt.Println(CommonString("hello", "hell"))

}

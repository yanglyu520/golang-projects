// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	i, j := 10, 20
	p1 := &i
	p2 := &j
	fmt.Println(i, j, p1, p2)

	//change the value of i and j
	*p1 = 30
	*p2 = 40
	fmt.Println(i, j, p1, p2)
}

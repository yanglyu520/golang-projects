
package main

import "fmt"

func main() {
	fmt.Println("--------map---------")
	// create a map
	map_a := make(map[int]string)
	//add value
	map_a[1] = "Yang"
	map_a[2] = "steve"
	//print the map
	fmt.Println("map: ", map_a)
	//delete a key value pair
	delete(map_a, 1)
	//find out the size of the map
	size1 := len(map_a)
	fmt.Println("size: ", size1)

	fmt.Println("--------set---------")
	//create a set
	set1 := make(map[int]bool)
	//add value
	set1[1] = true
	set1[3] = true
	//print the set
	for k := range set1 {
		fmt.Println(k)
	}
	//delete an item from the set
	delete(set1, 3)
	//find if the item still exists
	exist1 := set1[3]
	fmt.Println(exist1)
	//find out the size of the set
	size2 := len(set1)
	fmt.Println("size: ", size2)

	fmt.Println("--------set with void struct---------")
	type void struct{}
	var member void

	//create a set
	set2 := make(map[int]void)
	set2[1] = member
	set2[5] = member
	set2[9] = member
	//print
	for k := range set2 {
		fmt.Println(k)
	}
	//delete an item from the set
	delete(set2, 5)
	//find if an item still exists
	exist2 := set2[5]
	fmt.Println(exist2)
	//find out the size of the set
	size3 := len(set2)
	fmt.Println("size: ", size3)
}

package main

import "fmt"

func Color(colorName string) string {
	switch colorName {
	case "blue":
		return "#0000FF"
	case "white":
		return "#FFFFFF"
	case "grey":
		return "#888888"
	default:
		return "000000"
	}
}

func main(){
	colorName := "blue"
	fmt.Println(Color(colorName))
}
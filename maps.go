package main

import "fmt"

func main_maps() {

	// V1
	// colors := map[string]string{
	// 	"red":  "#ff0000",
	// 	"blue": "#abcabc",
	// }

	// V2
	// var colors map[string]string

	// V3
	colors := make(map[string]string)

	// add new item:
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	colors["green"] = "#faaaaf"
	// delete
	delete(colors, "green")

	fmt.Println(colors)

	printMap(colors)

}

func printMap(m map[string]string) {
	for key, v := range m {
		fmt.Println("key: ", key, " value: ", v)
	}
}

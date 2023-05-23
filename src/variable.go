package main

import "fmt"

func main() {
	var a int = 18
	var b string = "Hyun"
	var c, d, e int = 1, 2, 3
	var f float32 = 172.1
	var g bool = true
	var auto = 12.34
	fmt.Println(a, b, c, d, e, f, g, auto)

	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)
	height = 250
	weight = 350.56
	isRunning = true
	fmt.Println("name :", name, "\nheight :", height, "\nweight :", weight, "\nisRunning :", isRunning)

	short := "short variable"
	fmt.Println("short :", short)
}

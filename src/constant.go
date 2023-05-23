package main

import "fmt"

func main() {
	const str string = "constant"
	const auto = true

	fmt.Println("str :", str, "\nauto :", auto)

	const a, b int = 1, 2
	const c, d, e = true, "string", 3.4
	const (
		f    int = 56
		g, h     = "data", 7.89
	)
	fmt.Println("a :", a, "\nb :", b, "\nc :", c, "\nd :", d, "\ne :", e, "\nf :", f, "\ng :", g, "\nh :", h)
}

package main

import (
	"fmt"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	str := "one,two,,four"
	in := make(chan string)
	go func() {
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
		close(in)
	}()

	for word := range in {
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
}
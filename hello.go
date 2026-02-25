package main

import (
	"fmt"
	"github.com/fatih/color"
	"hexlet-go/greeting"
)

func main() {
	fmt.Println(greeting.Get())
	color.Red("We have red")
	color.Magenta("And many others ..")
}

package main

import ( . "fmt"; "time" )

func main() {
	go func(strings ...string) {
		for _, str := range strings {
			Println(str)
		}
	}("You", "are")

	go func(strings ...string) {
		for _, str := range strings {
			Println(str)
		}
	}("not")

	go func(strings ...string) {
		for _, str := range strings {
			Println(str)
		}
	}("seeing", "these", "strings")

	go func(strings ...string) {
		for _, str := range strings {
			Println(str)
		}
	}("in", "the")

	go func(strings ...string) {
		for _, str := range strings {
			Println(str)
		}
	}("proper")

	go func(strings ...string) {
		for _, str := range strings {
			Println(str)
		}
	}("order.")

	time.Sleep(time.Second)
}

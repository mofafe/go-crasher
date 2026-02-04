package main

import "log"

func main() {
	a()
}

func a() {
	go a()
	go a()
	for {
		log.Println(1 + 1)
	}
}

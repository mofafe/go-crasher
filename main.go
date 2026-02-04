package main

import "log"

func main() {
	a()
}

func a() {
	go b()
	go b()
	for {
		log.Println(1 + 1)
	}
}

func b() {
	go a()
	go a()
	for {
		log.Println(1 + 1)
	}
}

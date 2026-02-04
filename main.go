package main

import "log"

func main() {
	go main()
	go main()
	for {
		log.Println(1 + 1)
	}
}

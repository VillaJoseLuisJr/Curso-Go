package main

import (
	"log"
)

func main() {
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
	//log.Panic("Hey, I'm an error log!")
	//fmt.Print("Can you see me?")
}

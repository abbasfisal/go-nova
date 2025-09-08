package main

import (
	"fmt"
	"log"
	"time"
	"tut/config"
)

func main() {

	cnf := config.NewConfig()

	log.Println("env:", cnf.App.Name)
	go func() {
		fmt.Println("go routine 1")
	}()
	//hi
	time.Sleep(1 * time.Second)
	fmt.Println("hiHI")
	fmt.Println("Hello World")

}

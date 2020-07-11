package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// defer: executes function call after final statement of parent function, but before return
	// last in, first out (lifo)
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	// a more realistic example
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// defer takes value of function arguments when it was called (eager evaluation)

	// panic: similar to throwing an error in javascript, user-defined error
	// panic happens after deferred statements: function statements -> deferred statements -> panic -> return
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		// handle panic
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err) // re-throw panic if unable to handle
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

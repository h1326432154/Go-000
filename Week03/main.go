package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello,GopherCon SG")
	})
	go func() {
		if err := http.ListenAndServe("8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}

// func main() {

// 	// Capture starting number of goroutines.
// 	startingGs := runtime.NumGoroutine()

// 	leak()

// 	// Hold the program from terminating for 1 second to see
// 	// if any goroutines created by leak terminate.
// 	time.Sleep(time.Second)

// 	// Capture ending number of goroutines.
// 	endingGs := runtime.NumGoroutine()

// 	// Report the results.
// 	fmt.Println("========================================")
// 	fmt.Println("Number of goroutines before:", startingGs)
// 	fmt.Println("Number of goroutines after :", endingGs)
// 	fmt.Println("Number of goroutines leaked:", endingGs-startingGs)
// }

// // leak is a buggy function. It launches a goroutine that
// // blocks receiving from a channel. Nothing will ever be
// // sent on that channel and the channel is never closed so
// // that goroutine will be blocked forever.
// func leak() {
// 	ch := make(chan int)

// 	go func() {
// 		val := <-ch
// 		fmt.Println("We received a value:", val)
// 	}()
// }

var a string

/*
// --- go关键字开启新的goroutine，先行发生于这个goroutine开始执行
func f() {
	fmt.Println(a, "111")
}

func main() {
	a = "hello, world"
	go f()
}
*/

func hello() {
	go func() { a = "hello" }()
	print(a)
}

package main

import (
	"fmt"
	"sync/atomic"
)

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "hello,GopherCon SG")
// 	})
// 	go func() {
// 		if err := http.ListenAndServe("8080", nil); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	select {}
// }

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

// ### memory model
/*
var a string

// --- go关键字开启新的goroutine，先行发生于这个goroutine开始执行
func f() {
	fmt.Println(a, "111")
}

func main() {
	a = "hello, world"
	go f()
}
// */

/*
func hello() {
	go func() { a = "hello" }()
	print(a)
}

func main() {
	a = "hello, world"
	go hello()
}
*/

/*
var c = make(chan int, 10)

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
*/

/*
var c = make(chan int)

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
*/

/*
var l sync.RWMutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
*/

/*
var a string
var once sync.Once

func setup() {
	println("aaa")
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	once.Do(setup)
	println(a)
}

func main() {
	go doprint()
	go doprint()
	time.Sleep(time.Millisecond)
}
*/

/*
var a, b int

func f() {
	a = 1
	b = 2

}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}

*/

/* 错误示范
var a string
var done bool
var once sync.Once

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func main() {
	go doprint()
	go doprint()
}

*/

/*
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}
func main() {
	go setup()
	for !done {
	}
	print(a)
}
*/

/*
// T .
type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}
func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
*/

/*
func main() {
	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(2)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}
*/

/*
func main() {
	ch := make(chan struct{})
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- struct{}{}
		}(i)
		<-ch
	}
}
*/

func main() {
	var count uint32

	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
		}
	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {})

}

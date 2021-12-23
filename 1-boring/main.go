package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (s *Student) getName() string {
	return s.name
}

func (s *Student) hello(msg string) string {
	return fmt.Sprintf("%s helloworld %s", s.name, msg)
}

func main() {
	cstZone := time.FixedZone("", 0)
	fmt.Println(time.Unix(1671, 0).In(cstZone).Format("15:04:05"))
	fmt.Println(time.Now().Format("15:04:05"))
	// after run this line, the main goroutine is finished.
	// main goroutine is a caller. It doesn't wait for func boring finished
	// Thus, we don't see anything
	//go boring("boring!") // spawn a goroutine. (1)

	// To solve it, we can make the main go routine run forever by `for {}` statement.

	// for {
	// }

	// A little more interesting is the main goroutine exit. the program also exited
	// This code hang
	//fmt.Println("I'm listening")
	//time.Sleep(2 * time.Second)
	//fmt.Println("You're boring. I'm leaving")
	//
	//var stu Person = &Student{
	//	name: "aa",
	//}
	//s := stu.(*Student)
	//fmt.Println(s.getName())

	// However, the main goroutine and boring goroutine does not communicate each other.
	// Thus, the above code is cheated because the boring goroutine prints to stdout by its own function.
	// the line `boring! 1` that we see on terminal is the output from boring goroutine.

	// real conversation requires a communication
}

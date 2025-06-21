package main

import (
	"fmt"
	"time"
)

func main() {
	// if use go it mean we want to work parallel

	//dones := make([]chan bool, 4)
	done := make(chan bool) // แปลว่า ทุก func ใช้ channel เดียวกัน

	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] = make(chan bool)
	go greet("How are you?", done)
	//dones[2] = make(chan bool)
	go slowGreet("How...are...you...?", done)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	// todo ต้องรอค่าจำนวนเท่ากับจำนวน goroutine หากต้องการให้ทั้งหมดเสร็จสิ้น
	//<-dones // แต่ละ channel รับได้ค่าเดียว
	//<-dones
	//<-dones
	//<-dones
	//
	//for _, done := range dones {
	//	<-done
	//}

	//for doneChane := range done { // doneChane type = bool
	//	fmt.Println(doneChane)
	//}

	for range done { // loop chan (buildin) ... WTF!!

	}
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // todo close chan ที่ func ที่ทำงานหลังสุด (ถ้าเรารู้ ถ้าไม่รู้ก็ defer)
}

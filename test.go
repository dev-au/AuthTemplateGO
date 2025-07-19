package main

//
//import (
//	"fmt"
//	"time"
//)
//
//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	if time.Now().Nanosecond()%2 != 0 {
//		time.Sleep(time.Second * 3)
//	}
//	c <- sum
//
//}
//
//func main() {
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(s[:4], c)
//	go sum(s[4:], c)
//	x := <-c
//	y := <-c
//
//	fmt.Println(x, y)
//}

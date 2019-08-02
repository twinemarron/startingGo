package main

import (
  "fmt"
)

func main() {
  ch1 := make(chan int)
  ch2 := make(chan int)
  ch3 := make(chan int)
  go func() {
    for {
      i := <-ch1
      ch2 <- (i * 2)
    }
  }()
  go func() {
    for {
      i := <-ch2
      ch3 <- (i - 1)
    }
  }()
  n := 1
  LOOP:
  for {
    select {
    case ch1 <- 1:
      n++
    case i := <-ch3:
      fmt.Println("recieved", i)
    default:
      if n > 100 {
        break LOOP
      }
    }
  }
  // ch1 := make(chan int, 1)
  // ch2 := make(chan int, 1)
  // ch3 := make(chan int, 1)
  // ch1 <- 1
  // ch2 <- 1
  //
  // select {
  // case <- ch1:
  //   fmt.Println("ch1から受信")
  // case <- ch2:
  //   fmt.Println("ch2から受信")
  // case ch3 <- 3:
  //   fmt.Println("ch3から受信")
  // default:
  //   fmt.Println("ここへは到達しない")
  // }
}

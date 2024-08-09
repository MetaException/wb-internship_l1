package main

import (
	"fmt"
	"time"
)

func main() {
	var duration uint
	fmt.Scanln(&duration)

	sleep(duration)
}

func sleep(duration uint) {
	<-time.After(time.Second * time.Duration(duration))
}

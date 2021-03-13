package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	d := 1
	if len(os.Args) > 1 {
		d, _ = strconv.Atoi(os.Args[1])
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%d percent process complete\n", i)
		time.Sleep(time.Duration(d) * time.Second)
	}

	return
}

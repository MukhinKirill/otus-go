package main

import (
	"fmt"
	"log"
	"os"

	ntp "github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}
	fmt.Println(time)
	os.Exit(0)
}

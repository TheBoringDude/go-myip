package main

import (
	"fmt"
	"log"

	"github.com/TheBoringDude/go-myip"
)

func main() {
	ip, err := myip.GetMyIP()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ip.IP)
}

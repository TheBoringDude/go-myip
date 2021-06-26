package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/TheBoringDude/go-myip"
)

var fullTemplate = `Your IP Information
------------------------------------
  IP: {{.IP}}
  Geolocation:
     Latitude: {{.Geo.Latitude}}
     Longitude: {{.Geo.Longitude}}
  Country: {{.Country}}
  City: {{.City}}
  Continent: {{.Continent}}
  PostalCode: {{.PostalCode}}
  Region: {{.Region}}
  RegionCode: {{.RegionCode}}
  Timezone: {{.Timezone}}
`

// prints help message
func showHelp() {
	fmt.Println(` MyIP

	Example: myip --full --retries=3
  
	Flags:
		--retries, int
			 the number of retries after a failed request
		--full
			 show all fields
		--help, -h, help
			 show this help message
	`)
	os.Exit(0)
}

func main() {
	retries := flag.Int("retries", 0, "how many times to retry after failed request")
	full := flag.Bool("full", false, "print all items in result")
	help := flag.Bool("help", false, "show the help message")
	flag.Parse()

	// show help message
	if *help {
		showHelp()
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "help" {
			showHelp()
		}
	}

	// request
	ip, err := myip.GetMyIPWithRetry(*retries, 0)
	if err != nil {
		log.Fatal(err)
	}

	// print if full
	if *full {
		t, err := template.New("full").Parse(fullTemplate)
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(os.Stdout, ip)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	fmt.Println(ip.IP)
}

// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"shodan/shodan"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: main <searchterm>, page number")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)
	
	a2,err:=strconv.Atoi(os.Args[2])//takes the second command line arguement and turns it into an int for the use in host search
	if err!= nil {
		log.Panicln(err)
	}
	hostSearch, err := s.HostSearch(os.Args[1], a2) // changed to accept second arguement for page numbers
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ",host.IPString,"====")
		h,_ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ",host.IPString,"====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}


	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}

	var newPage string
	fmt.Printf("New Page?(Y for yes)\n")
	fmt.Scanln(&newPage)

	for newPage == "Y"{
		fmt.Printf(
			"Query Credits: %d\nScan Credits:  %d\n\n",
			info.QueryCredits,
			info.ScanCredits)
		
		var newPageNum string
		fmt.Printf("Please select which page you wish to choose.\n")
		fmt.Scanln(&newPageNum)
		a2,err:=strconv.Atoi(newPageNum)//takes the input and uses it as the next page that wants to be looked at
		if err!= nil {
			log.Panicln(err)
		}
		hostSearch, err := s.HostSearch(os.Args[1], a2) // puts the argument in here
		if err != nil {
			log.Panicln(err)
		}
	
		fmt.Printf("Host Data Dump\n")
		for _, host := range hostSearch.Matches {
			fmt.Println("==== start ",host.IPString,"====")
			h,_ := json.Marshal(host)
			fmt.Println(string(h))
			fmt.Println("==== end ",host.IPString,"====")
			//fmt.Println("Press the Enter Key to continue.")
			//fmt.Scanln()
		}
	
	
		fmt.Printf("IP, Port\n")
	
		for _, host := range hostSearch.Matches {
			fmt.Printf("%s, %d\n", host.IPString, host.Port)
		}
		
		fmt.Printf("New Page?(Y for yes)\n")
		fmt.Scanln(&newPage)
	}



}
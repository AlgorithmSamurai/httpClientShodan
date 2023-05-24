package main

import (
	"fmt"
	"github.com/AlgorithmSamurai/httpClientShodan/shodan/host"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKeyPath := "../../.env"
	err := godotenv.Load(apiKeyPath)
	if err != nil {
		log.Fatalf("Error loading %s file: %s", apiKeyPath, err)
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	if apiKey == "" {
		log.Fatalln("SHODAN_API_KEY environment variable not set")
	}
	s := host.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}

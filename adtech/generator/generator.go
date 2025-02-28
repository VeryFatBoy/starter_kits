package main

import (
	"fmt"
	"io"
	"os"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {

	if len(os.Args) != 2 {
		printUsage()
	}

	var maxCampaigns int

	switch os.Args[1] {
	case "small":
		maxCampaigns = 100
	case "medium":
		maxCampaigns = 1_000
	case "large":
		maxCampaigns = 10_000
	default:
		printUsage()
	}

	file, err := os.Create("campaigns-" + os.Args[1] + ".tsv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= maxCampaigns; i++ {
		s := fmt.Sprintf("%v\t%v\t%v\t1\n", i, gofakeit.BS(), gofakeit.Date())

		_, err = io.WriteString(file, s)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	file.Close()

	var maxEvents int

	switch os.Args[1] {
	case "small":
		maxEvents = 100_000
	case "medium":
		maxEvents = 1_000_000
	case "large":
		maxEvents = 10_000_000
	default:
		printUsage()
	}

	file, err = os.Create("events-" + os.Args[1] + ".tsv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < maxEvents; i++ {
		s := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
			gofakeit.Uint32(),
			gofakeit.RandomString([]string{"Impression", "Click", "Downstream Conversion"}),
			gofakeit.Company(),
			gofakeit.Number(1, maxCampaigns),
			gofakeit.Gender(),
			gofakeit.RandomString([]string{"25k - 50k", "50k - 75k", "75k - 99k", "100k+", "25k and below", "unknown"}),
			gofakeit.URL(),
			gofakeit.CountryAbr(),
			gofakeit.Price(0.01, 20),
			gofakeit.Date())

		_, err = io.WriteString(file, s)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	file.Close()

	var maxNetworks int

	switch os.Args[1] {
	case "small":
		maxNetworks = 10
	case "medium":
		maxNetworks = 100
	case "large":
		maxNetworks = 1_000
	default:
		printUsage()
	}

	file, err = os.Create("networks-" + os.Args[1] + ".tsv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= maxNetworks; i++ {
		s := fmt.Sprintf("%v\t%v\n", i, gofakeit.Company())

		_, err = io.WriteString(file, s)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	file.Close()

	var maxPlatforms int

	switch os.Args[1] {
	case "small":
		maxPlatforms = 10
	case "medium":
		maxPlatforms = 100
	case "large":
		maxPlatforms = 1_000
	default:
		printUsage()
	}

	file, err = os.Create("platforms-" + os.Args[1] + ".tsv")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= maxPlatforms; i++ {
		s := fmt.Sprintf("%v\t%v\n", i, gofakeit.Company())

		_, err = io.WriteString(file, s)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	file.Close()
}

func printUsage() {
	fmt.Printf("Usage: %s <small | medium | large>\n", os.Args[0])
	os.Exit(1)
}

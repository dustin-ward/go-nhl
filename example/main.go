package main

import (
	"fmt"
	"log"

	"github.com/dustin-ward/go-nhl/v1/nhl"
)

func main() {
	playerIds := []int{
		8478402,
		8478403,
		8478404,
		8478405,
		8478406,
	}

	for _, pid := range playerIds {
		p, err := nhl.GetPlayer(pid)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Player ============")
		fmt.Printf("  Name: %s %s\n", p.GetFirstName().GetDefault(), p.GetLastName().GetDefault())
		fmt.Printf("  #: %d\n", p.GetSweaterNumber())
		fmt.Printf("  Birth: %s, %s - %s\n", p.GetBirthCity().GetDefault(), p.GetBirthCountry(), p.GetBirthDate())
		fmt.Printf("  Height: %d'%d\" Weight: %dlbs\n", p.GetHeightInInches()/12, p.GetHeightInInches()%12, p.GetWeightInPounds())
	}
}

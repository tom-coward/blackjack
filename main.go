/*
Copyright © 2023 Tom Coward tom@tomcoward.me

The executable CLI interface for the Blackjack game.
*/

package main

import (
	"fmt"
	"os"

	"github.com/tom-coward/blackjack/blackjack"
)

func main() {
	game := blackjack.NewGame()

	err := game.DealOpeningHands()
	if err != nil {
		fmt.Printf("There was an error dealing the opening hand: %s", err.Error())
		os.Exit(-1)
	}

	for !game.Complete && game.PlayerScore < 21 {
		fmt.Println("----------------------------------------")
		fmt.Print("Your current hand is: \n")
		for _, card := range game.PlayerDeck {
			fmt.Printf("- %s\n", card.Identifier)
		}

		fmt.Printf("\nThe dealer's revealed card is %s - they also have one more hidden card\n", game.HouseDeck[0].Identifier)

		fmt.Printf("\nYour current score is %d. Would you like to hit or stand? (\"H\" or \"S\"): ", game.PlayerScore)
		var playerChoice string
		fmt.Scan(&playerChoice)

		switch playerChoice {
		case "H":
			cards, err := game.DealToPlayer(1)

			if err != nil {
				fmt.Printf("There was an error Hitting: %s", err.Error())
				os.Exit(-1)
			}

			fmt.Printf("\nYou were dealt a %s", cards[0].Identifier)
		case "S":
			fmt.Printf("\nThe dealer's second card is %s", game.HouseDeck[1].Identifier)
			err := game.Stand()

			if err != nil {
				fmt.Printf("There was an error Standing: %s", err.Error())
				os.Exit(-1)
			}
		default:
			fmt.Println("\nYour input was invalid")
		}
	}

	fmt.Println("Your final hand is:")
	for _, card := range game.PlayerDeck {
		fmt.Printf("- %s\n", card.Identifier)
	}

	fmt.Println("The dealer's final hand is:")
	for _, card := range game.HouseDeck {
		fmt.Printf("- %s\n", card.Identifier)
	}

	if game.PlayerBust {
		fmt.Printf("You're bust! Your score exceeded 21 at %d", game.PlayerScore)
		os.Exit(1)
	}

	if game.HouseBust {
		fmt.Printf("The house is bust! The dealer's score exceeded 21 at %d", game.HouseScore)
	}

	if game.PlayerWon {
		fmt.Printf("You won! Your final score was %d vs. the house (dealer)'s score of %d", game.PlayerScore, game.HouseScore)
	}

	os.Exit(1)
}

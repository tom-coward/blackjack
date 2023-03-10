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
		fmt.Printf("\nThere was an error dealing the opening hand: %s", err.Error())
		os.Exit(-1)
	}

	for !game.Complete && game.PlayerScore < 21 {
		fmt.Println("\nYour current hand is:")
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
				fmt.Printf("\nThere was an error Hitting: %s", err.Error())
				os.Exit(-1)
			}

			fmt.Printf("\nYou were dealt a %s\n", cards[0].Identifier)
		case "S":
			fmt.Printf("\nThe dealer's second card is %s\n", game.HouseDeck[1].Identifier)
			err := game.Stand()

			if err != nil {
				fmt.Printf("\nThere was an error Standing: %s", err.Error())
				os.Exit(-1)
			}
		default:
			fmt.Println("\nYour input was invalid")
		}
	}

	fmt.Println("\nYour final hand is:")
	for _, card := range game.PlayerDeck {
		fmt.Printf("- %s\n", card.Identifier)
	}

	fmt.Println("\nThe dealer's final hand is:")
	for _, card := range game.HouseDeck {
		fmt.Printf("- %s\n", card.Identifier)
	}

	if game.PlayerBust {
		fmt.Printf("\nYou're bust! Your score exceeded 21 at %d", game.PlayerScore)
		os.Exit(1)
	}

	if game.HouseBust {
		fmt.Printf("\nThe house is bust, so you won! The dealer's score exceeded 21 at %d", game.HouseScore)
		os.Exit(1)
	}

	if game.PlayerWon {
		fmt.Printf("\nYou won! Your final score was %d vs. the house (dealer)'s score of %d", game.PlayerScore, game.HouseScore)
		os.Exit(1)
	}

	if game.Draw {
		fmt.Printf("\nYou drew! You and the dealer's both had a final score of %d", game.PlayerScore)
		os.Exit(1)
	}

	fmt.Printf("\nYou lost! Your final score was %d vs. the house (dealer)'s score of %d", game.PlayerScore, game.HouseScore)
	os.Exit(1)
}

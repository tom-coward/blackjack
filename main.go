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

	game.DealOpeningHands()

	for !game.Complete {
		fmt.Print("Your current hand is: \n")
		for _, card := range game.PlayerDeck {
			fmt.Printf("%s", card.Identifier)
		}

		fmt.Printf("\nYour current score is %d. Would you like to hit or stand? (\"H\" or \"S\")", game.PlayerScore)
		var playerChoice string
		fmt.Scan(&playerChoice)

		switch playerChoice {
		case "H":
			game.DealToPlayer(1)
		case "S":
			game.Stand()
		}
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
}

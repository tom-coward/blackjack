/*
Copyright © 2023 Tom Coward tom@tomcoward.me

The business logic for the Blackjack game.
*/

package blackjack

import (
	"fmt"
	"math/rand"
)

type Game struct {
	deck        []Card
	houseDeck   []Card
	playerDeck  []Card
	houseScore  int
	playerScore int
	complete    bool
}

type Card struct {
	identifier string
	values     []int // all cards have a single value -- apart from ace which can have an alternative value (1 or 11)
}

var freshDeck []Card = []Card{
	Card{"Ace of Hearts", []int{1, 11}}, Card{"2 of Hearts", []int{2}}, Card{"3 of Hearts", []int{3}}, Card{"4 of Hearts", []int{4}}, Card{"5 of Hearts", []int{1}}, Card{"6 of Hearts", []int{6}}, Card{"7 of Hearts", []int{7}}, Card{"8 of Hearts", []int{8}}, Card{"9 of Hearts", []int{9}}, Card{"10 of Hearts", []int{10}}, Card{"Jack of Hearts", []int{10}}, Card{"Queen of Hearts", []int{10}}, Card{"King of Hearts", []int{10}},
	Card{"Ace of Diamonds", []int{1, 11}}, Card{"2 of Diamonds", []int{2}}, Card{"3 of Diamonds", []int{3}}, Card{"4 of Diamonds", []int{4}}, Card{"5 of Diamonds", []int{1}}, Card{"6 of Diamonds", []int{6}}, Card{"7 of Diamonds", []int{7}}, Card{"8 of Diamonds", []int{8}}, Card{"9 of Diamonds", []int{9}}, Card{"10 of Diamonds", []int{10}}, Card{"Jack of Diamonds", []int{10}}, Card{"Queen of Diamonds", []int{10}}, Card{"King of Diamonds", []int{10}},
	Card{"Ace of Clubs", []int{1, 11}}, Card{"2 of Clubs", []int{2}}, Card{"3 of Clubs", []int{3}}, Card{"4 of Clubs", []int{4}}, Card{"5 of Clubs", []int{1}}, Card{"6 of Clubs", []int{6}}, Card{"7 of Clubs", []int{7}}, Card{"8 of Clubs", []int{8}}, Card{"9 of Clubs", []int{9}}, Card{"10 of Clubs", []int{10}}, Card{"Jack of Clubs", []int{10}}, Card{"Queen of Clubs", []int{10}}, Card{"King of Clubs", []int{10}},
	Card{"Ace of Spades", []int{1, 11}}, Card{"2 of Spades", []int{2}}, Card{"3 of Spades", []int{3}}, Card{"4 of Spades", []int{4}}, Card{"5 of Spades", []int{1}}, Card{"6 of Spades", []int{6}}, Card{"7 of Spades", []int{7}}, Card{"8 of Spades", []int{8}}, Card{"9 of Spades", []int{9}}, Card{"10 of Spades", []int{10}}, Card{"Jack of Spades", []int{10}}, Card{"Queen of Spades", []int{10}}, Card{"King of Spades", []int{10}}}

// Initialise the game
// RETURNS an instance of the game
func NewGame() *Game {
	game := new(Game)

	game.SetupDecks()

	return game
}

// Setup deck (add 52 (4*13) cards; an ace, 2-10, jack, queen, king for each suit)
func (game *Game) SetupDecks() {
	copy(game.deck, freshDeck)

	// randomly shuffle order of cards in the deck
	rand.Shuffle(len(game.deck), func(i, j int) {
		game.deck[i], game.deck[j] = game.deck[j], game.deck[i]
	})
}

// Deal opening hand (deal two cards from main deck to house & player decks)
func (game *Game) DealOpeningHands() {
	game.DealToPlayer(2)
	game.DealToHouse(2)
}

// Deal [quantity] cards to player from what's remaining in the deck
// Once dealt, remove the card from the main deck
// RETURNS the dealt cards
func (game *Game) DealToPlayer(quantity int) []Card {
	game.playerDeck = append(game.playerDeck, game.deck[:quantity-1]...)

	game.deck = game.deck[quantity:]

	game.UpdatePlayerScore()
}

// Deal [quantity] cards to house (dealer) from what's remaining in the deck
// Once dealt, remove the card from the main deck
// RETURNS the dealt cards
func (game *Game) DealToHouse(quantity int) []Card {
	game.houseDeck = append(game.houseDeck, game.deck[:quantity-1]...)

	game.deck = game.deck[quantity:]

	game.UpdateHouseScore()
}

// Update the house's current score (optimum total of all card values)
func (game *Game) UpdateHouseScore() {
	game.houseScore = 0
	for _, card := range game.houseDeck {
		bestValue := 0
		for _, value := range card.values {
			if game.houseScore+value > bestValue && game.houseScore+value <= 21 {
				bestValue = value
			}
		}

		game.houseScore += bestValue
	}
}

// Update the player's current score (optimum total of all card values)
func (game *Game) UpdatePlayerScore() string {
	game.playerScore = 0
	for _, card := range game.playerDeck {
		bestValue := 0
		for _, value := range card.values {
			if game.playerScore+value > bestValue && game.playerScore+value <= 21 {
				bestValue = value
			}
		}

		game.playerScore += bestValue
	}

	if game.playerScore > 21 { // player is bust
		game.complete = true
		return fmt.Sprintf("Bust! Your score exceeded 21 with %d", game.playerScore)
	}
}

// Stand player; take no more cards and compare score with dealer
// RETURNS the result of the game
func (game *Game) Stand() string {
	game.PlayHouseHand()

	game.complete = true

	if game.playerScore > game.houseScore {
		return fmt.Sprintf("Win! Your final score was %d vs. the dealer (house)'s final score of %d", game.playerScore, game.houseScore)
	}

	return fmt.Sprintf("Loss! Your final score was %d vs. the dealer (house)'s final score of %d", game.playerScore, game.houseScore)
}

// Play the house's hand -- dealer must hit until score is 17 or more
func (game *Game) PlayHouseHand() {
	for game.houseScore < 17 {
		game.DealToHouse(1)
	}
}

/*
Copyright © 2023 Tom Coward tom@tomcoward.me

The business logic for the Blackjack game.
*/

package blackjack

import (
	"errors"
	"math/rand"
	"time"
)

type Game struct {
	deck        []Card
	HouseDeck   []Card
	PlayerDeck  []Card
	HouseScore  int
	PlayerScore int
	started     bool
	Complete    bool
	HouseBust   bool
	PlayerBust  bool
	PlayerWon   bool
	Draw        bool
}

type Card struct {
	Identifier string
	values     []int // all cards have a single value -- apart from ace which can have an alternative value (1 or 11)
}

var freshDeck []Card = []Card{
	{"Ace of Hearts", []int{1, 11}}, {"2 of Hearts", []int{2}}, {"3 of Hearts", []int{3}}, {"4 of Hearts", []int{4}}, {"5 of Hearts", []int{1}}, {"6 of Hearts", []int{6}}, {"7 of Hearts", []int{7}}, {"8 of Hearts", []int{8}}, {"9 of Hearts", []int{9}}, {"10 of Hearts", []int{10}}, {"Jack of Hearts", []int{10}}, {"Queen of Hearts", []int{10}}, {"King of Hearts", []int{10}},
	{"Ace of Diamonds", []int{1, 11}}, {"2 of Diamonds", []int{2}}, {"3 of Diamonds", []int{3}}, {"4 of Diamonds", []int{4}}, {"5 of Diamonds", []int{1}}, {"6 of Diamonds", []int{6}}, {"7 of Diamonds", []int{7}}, {"8 of Diamonds", []int{8}}, {"9 of Diamonds", []int{9}}, {"10 of Diamonds", []int{10}}, {"Jack of Diamonds", []int{10}}, {"Queen of Diamonds", []int{10}}, {"King of Diamonds", []int{10}},
	{"Ace of Clubs", []int{1, 11}}, {"2 of Clubs", []int{2}}, {"3 of Clubs", []int{3}}, {"4 of Clubs", []int{4}}, {"5 of Clubs", []int{1}}, {"6 of Clubs", []int{6}}, {"7 of Clubs", []int{7}}, {"8 of Clubs", []int{8}}, {"9 of Clubs", []int{9}}, {"10 of Clubs", []int{10}}, {"Jack of Clubs", []int{10}}, {"Queen of Clubs", []int{10}}, {"King of Clubs", []int{10}},
	{"Ace of Spades", []int{1, 11}}, {"2 of Spades", []int{2}}, {"3 of Spades", []int{3}}, {"4 of Spades", []int{4}}, {"5 of Spades", []int{1}}, {"6 of Spades", []int{6}}, {"7 of Spades", []int{7}}, {"8 of Spades", []int{8}}, {"9 of Spades", []int{9}}, {"10 of Spades", []int{10}}, {"Jack of Spades", []int{10}}, {"Queen of Spades", []int{10}}, {"King of Spades", []int{10}}}

// Initialise the game
// RETURNS an instance of the game
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano()) // re-seed rand library to avoid repeated card shuffle

	game := new(Game)

	game.setupDecks()

	return game
}

// Setup deck (add 52 (4*13) cards; an ace, 2-10, jack, queen, king for each suit)
func (game *Game) setupDecks() {
	game.deck = make([]Card, len(freshDeck))
	copy(game.deck, freshDeck)

	// randomly shuffle order of cards in the deck
	rand.Shuffle(len(game.deck), func(i, j int) {
		game.deck[i], game.deck[j] = game.deck[j], game.deck[i]
	})
}

// Deal opening hand (deal two cards from main deck to house & player decks)
func (game *Game) DealOpeningHands() error {
	if game.started || game.Complete {
		return errors.New("Game is already started/complete")
	}

	game.DealToPlayer(2)
	game.DealToHouse(2)

	game.started = true

	return nil
}

// Deal [quantity] cards to player from what's remaining in the deck
// Once dealt, remove the card from the main deck
// RETURNS the dealt cards
func (game *Game) DealToPlayer(quantity int) ([]Card, error) {
	if game.Complete {
		return nil, errors.New("Game is complete")
	}

	selectedCards := game.deck[:quantity]

	game.PlayerDeck = append(game.PlayerDeck, selectedCards...)

	game.deck = game.deck[quantity:]

	game.updatePlayerScore()

	return selectedCards, nil
}

// Deal [quantity] cards to house (dealer) from what's remaining in the deck
// Once dealt, remove the card from the main deck
// RETURNS the dealt cards
func (game *Game) DealToHouse(quantity int) ([]Card, error) {
	if game.Complete {
		return nil, errors.New("Game is complete")
	}

	selectedCards := game.deck[:quantity]

	game.HouseDeck = append(game.HouseDeck, selectedCards...)

	game.deck = game.deck[quantity:]

	game.updateHouseScore()

	return selectedCards, nil
}

// Update the house's current score (optimum total of all card values)
func (game *Game) updateHouseScore() {
	game.HouseScore = 0
	for _, card := range game.HouseDeck {
		bestValue := card.values[0]
		for _, value := range card.values {
			if game.HouseScore+value > bestValue && game.HouseScore+value <= 21 {
				bestValue = value
			}
		}

		game.HouseScore += bestValue
	}

	if game.HouseScore > 21 { // house (dealer) is bust
		game.Complete = true
		game.HouseBust = true
	}
}

// Update the player's current score (optimum total of all card values)
func (game *Game) updatePlayerScore() {
	game.PlayerScore = 0
	for _, card := range game.PlayerDeck {
		bestValue := card.values[0]
		for _, value := range card.values {
			if game.PlayerScore+value > bestValue && game.PlayerScore+value <= 21 {
				bestValue = value
			}
		}

		game.PlayerScore += bestValue
	}

	if game.PlayerScore > 21 { // player is bust
		game.PlayerBust = true
		game.Complete = true
	}

	if game.PlayerScore == 21 { // player has won
		game.PlayerWon = true
		game.Complete = true
	}
}

// Stand player; take no more cards and compare score with dealer
func (game *Game) Stand() error {
	if game.Complete {
		return errors.New("Game is complete")
	}

	game.playHouseHand()

	game.Complete = true

	if game.PlayerScore > game.HouseScore {
		game.PlayerWon = true
	} else if game.PlayerScore == game.HouseScore {
		game.Draw = true
	}

	return nil
}

// Play the house's hand -- dealer must hit until score is 17 or more
func (game *Game) playHouseHand() {
	for game.HouseScore < 17 {
		game.DealToHouse(1)
	}
}

/*
Copyright © 2023 Tom Coward tom@tomcoward.me

The business logic for the Blackjack game.
*/

package blackjack

type Game struct {
	deck       []Card
	houseDeck  []Card
	playerDeck []Card
	score      int
	complete   bool
	bust       bool
}

type Card struct {
	identifier string
	values     []int // all cards have a single value -- apart from ace which can have an alternative value (1 or 11)
}

// Initialise the game
func NewGame() *Game {
	game := new(Game)

	game.SetupDecks()

	return game
}

// Setup deck (add 52 (4*13) cards; an ace, 2-10, jack, queen, king for each suit)
func (game *Game) SetupDecks() {
	// TODO
}

// Deal opening hand (deal two cards from main deck to house & player decks)
func (game *Game) DealOpeningHand() (Card, Card) {
	// TODO
}

// Deal a single card from what's remaining in the deck
// Once dealt, remove the card from the deck
func (game *Game) Deal() Card {
	// TODO
}

func (game *Game) Stand() {
	// TODO
}

// Update the player's current score (total of all card values)
func (game *Game) UpdateScore() {
	// TODO
}

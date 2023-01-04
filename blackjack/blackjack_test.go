/*
Copyright © 2023 Tom Coward tom@tomcoward.me

A suite of unit tests for the Blackjack business logic.
*/

package blackjack

import "testing"

// Given I play a game of blackjack
// When I am dealt my opening hand
// Then I have two cards
func TestDealOpeningHands(t *testing.T) {
	game := NewGame()

	game.DealOpeningHands()

	if len(game.PlayerDeck) != 2 {
		t.Errorf("Expected 2 cards in player deck, got %d", len(game.PlayerDeck))
	}

	if len(game.houseDeck) != 2 {
		t.Errorf("Expected 2 cards in house deck, got %d", len(game.PlayerDeck))
	}

	if len(game.deck) != 50 {
		t.Errorf("Expected 50 cards remaining in main deck, got %d", len(game.deck))
	}
}

// Given I have a valid hand of cards
// When I choose to ‘hit’
// Then I receive another card
// And my score is updated
func TestHit(t *testing.T) {
	game := NewGame() // start new game with starting score of 0
	oldGameScore := game.PlayerScore

	cards, _ := game.DealToPlayer(1)

	if cards[0].values[1]+oldGameScore != game.PlayerScore {
		t.Errorf("Expected score to be %d, got %d", cards[0].values[1]+oldGameScore, game.PlayerScore)
	}
}

// Given I have a valid hand of cards
// When I choose to ‘stand’
// Then I receive no further cards
// And my score is evaluated
func TestStand(t *testing.T) {
	game := NewGame()
	oldPlayerDeckLen := len(game.PlayerDeck)

	game.Stand()

	if len(game.PlayerDeck) != oldPlayerDeckLen {
		t.Errorf("Expected player deck length to be %d, got %d", oldPlayerDeckLen, len(game.PlayerDeck))
	}

	if !game.Complete {
		t.Errorf("Expected game to be complete, but it's incomplete")
	}
}

// Given my score is updated or evaluated
// When it is 21 or less
// Then I have a valid hand
func TestUpdateScoreUnder21(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{Card{"Ace", []int{1, 11}}, Card{"10", []int{10}}}

	game.updatePlayerScore()

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}

	if game.PlayerBust {
		t.Errorf("Expected game to still be active, but it's bust")
	}
}

// Given my score is updated
// When it is 22 or more
// Then I am ‘bust’ and do not have a valid hand
func TestUpdateScoreOver21(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{Card{"Ace", []int{1, 11}}, Card{"10", []int{10}}, Card{"2", []int{2}}}

	game.updatePlayerScore()

	if game.PlayerScore != 23 {
		t.Errorf("Expected score to be 23, got %d", game.PlayerScore)
	}

	if !game.PlayerBust {
		t.Errorf("Expected game to be bust, but it's still active")
	}
}

// Given I have a king and an ace
// When my score is evaluated
// Then my score is 21
func TestKingAceEquals21Score(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{Card{"King", []int{10}}, Card{"Ace", []int{1, 11}}}

	game.updatePlayerScore()

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}
}

// Given I have a king, a queen, and an ace
// When my score is evaluated
// Then my score is 21
func TestKingQueenAceEquals21Score(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{Card{"King", []int{10}}, Card{"Queen", []int{10}}, Card{"Ace", []int{1, 11}}}

	game.updatePlayerScore()

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}
}

// Given that I have a nine, an ace, and another ace
// When my score is evaluated
// Then my score is 21
func TestNineAceAceEquals21Score(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{Card{"9", []int{9}}, Card{"Ace", []int{1, 11}}, Card{"Ace", []int{1, 11}}}

	game.updatePlayerScore()

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}
}

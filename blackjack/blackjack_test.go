/*
Copyright © 2023 Tom Coward tom@tomcoward.me

A suite of unit tests for the Blackjack business logic.
*/

package blackjack

import "testing"

// Given I play a game of blackjack
// When I am dealt my opening hand
// Then I have two cards, and the house has two cards
// And the main deck is updated
func TestDealOpeningHands(t *testing.T) {
	game := NewGame()

	game.DealOpeningHands()

	if len(game.PlayerDeck) != 2 {
		t.Errorf("Expected 2 cards in player deck, got %d", len(game.PlayerDeck))
	}

	if len(game.HouseDeck) != 2 {
		t.Errorf("Expected 2 cards in house deck, got %d", len(game.HouseDeck))
	}

	if len(game.deck) != 48 {
		t.Errorf("Expected 48 cards remaining in main deck, got %d", len(game.deck))
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

	// try each possible card value to check if score is updated
	scoreValidated := false
	for _, value := range cards[0].values {
		if value+oldGameScore == game.PlayerScore {
			scoreValidated = true
		}
	}

	if !scoreValidated {
		t.Errorf("Expected score to be %d, got %d", cards[0].values[0]+oldGameScore, game.PlayerScore)
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
// Then I have a valid hand (not bust & game is still active)
func TestUpdateScoreUnder21(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{{"5", []int{5}}, {"10", []int{10}}}

	game.updatePlayerScore() // score should be 15

	if game.PlayerBust {
		t.Errorf("Expected player to have valid hand, but they're bust")
	}

	if game.Complete {
		t.Errorf("Expected game to still be active, but it's complete")
	}
}

// Given my score is updated
// When it is 22 or more
// Then I am ‘bust’ and do not have a valid hand
func TestUpdateScoreOver21(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{{"10", []int{10}}, {"10", []int{10}}, {"2", []int{2}}}

	game.updatePlayerScore() // score should be 22

	if game.PlayerScore != 22 {
		t.Errorf("Expected score to be 22, got %d", game.PlayerScore)
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

	game.PlayerDeck = []Card{{"King", []int{10}}, {"Ace", []int{1, 11}}}

	game.updatePlayerScore() // score should be 21 (also testing correct ace value is chosen)

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}
}

// Given I have a king, a queen, and an ace
// When my score is evaluated
// Then my score is 21
func TestKingQueenAceEquals21Score(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{{"King", []int{10}}, {"Queen", []int{10}}, {"Ace", []int{1, 11}}}

	game.updatePlayerScore() // score should be 21 (also testing correct ace value is chosen)

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}
}

// Given that I have a nine, an ace, and another ace
// When my score is evaluated
// Then my score is 21
func TestNineAceAceEquals21Score(t *testing.T) {
	game := NewGame()

	game.PlayerDeck = []Card{{"9", []int{9}}, {"Ace", []int{1, 11}}, {"Ace", []int{1, 11}}}

	game.updatePlayerScore() // score should be  21 (also testing correct ace values are chosen)

	if game.PlayerScore != 21 {
		t.Errorf("Expected score to be 21, got %d", game.PlayerScore)
	}
}

# tom-coward/blackjack
This is a straightforward Blackjack (otherwise known as Twenty-One) CLI game, written in Go.

## Assumptions
- The game is one player vs. the house (dealer)
- The dealer (house) draws two cards at the opening hand (as does the player) - then once the player has decided to Stand, they will draw more cards until the score is >= 17
- A score for the player or house > 21 means that they're bust, and the game is over (won by the opposing player to whoever is bust)
- There are four instances of each card in the main deck at the beginning of the game (one for each suit)
- As a card is dealt to either the player or dealer, it is simultaneously removed from the main deck so it cannot be dealt again

## Gameplay
**To play a round of the game, simply run `go run .` from the root directory.**

The game will setup & randomly shuffle the order of cards in the main deck (consiting of 52 cards), and deal the opening hand (deal 2 cards to the player & the dealer). As the player, you will be informed of one of the dealer's cards but one will remain secret until the dealer deals their final hand. The game will automatically calculate the optimum score of the current hand of cards (i.e. should an ace have the value 1 or 11?).

The aim of the game is for you (as the player) to achieve a higher score (sum of the value of all your cards) than the dealer, who's representing the house.

At this point, the dealer's cards are fixed until the conclusion of the game and it's up to you as the player how to proceed - you may choose to either *Hit* (take another card from the main deck) or *Stand* (stick with your current hand of cards).

**Standing** will result in the dealer revealing their second card, and then dealing themselves cards until their score reaches 17 or more - at this point (provided the dealer isn't bust), the player/dealer with the highest score wins.

If at any point your score exceeds 21, you're *bust*, and the house (dealer) has won the game; the inverse applies if the dealer is bust when they deal more hands at the end of the game.

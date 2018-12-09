package day9

import (
	"fmt"
)

// Player plays the Game
type Player struct {
	score int
}

// Game describes an elf game
type Game struct {
	players []Player
	LastMarble int
}

// NewGameParams gets a new Game with parameters
func NewGameParams(players, lastMarble int) Game {
	return Game{
		players: make([]Player, players),
		LastMarble: lastMarble,
	}
}

// NewGameStr get a new Game by parsing a string
func NewGameStr(input string) Game {
	var players, lastMarble int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &lastMarble)

	return Game{
		players: make([]Player, players),
		LastMarble: lastMarble,
	}
}
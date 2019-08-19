package Structures

import (
	"github.com/satori/go.uuid"
	"errors"
	"strconv"
)

	func NewGame() Game {

		game := Game{}
		game.ID = uuid.Must(uuid.NewV4()).String()
		return game
	}

	type Game struct {

		ID string
		PlayerOne, PlayerTwo Player
		nextPlayer int
		Board [3][3]Point
	}

	func (g *Game) IsWon() bool {

		return false
	}

	func (g *Game) Play(playerCode string, x, y int) (bool, error) {

		// Is this a valid player of this game?
		var ourPlayer Player
		if(g.PlayerOne.Code == playerCode && g.nextPlayer == 1) {

			ourPlayer = g.PlayerOne

		} else if(g.PlayerTwo.Code == playerCode && g.nextPlayer == 2) {

			ourPlayer = g.PlayerTwo

		} else {

			// Find your own game
			return false, errors.New("You are not playing in this game")
		}

		// Test for invalid positions
		if(x < 0 || x > 2 || y < 0 || y > 2) {

			return false, errors.New("You have supplied a missing or invalid board location")
		}

		// Now, has someone claimed this spot already?
		if(len(g.Board[x][y].Player.Code) > 0) {

			// This is taken
			return false, errors.New("This position (" + strconv.Itoa(x) + "," + strconv.Itoa(y) + ") has been claimed")
		}

		// Take this
		g.Board[x][y].Player = ourPlayer

		// Done
		return true, nil
	}
package Structures

import (
	"github.com/satori/go.uuid"
	"errors"
	"strconv"
)

	func NewGame() Game {

		game := Game{ ID: uuid.Must(uuid.NewV4()).String(), Active: false, WaitingForPlayer: true, Won: false, nextPlayer: 1 }
		return game
	}

	type Game struct {

		ID string
		PlayerOne, PlayerTwo Player
		nextPlayer int
		Board [3][3]Point
		Winner Player
		WaitingForPlayer, Active, Won bool
	}

	func (g *Game) IsWon() bool {

		return false
	}

	func (g *Game) IsPlayer(playerCode string) bool {

		if(g.PlayerOne.Code != playerCode && g.PlayerTwo.Code != playerCode) {

			return false
		}

		return true
	}

	func (g *Game) Play(playerCode string, x, y int) (bool, error) {

		if(!g.Active) {

			if(g.Won) {

				return false, errors.New("This game has ended. It was won by " + g.Winner.Name)

			} else if(g.WaitingForPlayer) {

				return false, errors.New("This game is waiting for a second player to join")

			} else {

				return false, errors.New("This game is not active for some reason")
			}
		}

		// Is this a valid player of this game?
		var ourPlayer Player
		var nextPlayer int
		if(g.PlayerOne.Code == playerCode) {

			if(g.nextPlayer != 1) {

				return false, errors.New("It is not your turn to play - " + g.PlayerTwo.Name + " is up next")
			}

			ourPlayer = g.PlayerOne
			nextPlayer = 2

		} else if(g.PlayerTwo.Code == playerCode) {

			if(g.nextPlayer != 2) {

				return false, errors.New("It is not your turn to play - " + g.PlayerOne.Name + " is up next")
			}

			ourPlayer = g.PlayerTwo
			nextPlayer = 1

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
			return false, errors.New("This position (" + strconv.Itoa(x) + "," + strconv.Itoa(y) + ") has been claimed by " + g.Board[x][y].Player.Name)
		}

		// Take this
		g.Board[x][y].Owned = true
		g.Board[x][y].Player = ourPlayer
		g.nextPlayer = nextPlayer

		// Also calculate if we have won
		g.CalculateWinner()

		// Done
		return true, nil
	}

	// Calculate if there is a winner at this stage
	func (g *Game) CalculateWinner() {

		// There are only eight winning lines in tic tac toe. if this were a 4x4, 5x5 board etc I would pobably need
		// a smarter line detection algorithm but that is I think, a whole separate exercise in itself.

		// Left column, straight down
		if(g.Board[0][0].Owned && g.Board[0][0].Player.Code == g.Board[0][1].Player.Code && g.Board[0][0].Player.Code == g.Board[0][2].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[0][0].Player
		}

		// Middle column, straight down
		if(g.Board[1][0].Owned && g.Board[1][0].Player.Code == g.Board[1][1].Player.Code && g.Board[1][0].Player.Code == g.Board[1][2].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[1][0].Player
		}

		// Right column, straight down
		if(g.Board[2][0].Owned && g.Board[2][0].Player.Code == g.Board[2][1].Player.Code && g.Board[2][0].Player.Code == g.Board[2][2].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[2][0].Player
		}

		// Left row, across
		if(g.Board[0][0].Owned && g.Board[0][0].Player.Code == g.Board[1][0].Player.Code && g.Board[0][0].Player.Code == g.Board[2][0].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[0][0].Player
		}

		// Middle row, straight across
		if(g.Board[0][1].Owned && g.Board[0][1].Player.Code == g.Board[1][1].Player.Code && g.Board[0][1].Player.Code == g.Board[2][1].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[0][1].Player
		}

		// Bottom row, straight across
		if(g.Board[0][2].Owned && g.Board[0][2].Player.Code == g.Board[1][2].Player.Code && g.Board[0][2].Player.Code == g.Board[2][2].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[0][2].Player
		}

		// Diagonal left to right
		if(g.Board[0][0].Owned && g.Board[0][0].Player.Code == g.Board[1][1].Player.Code && g.Board[0][0].Player.Code == g.Board[2][2].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[0][0].Player
		}

		// Diagonal, right to left
		if(g.Board[2][0].Owned && g.Board[2][0].Player.Code == g.Board[1][1].Player.Code && g.Board[2][0].Player.Code == g.Board[0][2].Player.Code) {

			g.Active = false
			g.Won = true
			g.Winner = g.Board[2][0].Player
		}
	}
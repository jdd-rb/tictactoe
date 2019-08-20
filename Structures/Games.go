package Structures

import (
	"sync"
	"errors"
)

	func SetupGames() Games {

		games := Games{}

		// Waiting list of games that need a partner
		games.waitingList = make([]string, 0)
		games.games = make(map[string]Game)

		return games
	}

	type Games struct {

		lock sync.RWMutex
		waitingList []string
		games map[string]Game
	}

	func (g *Games) Exists(gameID string) bool {

		// Read lock for this one
		g.lock.RLock()
		defer g.lock.RUnlock()

		_, exists := g.games[gameID]
		return exists
	}

	func (g *Games) GetStatus(gameID string) (GameStatus,bool) {

		// Get a read lock
		g.lock.RLock()
		defer g.lock.RUnlock()

		game, exists := g.games[gameID]

		if(!exists) {

			return GameStatus{}, false
		}

		status := GameStatus{ Active: game.Active, WaitingForPlayer: game.WaitingForPlayer, Won: game.IsWon()}
		status.Board = game.Board
		status.PlayerOne = game.PlayerOne.Name
		status.PlayerTwo = game.PlayerTwo.Name

		if(game.IsWon()) {

			status.Winner = game.Winner.Name
		}

		return status, true
	}

	func (g *Games) JoinGame(player Player) string {

		// Gain exclusivity while we do this
		g.lock.Lock()
		defer g.lock.Unlock()

		// Ultimately this is the game the player has joined
		var gameID string

		// Do we have at least one game in the queue?
		if(len(g.waitingList) < 1) {

			// Nothing in the queue right now. This player can be player one in a brand new game
			newGame := NewGame()
			newGame.PlayerOne = player

			// Add this to the roster of games
			g.games[newGame.ID] = newGame

			// This is also our player's new game ID
			gameID = newGame.ID

			// Add this to the waiting list
			g.waitingList = append(g.waitingList, newGame.ID)

		} else {

			// Just pick the first element from the waiting list
			joinGameID := g.waitingList[0]
			joinGame := g.games[joinGameID]
			joinGame.PlayerTwo = player
			joinGame.WaitingForPlayer = false
			joinGame.Active = true
			g.games[joinGameID] = joinGame
			gameID = joinGameID
		}

		return gameID
	}

	func (g *Games) MakeAMove(gameID string, playerID string, x, y int) (bool, error) {

		// Firstly, gain exclusivity, for a moment
		g.lock.Lock()
		defer g.lock.Unlock()

		// Grab our game and a flag if this exists or not
		game, exists := g.games[gameID]

		if(!exists) {

			// OK so this game doesn't even exist
			return false, errors.New("This game does not exist")
		}

		// Is this player either player one, or player two, of this game?
		if(game.PlayerOne.Code != playerID && game.PlayerTwo.Code != playerID) {

			return false, errors.New("You are not playing in this game")
		}

		// OK, so play the move, if that's valid
		isValid, err := game.Play(playerID, x, y)

		if(!isValid) {

			return false, err
		}

		// Save the gaem back
		g.games[gameID] = game
		return true, nil
	}

	func (g *Games) IsPlayingGame(gameID string, playerID string) bool {

		// Get a read lock
		g.lock.RLock()

		// Grab the game
		game, exists := g.games[gameID]

		// Release the read lock
		g.lock.RUnlock()

		if(!exists) {

			return false
		}

		// Are they a player?
		return game.IsPlayer(playerID)
	}

	func (g *Games) GetBoard(gameID string) ([3][3]Point, bool) {

		// Write lock - ensure this readout of the board is to the split second accurate
		g.lock.RLock()
		defer g.lock.Unlock()

		// Grab the game
		game, exists := g.games[gameID]

		if(!exists) {

			return [3][3]Point{}, false
		}

		return game.Board, true
	}

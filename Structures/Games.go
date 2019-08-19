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
	}

	type Games struct {

		lock sync.RWMutex
		waitingList []string
		games map[string]Game
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

		return true, nil
	}

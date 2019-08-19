package Structures

import (
	"sync"
	"github.com/satori/go.uuid"
)

	func SetupPlayers() Players {

		myPlayers := Players{}
		myPlayers.players = make(map[string]Player)

		return myPlayers
	}

	type Players struct {

		players map[string]Player
		lock sync.RWMutex
	}

	func (p *Players) Create(name string) Player {

		// Write lock while we do this
		p.lock.Lock()
		defer p.lock.Unlock()

		// Names don't have to be exclusive - just create the new player
		newPlayer := Player{ Name: name, Code: uuid.Must(uuid.NewV4()).String() }
		p.players[newPlayer.Code] = newPlayer

		// Done
		return newPlayer
	}

	func (p *Players) Get(playerID string) (Player,bool) {

		// Read lock while we do this
		p.lock.RLock()
		defer p.lock.RUnlock()

		// Get the player and a check if they exist
		player, exists := p.players[playerID]

		if(!exists) {

			return Player{}, false
		}

		return player, true
	}

package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"os"
	"github.com/jddcode/tictactoe/Structures"
	"strconv"
	"encoding/json"
)

var games Structures.Games
var players Structures.Players

	func main() {

		// Our players
		players = Structures.SetupPlayers()

		// Julienschmidt has a great library for mapping URL parameters in the format /x/y/ to usable parameters in
		// http handler functions
		router := httprouter.New()

		// Set up our array of games
		games = Structures.SetupGames()

		// Create a player (choose a 'cool' name)
		router.GET("/players/create/:name", CreatePlayer)

		// Either join a match, or join a newly created match
		router.GET("/games/join/:playerID", JoinAGame)

		// Make a move
		router.GET("/games/move/:playerID/:gameID/:xPos/:yPos", MakeAMove)

		// Get the status of the game
		router.GET("/games/status/:gameID", GetGameStatus)

		err := http.ListenAndServe(":8080",router)

		if(err != nil) {

			fmt.Println("Error: " + err.Error())
			os.Exit(0)
		}
	}

	// Create a new player and send back the player ID - use that to play your games
	func CreatePlayer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		name := p.ByName("name")

		if(len(name) < 2 || len(name) > 50) {

			// This is a bad request
			w.WriteHeader(400)
			w.Write([]byte("Please supply a player name of between two and 50 characters in the URL eg /players/create/john"))
			return
		}

		// Add this to the players' list
		player := players.Create(name)

		// This has gone OK
		w.WriteHeader(200)
		w.Write([]byte(player.Code))
	}

	// Join a game of tic tac toe. Either create a new game and register as waiting for a partner, or join
	// a game which is already waiting for a partner
	func JoinAGame(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// Get the player we are working with
		player, exists := players.Get(p.ByName("playerID"))

		if(!exists) {

			// This is a not found situation
			w.WriteHeader(404)
			w.Write([]byte("Player not found - try creating a player (/players/create/<name>"))
			return
		}

		// Join a game and send back the game ID
		gameID := games.JoinGame(player)

		// That's great, send back the game ID
		w.WriteHeader(200)
		w.Write([]byte(gameID))
	}

	func GetGameStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		status, outcome := games.GetStatus(p.ByName("gameID"))

		if(!outcome) {

			w.WriteHeader(404)
			w.Write([]byte("This game does not exist"))
			return
		}

		// Give a status update
		w.WriteHeader(200)
		statusJson, _ := json.Marshal(status)
		w.Write(statusJson)
	}

	func MakeAMove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// We need a playerID, gameID and x and y co-ordinates
		playerID := p.ByName("playerID")
		gameID := p.ByName("gameID")
		xPos, _ := strconv.Atoi(p.ByName("xPos"))
		yPos, _ := strconv.Atoi(p.ByName("yPos"))

		// Don't bother to check any of the four variables above becuase actually the various structs do this anyway

		// So move right into assuming those are right, and attempting the play
		outcome, err := games.MakeAMove(gameID, playerID, xPos, yPos)

		// Let the player know the outcome
		if(!outcome) {

			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		// Otherwise just a straight 200 OK will do
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}

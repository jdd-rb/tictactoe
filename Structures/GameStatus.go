package Structures

	// A summary struct for sending out to players on request.
	// Simplified player variables just have the player name, not code
	// The board is in there of course, plus some lazy booleans for the end user to quickly see what is going on
	// The real work of calculating game state is done in the Game struct - not this summary struct
	type GameStatus struct {

		PlayerOne, PlayerTwo string
		Board [3][3]Point
		Winner string
		WaitingForPlayer, Active, Won bool
	}

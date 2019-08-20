# Noughts And Crosses (Tic Tac Toe)

This is a tech test I was given which tests creation of APIs. The objective is to implement a noughts and crosses game.

This implementation uses a REST api and is fully multi thread safe, with locking to prevent crashes due to simultaneous read/write operations on maps where used.

Any number of players can register to play, and they can then join a game. Players will either be entered into a new game or paired up with another player waiting to play a game.

There is full validation to make sure that only players in the game can submit a move, and that it must be their turn when they submit a move. Also, once the game is won no more moves can be submitted.

It is also possible to get a report on the status of the game at any time during or after the game.

## Endpoints

This test contains the following endpoints:

/players/create/:name

This creates a new player with the given name. The success response is a UUID for the player, which should be given in subsequent API calls.

/games/join/:playerID

This takes the player UUID from above, and joins a game, either by creating a new game or by joining an existing one without two players already.

/games/move/:playerID/:gameID/:xPos/:yPos

This allows a player to make a move. There is automatic checking to make sure that the player ID given is in this specific game, and that it is their turn to move. Also there is checking that the game isn't over yet and that the square they want to occupy has not already been taken.

/games/status/:gameID

This allows you to check the status of any current game. Observers are welcome; you do not need to be playing in the game to check the status.

# Battleship

## Project description

To play Battleship, you need a game board where each player has a grid numbered from 1 to 10 horizontally and labeled from A to J vertically. Players must not be able to see their opponents' grids. Each player has:

- An aircraft carrier (5 squares)
  - A cruiser (4 squares)
  - A light cruiser (3 squares)
  - A submarine (3 squares)
  - A torpedo boat (2 squares)

Place your ships as you wish, without overlapping them. Once both players have placed all their ships, the game can begin.

## Functional rules

### Managing start game

- Each player must place their boats and validate them.
- Then the battle begins

#### Managing game sessions

- Each player, one after the other, enters a coordinate where they want to send a torpedo.

#### Saving data

Record:

- Each player's initial grid
- With each shot fired, the players' grids are updated, indicating whether or not it was a hit

## Models

``` go
Player {
  PlayerID
  Name string
  WinGames int
  LoseGames int
}
```

``` go
Session {
  SessionID
  PlayerIDA
  PlayerIDB
  Winner (playerID)
  StartDateTime 
  playerIDTour
}
```

``` go
Grid {
  GridID
  SessionID
  PlayerID
  BoatPosition[][]
  HitPosition[][]
}
```

## Ressources

### /players

``` go
Create
Update
Search
getbyID
getbyName
```

#### /sessions

``` go
Create => new game
```

#### /sessions/:id

``` go
Get
  => Params : idSession
  => Reponse : 
```

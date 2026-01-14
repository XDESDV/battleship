package models

import "time"

// Player represent a player
type Player struct {
	CustomID string `bson:"customID" json:"customId"`
	// ....
	CreatedAt time.Time `json:"creationDate"`
	Suspended bool      `json:"suspended"`
}

// Collection Mongodb collection
func (p *Player) Collection() string {
	return "player"
}

// PlayerInput
type PlayerInput struct {
	//.....
}

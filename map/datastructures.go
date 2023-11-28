package _map

import "hellavtor/player"

type Building struct {
	Floors []Floor
}

type Floor struct {
	Name string
	Number int
	Rooms []Room
	OnFloor bool

}

type Room struct {
	Name string
	Description string
	Items []Items
	InRoom bool
}


type Items struct {
	Weapons []player.Weapon
	Tools []player.Tool
}
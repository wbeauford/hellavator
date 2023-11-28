package _map

import (
	"fmt"
	"hellavtor/player"
)
var building Building
var currentFloor int
var currentRoom string
func CreateGameMap() *Building {
	var b Building
	b.Floors = createFloors()
	building = b
	return &b

}

func ListBuildingInfo(b *Building) {
	floors := b.Floors
	for _, x := range floors {
		fmt.Println(x.Name)
	}
}

func GetBuilding() *Building {
	return &building
}

func (b *Building) SetCurrentFloor(floorNumber int) {
	for i, _ := range b.Floors {
		if b.Floors[i].Number == floorNumber {
			if b.Floors[i].OnFloor {
				fmt.Println(fmt.Sprintf("You are already on the floor: %s", b.Floors[i].Name))
				return
			}
			b.Floors[i].OnFloor = true
		}
	}
}

func (b *Building) GetCurrentFloor() string {
	var currentFloor string
	for i, _ := range b.Floors {
		if b.Floors[i].OnFloor {
			currentFloor = b.Floors[i].Name
		}

	}

	return currentFloor
}

func (b *Building)ListRooms() {
	for i, _ := range b.Floors {
		if b.Floors[i].OnFloor {
			for _, x := range b.Floors[i].Rooms {
				fmt.Println(x.Name)
			}
		}
	}
}

func (b *Building)SetCurrentRoom(roomName string) {
	for i, _ := range b.Floors {
		if b.Floors[i].OnFloor {
			for x, _ := range b.Floors[i].Rooms {
				if b.Floors[i].Rooms[x].Name == roomName {
					b.Floors[i].Rooms[x].InRoom = true

				}
			}
		}
	}
}

func (b *Building)GetCurrentRoom() string {
	var currentRoom string
	for i, _ := range b.Floors {
		if b.Floors[i].OnFloor {
			for x, _ := range b.Floors[i].Rooms {
				if b.Floors[i].Rooms[x].InRoom {
					currentRoom = b.Floors[i].Rooms[x].Name
				}
			}
		}
	}

	return currentRoom
}

func createFloors() []Floor {
	var f []Floor
	f = append(f, createSingleFloor("LA Bank", 1, createRooms()))
	f = append(f, createSingleFloor("CTC Development", 2, createRooms()))
	f = append(f, createSingleFloor("Jitter Marketing", 3, createRooms()))
	f = append(f, createSingleFloor("Bryan Motors Co", 4, createRooms()))
	f = append(f, createSingleFloor("Roof", 5, nil))

	return f
}

func createRooms() []Room {
	var rs []Room
	var stockItems []Items
	var bathroomItems []Items
	var officeItems []Items
	var kitchemItems []Items

	stockItems = fillRoom("stockRoom")
	bathroomItems = fillRoom("bathroom")
	officeItems = fillRoom("office")
	kitchemItems = fillRoom("kitchen")

	rs = append(rs, createSingleRoom("Stock Room", "Office Supplies are stored here", stockItems))
	rs = append(rs, createSingleRoom("Bath Room", "A Bathroom", bathroomItems))
	rs = append(rs, createSingleRoom("Office Room", "A Cube Farm", officeItems))
	rs = append(rs, createSingleRoom("Kitchen", "A Kitchen", kitchemItems))

	return rs
}

func createSingleFloor(name string, number int, r []Room) Floor {
	var f Floor
	f.Name = name
	f.Number = number
	f.Rooms = r
	return f

}

func createSingleRoom(name string, description string, items []Items) Room {
	var r Room
	r.Name = name
	r.Description = description
	r.Items = items
	return r
}

func fillRoom(roomType string) []Items {
	var i Items
	var it []Items
	switch roomType {
	case "stockRoom":
		i.Tools = append(i.Tools, addToolsToRoom("drill", "Standard Electric Drill", false, "tool"))
		i.Tools = append(i.Tools, addToolsToRoom("Body Armor", "Up your health plus ten when using this", false, "health"))
	case "bathroom":
		i.Tools = append(i.Tools, addToolsToRoom("smokes", "Cigarettes - bad for your health but relaxing", true, "health"))
	case "office":
		i.Tools = append(i.Tools, addToolsToRoom("wire cutters", "Wire Cutters - used for cutting wires", false, "tool"))
		i.Weapons = append(i.Weapons, addWeaponsToRoom("p90", "p90 Sub Machine Gun", 15, "firearm"))
	case "kitchen":
		i.Weapons = append(i.Weapons, addWeaponsToRoom("Stun Gun", "A gun that fires electric current", 3, "firearm"))
		i.Tools = append(i.Tools, addToolsToRoom("Key Card", "A Key Card that gives you entry to secure areas", false, "tool"))
	}

	it = append(it, i)
	return it
}

func addToolsToRoom(name string, description string, recovery bool, itemtype string) player.Tool {
	var t player.Tool
	t.Name = name
	t.Description = description
	t.Type = itemtype
	t.Recovery = recovery
	return t
}

func addWeaponsToRoom(name string, description string, damage int, wtype string) player.Weapon {
	var w player.Weapon
	w.Name = name
	w.Description = description
	w.Damage = damage
	w.Type = wtype
	return w
}

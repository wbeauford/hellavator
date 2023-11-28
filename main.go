package main

import (
	"fmt"
	"hellavtor/character"
	_map "hellavtor/map"
	"hellavtor/player"
	"os"
	"strconv"
)



func main() {
	startScreen()
}



func startScreen() {
	fmt.Println("==================================================")
	fmt.Println("------------------ HELLAVATOR---------------------")
	fmt.Println("-------------A Game by William Beauford-----------")
	fmt.Println("==================================================")

	fmt.Println("Type: 'Start' to begin")
	var start string
	fmt.Scanln(&start)
	if start == "Start" {
		fmt.Println("Prepare to begin.......")
		playerCreateScreen()
	}
}


func playerCreateScreen() {
	var playerName string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&playerName)

	p := player.CreatePlayer(playerName)

	fmt.Println(fmt.Sprintf("Welcome %s %s, I hope you're ready for this.", p.Rank, p.Name))
	fmt.Println("Continue to game?")
	var continuegame string
	fmt.Scanln(&continuegame)
	fmt.Println(continuegame)
	if continuegame == "Y" || continuegame == "Yes" {
		GameStart(p)
	} else {
		fmt.Println("Goodbye...")
		os.Exit(1)
	}
}

func GameStart(p *player.Player) {
	var Harry character.Character
	var Mac character.Character
	Harry.Name = "Harry"
	Harry.Rank = "Sgt"
	Harry.Hero = true
	Harry.Age = 39
	Mac.Name = "MacMahon"
	Mac.Rank = "Cpt"
	Mac.Age = 47
	Mac.Hero = true
	var characters []character.Character
	characters = append(characters, Harry)
	characters = append(characters, Mac)
	character.SetCharacters(characters)
	building := _map.CreateGameMap()
	SelectionScreen(p, building)


}

func SelectionScreen(p *player.Player, building *_map.Building) {
	var selection string
	fmt.Println("Make a selection of what you want to do:")
	fmt.Println("1 List your Inventory")
	fmt.Println("2 List Your Status")
	fmt.Println("3 Call Someone")
	fmt.Println("4 View Building Info")
	fmt.Println("5 Move to Floor")
	fmt.Println("6 View Rooms")
	fmt.Println("7 Move to a Room")
	fmt.Println("8 List Items in Room")
	//fmt.Println("9 Add Item to Inventory")
	fmt.Scanln(&selection)
	num, _ := strconv.Atoi(selection)
	MakeSelection(num, p, building)
}

func MakeSelection(selection int, p *player.Player, building *_map.Building) {
	switch selection {
	case 1:
		fmt.Println("Which do you want to view?")
		fmt.Println("1: Weapons")
		fmt.Println("2: Tools")
		var iselect string
		fmt.Scanln(&iselect)
		MakeInventorySelection(iselect, p)
	case 2:
		fmt.Println(fmt.Sprintf("Name: %s", p.Name))
		fmt.Println(fmt.Sprintf("Rank: %s", p.Rank))
		fmt.Println(fmt.Sprintf("Health: %d", p.Health))
	case 3:
		player.DisplayRadioContacts()
	case 4:
		ListFloors(building)
	case 5:
		MoveToFloorScreen()
	case 6:
		ListRooms(building)
	case 7:
		MoveToRoomScreen()
	case 8:
		building := _map.GetBuilding()
		ListItemsInRoom(building)
	//case 9:
		//AddToInventoryScreen()

	}
}

func MoveToFloorScreen() {
	var floornum string
	fmt.Println("Enter Floor Number to move to")
	fmt.Scanln(&floornum)
	num, _ := strconv.Atoi(floornum)
	MoveToFloor(num ,_map.GetBuilding())
}

func MoveToRoomScreen() {
	var roomName string
	fmt.Println("Enter a Room Name to move to")
	fmt.Scanln(&roomName)
	b := _map.GetBuilding()
	MoveToRoom(roomName, b)
}

func MakeInventorySelection(selection string, p *player.Player) {
	num, _ := strconv.Atoi(selection)
	if num == 1 {
		player.OpenInventory(p, "Weapons")
	} else if num == 2 {
		player.OpenInventory(p, "Tools")
	} else {
		fmt.Println("Err: Invalid Selection")
	}
}


func ListFloors(b *_map.Building) {
	_map.ListBuildingInfo(b)
}

func MoveToFloor(floorNumber int, b *_map.Building) {
	currentFloor := GetCurrentFloor(b)
	floors := b.Floors
	for _, x := range floors {
		if currentFloor == x.Name {
			fmt.Println("You are already on this floor")
			continue
		}
		if floorNumber == x.Number {
			x.OnFloor = true
			fmt.Println(fmt.Sprintf("Moved to Floor: %s", x.Name))
			b.SetCurrentFloor(floorNumber)
		}
	}
	p := player.GetPlayer()
	SelectionScreen(p, b)
}

func GetCurrentFloor(b *_map.Building) string{
	return _map.GetCurrentFloor(b)

}

func ListRooms(b *_map.Building) {
	_map.ListRooms(b)
}

func MoveToRoom(roomName string, b *_map.Building) {
	currentRoom := GetCurrentRoom()
	floors := b.Floors
	for _, x := range floors {
		for _, y := range x.Rooms {
			if roomName == currentRoom {
				fmt.Println("You are already in this room")
				continue
			}
			if roomName == y.Name {
				y.InRoom = true
				fmt.Println(fmt.Sprintf("Moved inside Room : %s", y.Name))
				_map.SetCurrentRoom(roomName)
			}
		}
	}
}

func GetCurrentRoom() string {
	return _map.GetCurrentRoom()
}

func ListItemsInRoom(building *_map.Building) {
	var weapons []player.Weapon
	var tools []player.Tool
	weapons, tools = GetItemsInRoom(building)
	fmt.Println("Current Weapons in this room:")
	for _, x := range weapons {
		fmt.Println(x.Name)
	}

	fmt.Println("Current Items in this room")
	for _, y := range tools {
		fmt.Println(y.Name)
	}
}

func GetItemsInRoom(building *_map.Building) ([]player.Weapon, []player.Tool) {
	floors := building.Floors
	var weapons []player.Weapon
	var tools []player.Tool
	for _, x := range floors {
		if x.OnFloor {
			for _, y := range x.Rooms {
				if y.InRoom {
					for _, z := range y.Items {
						weapons = z.Weapons
						tools = z.Tools
					}
				}
			}
		}
	}
	return weapons, tools
}

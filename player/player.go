package player

import (
	"fmt"
	"hellavtor/character"
	"strconv"
)

var player Player

func CreatePlayer(name string) *Player {
	var p Player
	p.Name = name
	p.Health = 100
	p.Rank = "Officer"
	p.Inventory = createDefaultInventory()
	player = p
	return &p

}

func GetPlayerName() string {
	return player.Name
}

func GetPlayer() *Player {
	return &player
}

func createDefaultInventory() Inventory {
	var i Inventory
	i.Weapons = createDefaultWeapons()
	i.Tools = createDefaultTools()
	return i

}

func createDefaultWeapons() []Weapon {
	var w []Weapon
	var handGun Weapon
	var knife Weapon

	handGun.Name = "Rock Island 1911"
	handGun.Description = "Your favorite handgun, traded in your standard issue for this variant."
	handGun.Type = "firearm"
	handGun.Damage = 10

	knife.Name = "Tactical Knife"
	knife.Description = "A Knife you got during your military service. Very sharp."
	knife.Damage = 6
	knife.Type = "blade"

	w = append(w, handGun)
	w = append(w, knife)
	return w

}

func createDefaultTools() []Tool {
	var t []Tool
	var multiTool Tool
	var radio Tool

	multiTool.Name = "Multi Tool"
	multiTool.Description = "Standard multitool bought from a hardware store. Contains: small screwdriver, file, pliers."
	multiTool.Type = "tool"
	multiTool.Recovery = false

	radio.Name = "Radio"
	radio.Description = "Standard Police Radio. Used for communications"
	radio.Type = "communications"
	radio.Recovery = false

	t = append(t, multiTool)
	t = append(t, radio)
	return t

}

func AddWeaponToInventory(p *Player, w Weapon) *Player {
	currentInventory := p.Inventory
	currentInventory.Weapons = append(currentInventory.Weapons, w)
	p.Inventory = currentInventory
	return p

}

func AddToolToInventory(p *Player, t Tool) *Player {
	currentInventory := p.Inventory
	currentInventory.Tools = append(currentInventory.Tools, t)
	p.Inventory = currentInventory
	return p
}

func OpenInventory(p *Player, selection string) {
	inventory := p.Inventory
	if selection == "Weapons" {
		displayWeapons(inventory.Weapons)
	} else if selection == "Tools" {
		displayTools(p.Inventory.Tools)
	} else {
		fmt.Println("Err: Invalid selection")
	}
}

func displayWeapons(w []Weapon) {
	for _, x := range w {
		if x.Selected {
			fmt.Println(fmt.Sprintf("%s - %s - *", x.Name, x.Description))
		} else {
			fmt.Println(fmt.Sprintf("%s - %s", x.Name, x.Description))
		}
	}
}

func displayTools(t []Tool) {
	for _, x := range t {
		if x.Selected {
			fmt.Println(fmt.Sprintf("%s - %s - *", x.Name, x.Description))
		} else {
			fmt.Println(fmt.Sprintf("%s - %s", x.Name, x.Description))
		}
	}
}

func (p *Player)SelectWeapon(selection string) {
	for i, _ := range p.Inventory.Weapons {
		if p.Inventory.Weapons[i].Name == selection {
			p.Inventory.Weapons[i].Selected = true
		}
	}

}

func (p *Player)SelectTool(selection string) {
	for i, _ := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == selection {
			p.Inventory.Tools[i].Selected = true
		}
	}

}

func (p *Player)UseWeapon() {
	var selectedWeapon string
	for i, _ := range p.Inventory.Weapons {
		if p.Inventory.Weapons[i].Selected {
			selectedWeapon = p.Inventory.Weapons[i].Type
		}
	}
	switch selectedWeapon {
	case "firearm":
		fmt.Println("Bang, Bang, Bang")
	case "blade":
		fmt.Println("Slice")
	}

}

func (p *Player)UseTool {
	var selectedTool string
	var recovery bool
	var selectedType string
	for i, _ := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Selected {
			selectedTool = p.Inventory.Tools[i].Name
			recovery = p.Inventory.Tools[i].Recovery
			selectedType = p.Inventory.Tools[i].Type
		}
	}

	parseTool(p, selectedTool, recovery, selectedType)
}

func parseTool(p *Player, tool string, recovery bool, toolType string) {
	switch toolType {
	case "communications":
		DisplayRadioContacts()
	case "tool":
		fmt.Println(fmt.Sprintf("Used %s", tool))
	case "health":
		if recovery {
			p.Health = p.Health + 3
			fmt.Println(fmt.Sprintf("Used %s got some relief, current health: %d", tool, p.Health))
		} else {
			p.Health = p.Health +10
			fmt.Println(fmt.Sprintf("Equiped %s, current health: %d", tool, p.Health))

		}
		removeTool(p, tool)
	}
}

func removeTool(p *Player, tool string) {
	for i, _ := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == tool {
			p.Inventory.Tools = append(p.Inventory.Tools[:i], p.Inventory.Tools[i+1:]...)
		}
	}
}


func UseTool(p *Player) {
	tools := p.Inventory.Tools
	var selectedTool string
	var selectedToolName string
	var recovery bool
	for _, x := range tools {
		if x.Selected {
			selectedTool = x.Type
			selectedToolName = x.Name
			recovery = x.Recovery
		}
	}

	switch  selectedTool {
	case "communications":
		fmt.Println("Calling...")
	case "tool":
		fmt.Println(fmt.Sprintf("Used %s", selectedToolName))
	case "health":
		if recovery {
			p.Health = p.Health + 3
			fmt.Println(fmt.Sprintf("Got some health, current health: %d", p.Health))
		} else {
			p.Health = p.Health + 10
			fmt.Println(fmt.Sprintf("Equiped Body Army, current health: %d", p.Health))
		}
	}
}

func DisplayRadioContacts() {
	var selection string
	characters := character.GetCharacters()
	for _, x := range *characters {
		if x.Hero {
			fmt.Println(fmt.Sprintf("%s %s", x.Rank, x.Name))
		}
	}
	fmt.Println("Type the name of the person you want to call")
	fmt.Scanln(&selection)
	CallCharacter(selection)
}

func CallCharacter(name string) {
	characters := character.GetCharacters()
	for _, x := range *characters {
		if name == x.Name {
			Call(name)
		}
	}
}

func Call(name string) {
	switch name {
	case "MacMahon":
		CallMac()
	case "Harry":
		CallHarry()

	}
}

func CallMac() {
	var selection string
	var selectNum int
	playerName := GetPlayerName()
	fmt.Println(fmt.Sprintf("Talk to me %s", playerName))
	fmt.Println(fmt.Sprintf("%d Ask About Mission Paramaters ", 1))
	fmt.Println(fmt.Sprintf("%d Ask About Bombers location", 2))
	fmt.Println(fmt.Sprintf("%d Ask About The Money", 3))
	fmt.Scanln(&selection)
	selectNum, _ = strconv.Atoi(selection)
	MacSays(selectNum)


}

func MacSays(selection int) {
	switch selection {
	case 1:
		fmt.Println("The hostages should be stuck on an elevator on floor 3. Make contact and observe, but stay in a holding pattern")
	case 2:
		fmt.Println("We're working on that, I need you to be with the hostages")
	case 3:
		fmt.Println("We dont have it yet, and hes not talking to me. I need the god damn money now! Just stay where you are")
	}
}

func CallHarry() {
	var selection string
	var selectNum int
	playerName := GetPlayerName()
	fmt.Println(fmt.Sprintf("Whats up %s?", playerName))
	fmt.Println(fmt.Sprintf("%d Ask About Mission Paramaters ", 1))
	fmt.Println(fmt.Sprintf("%d Ask About the bomb", 2))
	fmt.Println(fmt.Sprintf("%d Ask About the hostages", 3))
	fmt.Scanln(&selection)
	selectNum, _ = strconv.Atoi(selection)
	HarrySays(selectNum)

}

func HarrySays(selection int) {
	switch selection {
	case 1:
		fmt.Println("We're in a holding pattern, we need to get to the hostages on floor 3 and observe the device")
	case 2:
		fmt.Println("Whoever made this, is a pro, looks like a collapsible circuit if we try to cut it the whole thing will fire")
	case 3:
		fmt.Println("We hold here for now, if we move the hostages, hell trigger the bomb")
	}
}


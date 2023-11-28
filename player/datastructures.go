package player

type Player struct {
	Name string
	Rank string
	Inventory Inventory
	Health int

}

type Inventory struct {
	Weapons []Weapon
	Tools []Tool
}

type Weapon struct {
	Name string
	Description string
	Damage int
	Type string
	Selected bool
}

type Tool struct {
	Name string
	Description string
	Type string
	Recovery bool
	Selected bool
}

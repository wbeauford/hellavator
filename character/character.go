package character


var currentCharacters []Character


func CreateCharacter(name string, age int, rank string, hero bool) *Character {
	var c Character
	c.Name = name
	c.Age = age
	c.Rank = rank
	c.Hero = hero
	return &c

}

func SetCharacters(characters []Character) *[]Character {
	currentCharacters = characters
	return &currentCharacters
}

func GetCharacters() *[]Character {
	return &currentCharacters
}



package game

type Card struct {
	ID           int
	Name         string
	Rarity       CardRarity
	Intelligence int
	Dexterity    int
	Strength     int
	Combat       int
	Luck         int
}

type CardRarity int

const (
	RarityBasic CardRarity = iota
	RarityRare
	RarityUltraRare
	RarityBlessing
)

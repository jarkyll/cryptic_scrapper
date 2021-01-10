package models

import "github.com/kamva/mgm/v3"

type Deck struct {
	mgm.DefaultModel `bson:",inline"`
	Format 			string `bson:"string"`
	Companion 		[]Card `bson:"companion"`
	Commander 		[]Card `bson:"commander"`
	Creatures 		[]Card `bson:"creatures"`
	Planeswalkers   []Card `bson:"planeswalkers"`
	Spells 			[]Card `bson:"spells"`
	Artifacts 		[]Card `bson:"artifacts"`
	Enchantments 	[]Card `bson:"enchantments"`
	Lands 			[]Card `bson:"lands"`
	Sideboard   	[]Card `bson:"sideboard"`
}


func (d *Deck) SetField(field string, value []Card) []Card {
	switch field {
	case "Commander":
		d.Commander = value
	case "Creatures":
		d.Creatures = value
	case "Spells":
		d.Spells = value
	case "Artifacts":
		d.Artifacts = value
	case "Enchantments":
		d.Enchantments = value
	case "Lands":
		d.Lands = value
	case "Sideboard":
		d.Sideboard = value
	case "Planeswalkers":
		d.Planeswalkers = value
	default:
		break
	}
	return value
}
//func (Deck d) setField(field string, value []Card) {
//
//}

type Card struct {
	Name 	string
	Count 	int
}

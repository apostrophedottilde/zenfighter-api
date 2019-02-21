package domain

import (
	"errors"
)

type Arena struct{}

func (arena *Arena) Fight(fighter1 Fighter, fighter2 Fighter) Fighter {
	if fighter1.GetPower() > fighter2.GetPower() {
		return fighter1
	} else if fighter1.GetPower() < fighter2.GetPower() {
		return fighter2
	} else if fighter1.GetPower() == fighter2.GetPower() {
		errors.New("Error - It's a tiebreak!")
	}
	return nil
}

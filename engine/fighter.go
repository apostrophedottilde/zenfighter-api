package engine

import (
	"errors"
	"fmt"

	"bitcrunchy.com/zenfighter-api/domain"
)

func (engine *arenaEngine) GetKnight(ID string) (*domain.Knight, error) {
	fighter := engine.knightRepository.Find(ID)
	if fighter == nil {
		return nil, errors.New(fmt.Sprintf("fighter with ID '%s' not found!", ID))
	}
	fmt.Println("fighter: ", fighter)
	return fighter, nil
}

func (engine *arenaEngine) ListKnights() []*domain.Knight {
	fighters := engine.knightRepository.FindAll()
	return fighters
}

func (arena *arenaEngine) Create(knight *domain.Knight) {
	arena.knightRepository.Save(&domain.Knight{
		Name:        knight.Name,
		Strength:    knight.Strength,
		WeaponPower: knight.WeaponPower,
	})
}

func (arena *arenaEngine) DeleteAll() {
	arena.knightRepository.DeleteAll()
}

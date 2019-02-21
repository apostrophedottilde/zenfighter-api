package engine

import (
	"bitcrunchy.com/zenfighter-api/domain"
)

type Engine interface {
	GetKnight(ID string) (*domain.Knight, error)
	ListKnights() []*domain.Knight
	Fight(fighter1ID string, fighter2ID string) domain.Fighter
	Create(knight *domain.Knight)
	DeleteAll()
}

type FighterRepository interface {
	Find(ID string) *domain.Knight
	FindAll() []*domain.Knight
	Save(knight *domain.Knight) int64
	DeleteAll()
}

type DatabaseProvider interface {
	GetKnightRepository() FighterRepository
}

type arenaEngine struct {
	arena            *domain.Arena
	knightRepository FighterRepository
}

func (engine *arenaEngine) Fight(fighter1 string, fighter2 string) domain.Fighter {
	f1 := engine.knightRepository.Find(fighter1)
	f2 := engine.knightRepository.Find(fighter2)
	resp := engine.arena.Fight(f1, f2)
	return resp
}

func NewEngine(db DatabaseProvider) *arenaEngine {
	return &arenaEngine{
		arena:            &domain.Arena{},
		knightRepository: db.GetKnightRepository(),
	}
}

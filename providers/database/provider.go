package database

import (
	"bitcrunchy.com/zenfighter-api/engine"
	_ "github.com/go-sql-driver/mysql"
)

type Provider struct {
}

func (provider *Provider) GetKnightRepository() engine.FighterRepository {

	return &knightRepository{}
}

func (provider *Provider) Close() {

}

func NewProvider() *Provider {
	return &Provider{}
}

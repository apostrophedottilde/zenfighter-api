package domain

import "testing"

func TestArena_Fight(t *testing.T) {
	orc1 := &orc{"1", 10}
	orc2 := &orc{"2", 20}

	arena := &Arena{}

	result := arena.Fight(orc1, orc2)

	if result == nil {
		t.Fatal("fight result is nil")
	}

	if result.GetID() != "2" {
		t.Error("fight result is not the strongest")
	}
}

func TestArena_Fight_Draw(t *testing.T) {
	orc1 := &orc{"1", 10}
	orc2 := &orc{"2", 10}

	arena := &Arena{}

	result := arena.Fight(orc1, orc2)

	if result != nil {
		t.Error("fight result is nil")
	}
}

type orc struct {
	id       string
	strength int
}

func (o *orc) GetID() string {
	return o.id
}

func (o *orc) GetPower() float64 {
	return float64(o.strength)
}

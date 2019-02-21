package database

import (
	"os"
	"strconv"
	"testing"

	_ "github.com/ory/dockertest"
	"bitcrunchy.com/zenfighter-api/domain"
)

var (
	kr knightRepository
)

func TestMain(m *testing.M) {
	kr = knightRepository{}
	code := m.Run()
	os.Exit(code)
}

func TestFind(t *testing.T) {
	kr.DeleteAll()

	k2 := domain.Knight{
		Name:        "newguy",
		Strength:    78,
		WeaponPower: 122,
	}

	savedID := kr.Save(&k2)

	data := strconv.FormatInt(savedID, 10)

	g := kr.Find(data)

	if g.Name != "newguy" {
		t.Error("Was expecting the names to match.")
	}
}

func TestFindBadRequest(t *testing.T) {
	kr.DeleteAll()

	k2 := domain.Knight{
		Name:        "",
		Strength:    0,
		WeaponPower: 122,
	}

	savedID := kr.Save(&k2)

	if savedID != 0 {
		t.Error("Was not expecting to get an ID back.")
	}
}

func TestFindNotFound(t *testing.T) {
	g := kr.Find("newguy")

	if g != nil {
		t.Error("Was expecting the knight to be not found.")
	}
}

func TestFindAll(t *testing.T) {
	kr.DeleteAll()
	all := kr.FindAll()

	if len(all) != 0 {
		t.Error("Was expecting zero knights.")
	}

	k1 := domain.Knight{
		Name:        "Tester-san",
		Strength:    9001,
		WeaponPower: 23435,
	}
	k2 := domain.Knight{
		Name:        "Jaberwocky",
		Strength:    22113344,
		WeaponPower: 888,
	}
	kr.Save(&k1)
	kr.Save(&k2)

	all = kr.FindAll()

	if len(all) != 2 {
		t.Error("Was not exactly 2 test knights.")
	}

	s1 := all[0]
	s2 := all[1]

	if s1.ID == s2.ID {
		t.Error("Was not expecting the 2 IDs to match.")
	}

	if s2.Name == "Tester-san" {
		t.Error("The name does not match.")
	}

	if s1.Strength != 9001 {
		t.Error("The strength does not match.")
	}

	if s1.WeaponPower != 23435 {
		t.Error("The WeaponPower does not match.")
	}

	if s2.Name != "Jaberwocky" {
		t.Error("The name does not match.")
	}

	if s2.Strength != 22113344 {
		t.Error("The strength does not match.")
	}

	if s2.WeaponPower != 888 {
		t.Error("The WeaponPower does not match.")
	}
}

func TestCreate(t *testing.T) {
	kr.DeleteAll()

	k := domain.Knight{
		Name:        "Tester-san",
		Strength:    9001,
		WeaponPower: 23435,
	}

	res := kr.Save(&k)

	if res == 0 {
		t.Error("Error saving.")
	}

	knights := kr.FindAll()

	if knights == nil {
		t.Error("Error fetching knights.")
	}

	k1 := knights[0]

	if k1.Name != "Tester-san" {
		t.Error("Name was not correct.")
	}

}

func TestCreateBadRequest(t *testing.T) {
	kr.DeleteAll()

	k := &domain.Knight{
		Name:        "Tester-san",
		Strength:    0,
		WeaponPower: 456,
	}

	res := kr.Save(k)

	if res != 0 {
		t.Error("Was not expecting to persist this due to a malformed request.")
	}

	k2 := &domain.Knight{
		Name:        "test1",
		Strength:    0,
		WeaponPower: 456,
	}

	res2 := kr.Save(k2)

	if res2 != 0 {
		t.Error("Was not expecting to persist this due to a malformed request.")
	}
}

func TestDeleteAll(t *testing.T) {
	kr.DeleteAll()

	k1 := domain.Knight{
		Name:        "Tester-san",
		Strength:    9001,
		WeaponPower: 23435,
	}

	k2 := domain.Knight{
		Name:        "Jaberwocky",
		Strength:    9001,
		WeaponPower: 23435,
	}

	kr.Save(&k1)
	kr.Save(&k2)

	all := kr.FindAll()

	if len(all) != 2 {
		t.Error("Was not exactly 2 test knights.")
	}

	kr.DeleteAll()

	all = kr.FindAll()

	if len(all) != 0 {
		t.Error("Was not zero knights in the database.")
	}

}

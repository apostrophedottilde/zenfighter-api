package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"bitcrunchy.com/zenfighter-api/domain"
)

var (
	dbName = os.Getenv("dbName")
	dbHost = os.Getenv("dbHost")
	dbPort = os.Getenv("dbPort")
	dbUser = os.Getenv("dbUser")
	dbPass = os.Getenv("dbPass")
	dbType = "mysql"
)

type knightRepository struct {
}

func (repo *knightRepository) Find(ID string) *domain.Knight {

	db, err := sql.Open(dbType, repo.connectURL())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	nid, _ := strconv.Atoi(ID)

	res, err := db.Query("SELECT id, name, strength, weaponpower FROM fighters WHERE id = ?", nid)

	id := 0
	name := ""
	strength := 0
	weaponPower := 0

	if !res.Next() {
		return nil
	}

	res.Scan(&id, &name, &strength, &weaponPower)

	mid := strconv.Itoa(id)

	db.Close()

	knt := repo.makeKnight(mid, name, strength, weaponPower)

	return knt
}

func (repo *knightRepository) FindAll() []*domain.Knight {
	db, err := sql.Open(dbType, repo.connectURL())

	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("DB: #v", db)

	res, err := db.Query("SELECT id, name, strength, weaponpower FROM fighters")

	if err != nil {
		fmt.Println("err: ", err)
	}

	if res != nil {
		fmt.Println("res: ", res)
	}

	var knights []*domain.Knight

	for res.Next() {
		id := 0
		name := ""
		strength := 0
		weaponPower := 0

		res.Scan(&id, &name, &strength, &weaponPower)

		nid := strconv.Itoa(id)

		k := &domain.Knight{
			ID:          nid,
			Name:        name,
			Strength:    strength,
			WeaponPower: weaponPower,
		}

		knights = append(knights, k)
	}

	db.Close()
	return knights
}

func (repo *knightRepository) Save(knight *domain.Knight) int64 {
	db, err := sql.Open(dbType, repo.connectURL())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if knight.Name == "" || knight.Strength == 0 || knight.WeaponPower == 0 {
		return 0
	}

	stmt, err := db.Prepare("INSERT fighters SET name=?,strength=?,weaponpower=?")

	res, err := stmt.Exec(knight.Name, knight.Strength, knight.WeaponPower)

	if err != nil {
		panic(err)
	}
	db.Close()

	id, errr := res.LastInsertId()

	if errr != nil {
		return 0
	} else {
		return id
	}
}

func (repo *knightRepository) DeleteAll() {
	db, err := sql.Open(dbType, repo.connectURL())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM fighters WHERE strength > -1")
	stmt.Exec()
	if err != nil {
		panic(err)
	}
	db.Close()
}

func (repo *knightRepository) makeKnight(id string, name string, strength int, weaponPower int) *domain.Knight {
	return &domain.Knight{
		ID:          id,
		Name:        name,
		Strength:    strength,
		WeaponPower: weaponPower,
	}
}

func (repo *knightRepository) connectURL() string {
	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

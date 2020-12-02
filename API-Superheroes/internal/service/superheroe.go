package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/liggerinileo/SeminarioGo/API-Superheroes/internal/config"
)

// Superheroe ...
type Superheroe struct {
	ID             int64
	Name           string
	Strength       int64
	Speed          int64
	FightingSkills int64
	Intelligence   int64
}

// SuperheroeService ...
type SuperheroeService interface {
	AddSuperheroe(Superheroe) (int64, error)
	FindByID(int) *Superheroe
	FindAll() []*Superheroe
	UpdateSuperheroe(Superheroe, int) (int64, error)
	DeleteByID(int)
}

type service struct {
	conf *config.Config
	db   *sqlx.DB
}

// NewSuperheroeService ...
func NewSuperheroeService(c *config.Config, db *sqlx.DB) (SuperheroeService, error) {
	return service{c, db}, nil
}

func (s service) AddSuperheroe(superheroe Superheroe) (int64, error) {
	exec := s.db.MustExec("INSERT INTO superheroe(name, strength, speed, fightingSkills, intelligence) VALUES ($1, $2, $3, $4, $5)", superheroe.Name, superheroe.Strength, superheroe.Speed, superheroe.FightingSkills, superheroe.Intelligence)
	return exec.LastInsertId()
}

func (s service) UpdateSuperheroe(superheroe Superheroe, id int) (int64, error) {
	exec := s.db.MustExec("UPDATE superheroe SET name = $1, strength = $2, speed = $3, fightingSkills = $4, intelligence = $5 WHERE ID = $4", superheroe.Name, superheroe.Strength, superheroe.Speed, superheroe.FightingSkills, superheroe.Intelligence, id)
	return exec.LastInsertId()
}

func (s service) FindByID(id int) *Superheroe {
	var superheroe Superheroe
	err := s.db.Get(&superheroe, "SELECT * FROM superheroe WHERE id=$1", int64(id))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &superheroe
}

func (s service) FindAll() []*Superheroe {
	var superheroes []*Superheroe
	err := s.db.Select(&superheroes, "SELECT * FROM superheroe")
	if err != nil {
		fmt.Println(err)
	}
	return superheroes
}

func (s service) DeleteByID(id int) {
	s.db.MustExec("DELETE FROM superheroes WHERE id = $1", id)
}

package crud

import "fmt"

// Superheroes ...
type Superheroes struct {
	superheroes map[int]*Superheroe
}

// Superheroe ...
type Superheroe struct {
	ID             int
	name           string
	strength       int
	speed          int
	fightingSkills int
	intelligence   int
}

// NewSuperheroe ...
func NewSuperheroe() Superheroes {
	return Superheroes{
		superheroes: make(map[int]*Superheroe),
	}
}

// Add new Superheroe
func (s Superheroes) Add(sh Superheroe) {
	s.superheroes[sh.ID] = &sh
}

// Print ...
func (s Superheroes) Print() {
	for _, v := range s.superheroes {
		fmt.Println(v.ID, v.name, v.strength, v.speed, v.fightingSkills, v.intelligence)
	}
}

// Update ...
func (s Superheroes) update(superheroe Superheroe) {
	s.superheroes[superheroe.ID] = &superheroe
}

// DeleteByID ...
func (s Superheroes) deleteByID(ID int) {
	delete(s.superheroes, ID)
}

// FindByID ...
func (s Superheroes) findByID(ID int) *Superheroe {
	return s.superheroes[ID]
}

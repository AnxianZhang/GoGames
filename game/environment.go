package game

import (
	"github.com/AnxianZhang/GoGames/entity/generic"
)

var _ generic.WorldView = Environment{}

type Environment struct {
	entities []generic.Entity
}

func NewEnvironment() *Environment {
	return &Environment{[]generic.Entity{}}
}

func (env Environment) SearchEntities(tag string) []generic.Entity {
	var result = make([]generic.Entity, 0, len(env.entities))

	for _, e := range env.entities {
		if e.Tag() == tag {
			result = append(result, e)
		}
	}

	return result
}

func (env Environment) GetEntites() []generic.Entity {
	return env.entities
}

func (env Environment) FindFirstEntity(tag string) (generic.Entity, bool) {
	for _, e := range env.entities {
		if e.Tag() == tag {
			return e, true
		}
	}

	return nil, false
}

func (env *Environment) AddEntity(other generic.Entity) *Environment {
	env.entities = append(env.entities, other)
	return env
}

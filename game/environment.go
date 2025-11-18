package game

import "github.com/AnxianZhang/GoGames/entity"

var _ entity.WorldView = Environment{}

type Environment struct {
	entities []entity.Entity
}

func NewEnvironment() *Environment {
	return &Environment{[]entity.Entity{}}
}

func (env Environment) SearchEntities(tag string) []entity.Entity {
	var result = make([]entity.Entity, 0, len(env.entities))

	for _, e := range env.entities {
		if e.Tag() == tag {
			result = append(result, e)
		}
	}

	return result
}

func (env Environment) GetEntites() []entity.Entity {
	return env.entities
}

func (env Environment) FindFirstEntity(tag string) (entity.Entity, bool) {
	for _, e := range env.entities {
		if e.Tag() == tag {
			return e, true
		}
	}

	return nil, false
}

func (env *Environment) AddEntity(other entity.Entity) *Environment {
	env.entities = append(env.entities, other)
	return env
}

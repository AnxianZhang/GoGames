package entity

type WorldView interface {
	SearchEntities(tag string) []Entity
}

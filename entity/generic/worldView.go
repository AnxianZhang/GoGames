package generic

type WorldView interface {
	SearchEntities(tag string) []Entity
	FindFirstEntity(tag string) (Entity, bool)
}

package inc

type Manager[T any] struct {
	nameOf   func(T) string
	entities []T
	index    map[string]int
}

func NewManager[T any](nameOf func(T) string) *Manager[T] {
	return &Manager[T]{
		nameOf: nameOf,
		index:  map[string]int{},
	}
}

func (m *Manager[T]) Register(entity T) {
	name := m.nameOf(entity)
	if i, ok := m.index[name]; ok {
		m.entities[i] = entity
		return
	}
	m.index[name] = len(m.entities)
	m.entities = append(m.entities, entity)
}

func (m *Manager[T]) GetByName(name string) (T, bool) {
	i, ok := m.index[name]
	if !ok {
		var missing T
		return missing, false
	}
	return m.entities[i], true
}

func (m *Manager[T]) GetAll() []T {
	return m.entities
}

func registerAll[T any](m *Manager[T], register func(T)) {
	for _, entity := range m.GetAll() {
		register(entity)
	}
}

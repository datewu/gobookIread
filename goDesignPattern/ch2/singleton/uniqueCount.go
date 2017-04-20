package singleton

// Singleton contract
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance return the Singleton interface
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne statisify the Singleton interface
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

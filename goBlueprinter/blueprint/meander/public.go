package meander

// Facade expose Public method
type Facade interface {
	Public() interface{}
}

// Public takes any object and return Facade or not
func Public(obj interface{}) interface{} {
	if p, ok := obj.(Facade); ok {
		return p.Public()
	}
	return obj
}

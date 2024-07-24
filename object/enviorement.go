package object

func NewEnclousedEnvironment(outer *Enviorment) *Enviorment {
	env := NewEnviorment()
	env.outer = outer
	return env
}
func NewEnviorment() *Enviorment {
	s := make(map[string]Object)
	return &Enviorment{store: s}
}

type Enviorment struct {
	store map[string]Object
	outer *Enviorment
}

func (e Enviorment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Enviorment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

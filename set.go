package dsal

//Implements an unordered set

type Set map[interface{}]bool

func NewSet() Set {
	return make(Set)
}

func (s Set) Length() int {
	return len(s)
}

func (s Set) Union(u Set) Set {
	x := make(Set)
	for k, _ := range u {
		x[k] = false
	}
	for k, _ := range s {
		x[k] = false
	}
	return x
}

func (s Set) Insert(v interface{}) {
	s[v] = false
}

func (s Set) Intersection(u Set) Set {
	x := make(Set)
	if s.Length() < u.Length() {
		//s is smaller in size
		for k, _ := range s {
			_, ok := u[k]
			if ok {
				x[k] = false
			}

		}
		return x
	} else {
		return u.Intersection(s)
	}
}

func (s Set) IsMember(v interface{}) bool {
	_, ok := s[v]
	return ok
}

//Remove, removes an item from the set
func (s Set) Remove(v interface{}) bool {
	l := s.Length()
	delete(s, v)
	return s.Length() < l
}

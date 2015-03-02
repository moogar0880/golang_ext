package golang_ext

import "py"

// Go's map object, wrapped as a Python class

type GoMap struct {
	AbstractObject
	o map[py.Object]py.Object
}

// TODO (jnappi) implement Python dict set/get attr methods for GoMap Type

// GoMap generator function that creates an allocated map for us on creation
func NewGoMap() *GoMap {
    return &GoMap{o: make(map[py.Object]py.Object)}
}

// Convenience function for returning a newly created GoMap instance
func Map() (py.Object, error) {
	return &NewGoMap()
}

func init() {
	methods := []py.Method{
		{"map", Map, "Thin wrapper around a Go mapping type"},
	}

	_, err := py.InitModule("_gotypes", methods)
	if err != nil {
		panic(err)
	}
}

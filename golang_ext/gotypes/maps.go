package gotypes

// #include "utils.h"
import "C"
import (
	"py"
)

// Go's map object, wrapped as a Python class
// GoMap is a Python object that wraps an abstract Go Map object
type GoMap struct {
	py.BaseObject
	py.MappingProtocol
	m map[py.Object]py.Object
}

// Go Utility function to create a new, empty, GoMap
func EmptyGoMap() (*GoMap, error){
	new_map, err := &GoMap{o: make(map[py.Object]py.Object)}
	if err {
		panic(err)
	}
	return new_map, nil
}

func NewGoMap(t *py.Type, args *py.Tuple, kw *py.Dict) (py.Object, error) {
	obj, err := t.Alloc(0)
	if err != nil {
		return nil, err
	}

	self, ok := obj.(*GoMap)
	if !ok {
		defer obj.Decref()
		return nil, py.TypeError.Err("Alloc returned wrong type: %T", obj)
	}

	// Replace with call to ParseTupleAndKeywords for creation with specified
	// values
	var buffer int

	err = py.ParseTuple(args, "i", &buffer)
	if err != nil {
		return nil, err
	}

	self.m = make(map[py.Object]py.Object)

	return self, nil
}

// PyDealloc is the deallocator for a GoMap instance, it is used internally -
// m.Decref() should be used as normal.
func (self *GoMap) PyDealloc() {
	self.m = nil
	self.Free()
}

// Provide getitem access from Python through to the GoMap instance. Note that
// in Go "default" is a keyword, so we can't use it as an argument
func (self *GoMap) Py___getitem__(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	var key py.Object
	kwlist := []string{"key"}

	err := ParseTupleAndKeywords(args, kw, "O", kwlist, &key)
	if err != nil {
		return nil, err
	}

	if val, ok := self.m[key]; ok {
		val.Incref()
		return &val, nil
	} else {
		py.raise(py.KeyError)
		return nil, nil
	}
}

// Provide setitem access from Python through to the GoMap instance. Return's
// Python's None instance
func (self *GoMap) Py___setitem__(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	var key, value py.Object
	kwlist := []string{"key", "value"}

	err := ParseTupleAndKeywords(args, kw, "O", kwlist, &key, &value)
	if err != nil {
		py.None.Incref()
		return py.None, err
	}

	value.Incref()
	self.m[key] <- value

	py.None.Incref()
	return py.None, nil
}

// Provide delitem access from Python through to the GoMap instance. Return's
// Python's None instance
func (self *GoMap) Py___delitem__(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	var key py.Object
	kwlist := []string{"key"}

	err := ParseTupleAndKeywords(args, kw, "O", kwlist, &key)
	if err != nil {
		return nil, err
	}

	if val, ok := self.m[key]; ok {
		val.Decref()
		delete(self.m, key)
		py.None.Incref()
		return py.None, nil
	} else {
		py.raise(py.KeyError)
		return nil, nil
	}
}

// Provide a clear method implementation which empties all key, value pairs
// from the GoMap instance
func (self *GoMap) Py_clear(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	for key, value := range self.m {
		value.Decref()
		delete(self.m, key)
	}
	py.None.Incref()
	return py.None, nil
}

// Copy a GoMap instance, and return a handle on the copy
func (self *GoMap) Py_copy(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	new_map, err := &EmptyGoMap()
	if err {
		return nil, err
	}
	for key, val := range self.m {
	  new_map[key] = val
	}
	return &new_map, nil
}

// Provide get(k, d) method from Python through to the GoMap instance
func (self *GoMap) Py_get(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	var key, default_value py.Object
	kwlist := []string{"key", "default"}

	err := ParseTupleAndKeywords(args, kw, "O", kwlist, &key, &default_value)
	if err != nil {
		py.None.Incref()
		return py.None, err
	}
	if val, ok := self.m[key]; ok {
		val.Incref()
		return val, nil
	} else {
		default_value.Incref()
		return default_val, nil
	}
}

// Provide get(k, d) method from Python through to the GoMap instance. Note that
// in Go "default" is a keyword, so we can't use it as an argument
func (self *GoMap) Py_pop(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	var key, default_value py.Object
	kwlist := []string{"key", "default"}

	err := ParseTupleAndKeywords(args, kw, "O", kwlist, &key, &default_value)
	if err != nil {
		py.None.Incref()
		return py.None, err
	}

	if val, ok := self.m[key]; ok {
		// If we found the key, remove it from the map and return it's value
		// We don't need to val.Decref() here, because were' both Incref-ing
		// and Decref-ing at the same time
		delete(self.m, key)
		return val, nil
	} else {
		// If we found nothing, return our provided default value
		return default_value, nil
	}
}

// Provide a clear method implementation which pops and returns a specified
// key, value pair as a 2-tuple. Will raise a KeyError if the GoMap instance
// is empty
func (self *GoMap) Py_popitem(args *py.Tuple, kw *py.Dict) (py.Object, error) {
	if len(self.m) == 0 {
		py.raise(py.KeyError)
		return nil, nil
	} else {
		var key, value py.Object
		kwlist := []string{"key", "value"}

		err := ParseTupleAndKeywords(args, kw, "O", kwlist, &key, &value)
		if val, ok := self.m[key]; ok {
			delete(self.m, key)
			new_tuple, err := py.PackTuple(key, val)
			if err != nil {
				py.None.Incref()
				return py.None, err
			}
			return new_tuple, nil
		}
	}
	// If m is not empty, but doesn't contain the key, value pair provided,
	// Raise a KeyError and return (py.None, nil)
	py.raise(py.KeyError)
	py.None.Incref()
	return py.None, nil
}

// Returns the map inside of a GoMap instance for ease of use within Go Code
func (self *GoMap) GoMap() map[py.Object]py.Object {
	return self.m
}

// TODO fromkeys :: Returns a new dict with keys from iterable and values equal to value.
// TODO items :: D.items() -> a set-like object providing a view on D's items
// TODO keys :: D.keys() -> a set-like object providing a view on D's keys
// TODO setdefault :: D.setdefault(k[,d]) -> D.get(k,d), also set D[k]=d if k not in D
// TODO update :: D.update([E, ]**F) -> None.  Update D from dict/iterable E and F. If E is present and has a .keys() method, then does:  for k in E: D[k] = E[k] If E is present and lacks a .keys() method, then does:  for k, v in E: D[k] = v In either case, this is followed by: for k in F:  D[k] = F[k]
// TODO values :: D.values() -> an object providing a view on D's values

// Docstring variables
var (
	GoMapDoc = "A native Go Mapping type that implements the same interface as a Python dict."
)

// Classtype Variables
var GoMapClass = py.Class {
	Name:    "_gotypes.GoMap",
	Pointer: (*GoMap)(nil),
	New:     EmptyGoMap,
	Doc:	 GoMapDoc,
}

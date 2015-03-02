package gotypes

import (
	"py"
	"sync"
)

var (
	ModuleLock		sync.Mutex
	GotypesModule  	*py.Module
)

var Version = "0.0.1.alpha"

func init() (*py.Module, error) {

	ModuleLock.Lock()
	defer ModuleLock.Unlock()

	if GotypesModule != nil {
		return GotypesModule, nil
	}

	GotypesModule, err := py.InitModule("_gotypes", methods)
	if err != nil {
		return nil, err
	}

	GotypesModule.AddStringConstant("__version__", Version)

	m, err := GoMapClass.Create()
	if err != nil {
		GotypesModule.Decref()
		GotypesModule = nil
		return nil, err
	}

	err = GotypesModule.AddObject("GoMap", m)
	if err != nil {
		GotypesModule.Decref()
		GotypesModule = nil
		m.Decref()
		return nil, err
	}

	return GotypesModule, nil
}

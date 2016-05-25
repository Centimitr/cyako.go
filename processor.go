package cyako

import (
	// "fmt"
	"errors"
	"reflect"
	"strings"
)

/*
	Processor & Processor Module
*/

type Processor struct {
	Module  string
	PkgPath string
	Name    string
	Func    func(*Ctx) []reflect.Value
}

// Client use a part of the table's key (usually is ProcessorName) to match a processor.
// Currently, the key is "PkgName.ProcessorName" so team members should use different package names to do a replacement though their package path is not the same.
func (c *Cyako) matchProcessor(suffix string) (func(*Ctx) []reflect.Value, error) {
	matchedList := []string{}
	// key is "PkgName.ProcessorName"
	for key, _ := range c.ProcessorMap {
		if strings.HasSuffix(key, suffix) {
			matchedList = append(matchedList, key)
		}
	}
	if len(matchedList) == 1 {
		return c.ProcessorMap[matchedList[0]].Func, nil
	}
	return nil, errors.New("Cannot select 1 processor.")
}

// module packages use LoadModule to load itselft
func (c *Cyako) loadModule(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	for i := 0; i < v.NumMethod(); i++ {
		index := i
		c.ProcessorMap[t.Name()+"."+t.Method(i).Name] = &Processor{
			PkgPath: t.PkgPath(),
			Module:  t.Name(),
			Name:    t.Method(i).Name,
			Func: func(ctx *Ctx) []reflect.Value {
				return t.Method(index).Func.Call([]reflect.Value{v, reflect.ValueOf(ctx)})
			},
		}
	}
}

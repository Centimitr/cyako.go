package cache

import (
	cyako "github.com/Cyako/Cyako.go"
)

type Cache struct{}

func (c Cache) BeforeProcess(ctx *cyako.Ctx) {

}

func (c Cache) AfterProcess(ctx *cyako.Ctx) {

}

func init() {
	cyako.LoadMiddleware(Cache{})
}

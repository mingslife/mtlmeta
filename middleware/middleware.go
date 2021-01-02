package middleware

import (
	"github.com/mingslife/mtlmeta/conf"
	"github.com/mingslife/mtlmeta/mtl"
)

type MTLMiddleware interface {
	Handle(*mtl.MTLFile) bool
}

var middlewares []MTLMiddleware

func RegisterMiddlewares(c *conf.Config) {
	middlewares = append(middlewares, NewDictionaryMiddleware(c))
}

func Handle(mtlFile *mtl.MTLFile) {
	for _, middleware := range middlewares {
		intercept := middleware.Handle(mtlFile)
		if intercept {
			break
		}
	}
}

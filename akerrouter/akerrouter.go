package akerrouter

import (
	"github.com/kevinjiaxu96/Aker/aker"
)

type router struct {
	getRequest    map[string]func(aker.Contex)
	postRequest   map[string]func(aker.Contex)
	putRequest    map[string]func(aker.Contex)
	deleteRequest map[string]func(aker.Contex)
}

func Instance() router {
	tmp := router{}
	return tmp
}

func (r *router) Get(path string, callback func(aker.Contex)) {
	if r.getRequest == nil {
		r.getRequest = make(map[string]func(aker.Contex))
	}
	r.getRequest[path] = callback
}

func (r *router) Post(path string, callback func(aker.Contex)) {
	if r.postRequest == nil {
		r.postRequest = make(map[string]func(aker.Contex))
	}
	r.postRequest[path] = callback
}

func (r *router) Put(path string, callback func(aker.Contex)) {
	if r.putRequest == nil {
		r.putRequest = make(map[string]func(aker.Contex))
	}
	r.putRequest[path] = callback
}

func (r *router) Delete(path string, callback func(aker.Contex)) {
	if r.deleteRequest == nil {
		r.deleteRequest = make(map[string]func(aker.Contex))
	}
	r.deleteRequest[path] = callback
}

func (r router) Routes(ctx aker.Contex, next func()) {
	if ctx.Request.Method == "GET" {
		tmp, exist := r.getRequest[ctx.Request.URL.Path]
		if exist == true {
			tmp(ctx)
		}
	} else if ctx.Request.Method == "POST" {
		tmp, exist := r.postRequest[ctx.Request.URL.Path]
		if exist == true {
			tmp(ctx)
		}
	} else if ctx.Request.Method == "PUT" {
		tmp, exist := r.putRequest[ctx.Request.URL.Path]
		if exist == true {
			tmp(ctx)
		}
	} else if ctx.Request.Method == "DELETE" {
		tmp, exist := r.deleteRequest[ctx.Request.URL.Path]
		if exist == true {
			tmp(ctx)
		}
	}
	next()
}

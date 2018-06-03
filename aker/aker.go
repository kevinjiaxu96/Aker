// Copyright 2018 Jiawei Xu. All rights reserved.
// MIT Liscense

// Package aker -> Simple package to chain middlewares for web server
package aker

import (
	"net/http"
	"strconv"
)

// Contex -> Context structure holding writer and request
type Contex struct {
	Response http.ResponseWriter // Sugar for ResponseWriter
	Request  *http.Request       // Sugar for Request
}

// Middlewares -> Middleware structrue holding the chain of middlewares
type Middlewares struct {
	chain []interface{} // Array of middlewares to chain
	index int           // index of middleware
	ctx   Contex
}

// Instance -> function which to create a new instance of middleware chaining
func Instance() Middlewares {
	myApp := Middlewares{} //new instance
	return myApp
}

// Use -> add middleware to the chain
func (m *Middlewares) Use(callback func(ctx Contex, next func())) {
	if callback != nil { //check if middleware is nil
		m.chain = append(m.chain, callback)
	}
}

// Callback -> function which toreturn a handler to be used
// by user instead of calling Listen
func (m *Middlewares) Callback() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(len(m.chain))
		if len(m.chain) > 0 {
			m.ctx = Contex{Response: w, Request: r}
			m.next()
			// firstMiddleware(Contex{ //invoke the first middleware
			// 	Response: w,
			// 	Request:  r,
			// }, m.next) // next is a function to next middleware
		}
	})
}

// next -> function to call the next middleware
func (m *Middlewares) next() {
	if m.index < (len(m.chain)) { //check if the next middleware exist
		nextMiddleware := m.chain[m.index].(func(ctx Contex, next func())) // cast empty interface back to function
		m.index = m.index + 1
		nextMiddleware(m.ctx, m.next) //invoke next middleware
	} else {
		m.index = 0 //if this is the final call to next, reset back to first
	}
}

// Listen -> function to start a http server
// and invoke callback through goroutine
func (m Middlewares) Listen(PORT int, CALLBACK func()) {
	desiredPort := ":" + strconv.Itoa(PORT) //cast the desired port
	if CALLBACK != nil {                    //if callback is supplied, invoke callback
		go CALLBACK()
	}
	if err := http.ListenAndServe(desiredPort, m.Callback()); err != nil { //start server, if fail panic
		panic(err)
	}
}

// JSON -> function to start a http server
// and invoke callback through goroutine
func (ctx Contex) JSON(data []byte) {
	ctx.Response.Header().Set("Content-Type", "application/json")
	ctx.Response.Write(data)
}

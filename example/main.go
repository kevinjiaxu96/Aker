// Copyright 2018 Jiawei Xu. All rights reserved.
// MIT Liscense

// package main -> testing code for aker
package main

import (
	"fmt"
	"net/http"

	"github.com/kevinjiaxu96/Aker/aker"
	"github.com/kevinjiaxu96/Aker/akerrouter"
)

// A -> first test middleware
func A(ctx aker.Contex, next func()) {
	fmt.Println("Middle A")
	next() //invoke next middleware
	fmt.Println("Middle A End")
}

// B -> second test middleware
func B(ctx aker.Contex, next func()) {
	fmt.Println("Middle B")
	next() //invoke next middleware
	fmt.Println("Middle B End")
}

// C -> third test middleware
func C(ctx aker.Contex, next func()) {
	fmt.Println("Middle C")
	next()                  //invoke next middleware
	testResponse := "hello" //just writing a simple response to client
	ctx.Response.Write([]byte(testResponse))
	fmt.Println("Middle C End")
}

// myOwnHandler -> my own handler middleware to invoke
func myOwnHandlerTestA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MyOwnHandlerTestA")
		next.ServeHTTP(w, r) //call next middleware
		fmt.Println("MyOwnHandlerTestA End")
	})
}

// myOwnHandler -> my own handler to invoke and also invoke aker
func myOwnHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MyOwnHandler")
		next.ServeHTTP(w, r) //call next middleware
		fmt.Println("MyOwnHandler End")
	})
}

// randomPrint -> random callback when starting server
func randomPrint() {
	fmt.Println("Starting Server!")
}

func main() {
	test := aker.Instance() //get a fresh instance of aker
	router := akerrouter.Instance()
	router.Get("/api", func(aker.Contex) {
		fmt.Println("You're good!")
	})
	router.Get("/api", func(aker.Contex) {
		fmt.Println("You're replaced!")
	})
	router.Get("/api/butts", func(ctx aker.Contex) {
		fmt.Println("Nice butts!")
	})
	test.Use(A)
	test.Use(B)
	test.Use(C)
	test.Use(router.Routes)
	test.Listen(8080, nil)
}

# Aker
Just a simple Golang middleware chaining package and router middleware package

> p.s syntax and signatures are mostly inspired by koa2 from nodejs

## Contex
```
type Contex struct {
	Response http.ResponseWriter // Sugar for ResponseWriter
	Request  *http.Request       // Sugar for Request
}
```

## How to write a Aker Middleware
```
func Foo(ctx aker.Contex, next func(aker.Contex)) {
	//middleware logic
	next() //invoke next middleware, must be called or chain will stop here
}
```


## Examples of router middle ware
```
func testAdd(ctx aker.Contex) {
	//post logic
}

app := aker.Instance()
router := akerrouter.Instance()
router.Get("/api/testing", func(ctx aker.Context){
	//get logic
})
router.Post("/api/add", testAdd)
app.Use(router.Routes)
```

## Examples of middleware chaining
```
Middleware A
Middleware B
Middleware C
Middleware C Ends
Middleware B Ends
Middleware A Ends
```

### Running with Aker's own listen function

```
Aker.Listen starts a http server

Function Signature: Listen(PORT, CALLBACK)

CALLBACK is ran using goroutine
```

## Author
* <a href="https://github.com/kevinjiaxu96">Kevin Xu</a>

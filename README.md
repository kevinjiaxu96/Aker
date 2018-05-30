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
	next(ctx) //invoke next middleware, must be called or chain will stop here
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

### Running with Aker's own listen function

```
Aker.Listen starts a http server

Function Signature: Listen(PORT, CALLBACK)

CALLBACK is ran using goroutine
```

<img width="460" alt="middlecode" src="https://user-images.githubusercontent.com/26973140/40585084-4efc99a0-6161-11e8-976c-2026033f17c9.png">

<img width="173" alt="middle" src="https://user-images.githubusercontent.com/26973140/40585081-483c2acc-6161-11e8-9862-eb9e7833f562.png">


### Running with User's handler function
```
Aker.Callback returns a http.handlerfunc
```

<img width="176" alt="myown" src="https://user-images.githubusercontent.com/26973140/40585110-d28e9ce6-6161-11e8-8c4b-f34a5a3ea906.png">


<img width="977" alt="myowncode" src="https://user-images.githubusercontent.com/26973140/40585129-030de994-6162-11e8-8a68-7b901fa304ee.png">

## Author
* <a href="https://github.com/kevinjiaxu96">Kevin Xu</a>

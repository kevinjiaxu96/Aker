## Methods

### `aker.Instance()`
Returns `Object`:
* `chain` - Holds the current list of middlewares
* `index` - Holds the current index of middleware executing
* `ctx` - Holds the context/ResonseWriter/Request of the middleware

### `aker.Use(middleware func(ctx aker.Contex, next func()))`
Appends the current middleware to the list of middlewares.

### `aker.Callback()`
Returns `Function`:
```
http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)
```

### `aker.Listen(PORT int, callback func())`
Starts http server listening on port 'PORT' and runs the callback function using Goroutine.

### `aker.JSON(data []byte)`
Writes the data as JSON back to the requested client.
```
Basically does http.ResponseWriter.Write()
```
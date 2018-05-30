## Methods
> Currently only supports GET, PUT, POST, DELETE.
### `akerrouter.Instance()`
Returns `Object`:
* `getRequests` - Holds the current list of get callbacks
* `postRequests` - Holds the current list of post callbacks
* `putRequests` - Holds the current list of put callbacks
* `deleteRequests` - Holds the current list delete get callbacks

### `akerrouter.Routes(ctx aker.Contex, next func())`
A middleware to be used with Aker.

### `akerrouter.Get(PATH string, Callback func(ctx aker.Contex))`
Appends the current callback to the list of get callbacks.

### `akerrouter.Post(PATH string, Callback func(ctx aker.Contex))`
Appends the current callback to the list of post callbacks.

### `akerrouter.Put(PATH string, Callback func(ctx aker.Contex))`
Appends the current callback to the list of put callbacks.

### `akerrouter.Delete(PATH string, Callback func(ctx aker.Contex))`
Appends the current callback to the list of delete callbacks.


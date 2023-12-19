# CacheKit

CacheKit is a thread-safe cache implementation in Go with support for generic key-value types.

## Features

- Thread-safe operations
- Generic key-value types
- Basic cache operations: Get, Set, Update, Delete
- Additional utility functions: Len, Clear, Keys, Values, Items, Has

## Usage

First, get the package:

```go
go get "github.com/velicanercan/cachekit"
```

Import the package:
```go
import "cachekit"
```
Create a new cache:
```
cache := cachekit.New[int, string]()
```
Set a value:
```
cache.Set(1, "value")
```
Get a value:
```
value, ok := cache.Get(1)
```
Update a value:
```
err := cache.Update(1, "new value")
```

Delete a value:
```
cache.Delete(1)
```
Get the length of the cache:
```
length := cache.Len()
```
Clear the cache:
```
cache.Clear()
```

Get all keys:
```
keys := cache.Keys()
```

Get all values:
```
values := cache.Values()
```

Get all items:
```
items := cache.Items()
```

Check if a key exists:
```
exists := cache.Has(1)

```
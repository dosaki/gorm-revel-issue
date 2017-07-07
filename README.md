# Issue with Revel and Gorm
Not sure if this is an actual problem or me doing something stupid.

## Problem
`gorm` throws an error when trying to call `db.Create()` from within a controller (see `app/controllers/app.go`).

The database initializing function is inside `app/datasource/datasource.go` and gets called in `app/init.go` with

```go
revel.OnAppStart(datasource.AutoInitDB)
```

## How to reproduce:
### Requirements
* A mysql database
* [revel](https://github.com/revel/revel/)
* [gorm](https://github.com/jinzhu/gorm)

### Steps
1. Start the app with `revel run gorm_issue`
2. Go to `http://localhost:9000/`
3. See error :(

```
Runtime Panic
runtime error: invalid memory address or nil pointer dereference

In /github.com/jinzhu/gorm/main.go (around line 407)

402: }
403:
404: // Create insert the value into database
405: func (s *DB) Create(value interface{}) *DB {
406:   scope := s.clone().NewScope(value)
407:   return scope.callCallbacks(s.parent.callbacks.creates).db
408: }
409:
410: // Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
411: func (s *DB) Delete(value interface{}, where ...interface{}) *DB {
```

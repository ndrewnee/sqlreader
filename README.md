# SqlReader

[![Build Status](https://travis-ci.org/ndrewnee/sqlreader.svg?branch=master)](https://travis-ci.org/ndrewnee/sqlreader)

Are writing like this:
```go
err := db.Exec("SELECT * FROM some_table")
```
or like this?
```go
const FleetUpdateStatus = `
  UPDATE fleets
  SET status = 3
  WHERE id = ?
`

err := db.Exec(FleetUpdateStatus)
```
I don't like writing sql statements in code. It has many disadvantages: no syntax-highliting, no auto-complete, etc.
This little library helps you to use *.sql files instead of const strings

## Install
```bash
go get "github.com/ndrewnee/sqlreader"
```

## Usage

Import the package:

```go
import (
	"github.com/ndrewnee/sqlreader"
)

```

## Example


```go
// First param is path with sqls
// Then names of sqls that are required (Optional)
sqls, err := sqlreader.New("path-with-sqls", "required_sql")
if err != nil {
  panic(err)
}

sql := sqls.Get("required_sql")
db.Exec(sql)

```

For more examples have a look at sqlreader_test.go

Running tests:

```bash
go test "github.com/ndrewnee/sqlreader"
```

## License 
MIT (see [LICENSE](https://github.com/ndrewnee/sqlreader/blob/master/LICENSE) file)

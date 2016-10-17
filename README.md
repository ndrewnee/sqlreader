# sqlreader
Are writing like this:
```go
err := db.Exec('SELECT * FROM some_table')
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

## Usage

Import the package:

```go
import (
	"github.com/ndrewnee/sqlreader"
)

```

```bash
go get "github.com/ndrewnee/sqlreader"
```

## Example


```go
sqls, err := sqlreader.New("path-with-sqls")
if err != nil {
  panic(err)
}

sql := sqls.Get("some_sql")
db.Exec(sql)

```

For more examples have a look at sqlreader_test.go

Running tests:

```bash
go test "github.com/ndrewnee/sqlreader"
```

## License 
MIT (see [LICENSE](https://github.com/ndrewnee/sqlreader/blob/master/LICENSE) file)

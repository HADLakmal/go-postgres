# GO Postgres #

go-postgres library provide NoSQL functionality which can execute queries with pool of connections.

### What is this repository for? ###

* Establish postgres connection pool
* Execute the SQL queries

### How do I get set up? ###

#### Establish Postgres Connection

```go
package main

import (
	"fmt"
	"bitbucket.org/mybudget-dev/go-postgres"
)

func main() {
	conf := gopostgres.DBConfig{
		User:     "User",
		Password: "Password",
		Host:     "<<Host>>",
		Port:     "<<PORT>>",
		Database: "Database",
	}
	DBPool, err := gopostgres.NewDBPool(conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer DBPool.Close()

}
```

#### Postgres Insert #####

```go
package main

import (
	"context"
	"fmt"
	"bitbucket.org/mybudget-dev/go-postgres"
)

func main() {
	conf := gopostgres.DBConfig{
		User:     "User",
		Password: "Password",
		Host:     "<<Host>>",
		Port:     "<<PORT>>",
		Database: "Database",
	}
	sql, err := gopostgres.NewDBPool(conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sql.Close()

	// context 
	ctx := context.Background()

	// insert data into postgres
	data := "data"
	// Transaction opt is optional and you can configure DB isolation levels  
	transaction, err := sql.BeginTransaction(ctx, nil)
	if err != nil {
		return
	}
	userQuery :=
		fmt.Sprintf(`INSERT INTO "%s"(data) VALUES($1)`,
			"table")
	statement, err := transaction.PrepareContext(ctx, userQuery)
	if statement != nil {
		defer statement.Close()
	} else {
		fmt.Println(err.Error())
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = statement.ExecContext(ctx, &data)
	if err != nil {
		if errorTransaction := transaction.Rollback(); errorTransaction != nil {
			fmt.Println(err.Error())
		}
	}
}
```

#### Postgres Get #####

````go
package main

import (
	"context"
	"fmt"
	"bitbucket.org/mybudget-dev/go-postgres"
)

func main() {
	conf := gopostgres.DBConfig{
		User:     "User",
		Password: "Password",
		Host:     "<<Host>>",
		Port:     "<<PORT>>",
		Database: "Database",
	}
	sql, err := gopostgres.NewDBPool(conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sql.Close()

	// context 
	ctx := context.Background()

	// get data from postgres
	key := 1
	query := fmt.Sprintf(`SELECT data FROM "%s" WHERE key=$1  ;`, "table")
	statementOfLocation, err := sql.Prepare(ctx, query)
	if statementOfLocation != nil {
		defer statementOfLocation.Close()
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	row, err := statementOfLocation.QueryContext(ctx, key)
	if row != nil {
		defer row.Close()
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data interface{}
	for row.Next() {
		err = row.Scan(&data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
````
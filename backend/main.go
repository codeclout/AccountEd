package main

import (
	"log"

	account2 "github.com/codeclout/AccountEd/adapters/app/api/account"
	"github.com/codeclout/AccountEd/adapters/core/account"
	"github.com/codeclout/AccountEd/adapters/framework/in/http"
	"github.com/codeclout/AccountEd/adapters/framework/out/db"
	ports2 "github.com/codeclout/AccountEd/ports/app"
	ports3 "github.com/codeclout/AccountEd/ports/core"
	ports4 "github.com/codeclout/AccountEd/ports/framework/in/http"
	ports "github.com/codeclout/AccountEd/ports/framework/out/db"
)

func main() {
	var (
		e error

		accountDbAdapter ports.AccountDbPort
		accountAdapter   ports3.AccountPort
		accountAPI       ports2.AccountAPIPort
		httpAdapter      ports4.HTTPPort
	)

	// uri := "mongodb://localhost:27017,localhost:27018,localhost:27019/accountEd?replicaSet=rs0"
	uri := "mongodb://db,db1,db2/accountEd?replicaSet=rs0"
	accountDbAdapter, e = db.NewAdapter(5, uri)

	defer accountDbAdapter.CloseConnection()

	if e != nil {
		log.Fatalf("Failed to instantiate db connection: %v", e)
	}

	accountAdapter = account.NewAdapter()
	accountAPI = account2.NewAdapter(accountAdapter, accountDbAdapter)
	httpAdapter = http.NewAdapter(accountAPI)

	httpAdapter.Run()

}

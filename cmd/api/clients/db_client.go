package clients

import (
	"database/sql"
	"fmt"
	"github.com/lmurature/api-boilerplate/cmd/api/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbClient struct {
	Client *sql.DB
}

func NewDbClient() *DbClient {
	database := &DbClient{}
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.GetDbUser(),
		config.GetDbPass(),
		config.GetDbHost(),
		config.GetDbName())
	database.Client, err = sql.Open("mysql", url)
	if err != nil {
		fmt.Println("LUCIO")
		panic(err)
	}
	database.Client.SetConnMaxLifetime(time.Minute * 3)
	database.Client.SetMaxOpenConns(10)
	database.Client.SetMaxIdleConns(10)
	return database
}

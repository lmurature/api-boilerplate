package clients

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbClient struct {
	Client *sql.DB
}

func NewDbClient() *DbClient {
	database := &DbClient{}
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "root", "localhost", "users_db")
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

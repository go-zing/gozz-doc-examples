package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-zing/gozz-kit/ztree"

	"github.com/go-zing/gozz-doc-examples/overview02"
)

func main() {
	app, cf, err := overview02.Initialize_Application(&overview02.Config{
		Server: overview02.ServerConfig{
			Addr: ":8080",
		},
		Sql: overview02.SqlConfig{
			Dsn: fmt.Sprintf("%s@%s:tcp(localhost:3306)/", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD")),
		},
		Redis: overview02.RedisConfig{
			Host: "localhost",
			Port: "6379",
		},
	})
	if err != nil {
		panic(err)
	}
	defer cf()
	fmt.Printf("%s", ztree.Draw("overview02", ztree.Parse(app)))
}

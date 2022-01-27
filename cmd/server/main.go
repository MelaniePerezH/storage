package main

import (
	"database/sql"

	"github.com/extlurosell/meli_bootcamp_go_w3-2/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/storage")
	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}

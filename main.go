package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pdrum/todo/handler"
	"github.com/pdrum/todo/model"
	"github.com/pdrum/todo/service"

	"github.com/labstack/echo"
)

func main() {
	pass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	url := fmt.Sprintf("host=%s port=5432 user=todo dbname=todo password=%s sslmode=disable", dbHost, pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	res := handler.ItemHandler{
		Service: service.ItemService{
			Repo: &model.SQLItemRepo{
				DB: db,
			},
		},
	}
	e.GET("/", res.ListAll)
	e.POST("/", res.Create)
	e.Logger.Fatal(e.Start(":1323"))
}

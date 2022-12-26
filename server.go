package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yoteerakrit/assessment/app/api"
	"github.com/yoteerakrit/assessment/repo.go"
	"github.com/yoteerakrit/assessment/usecase"
)

func main() {
	// fmt.Println("Please use server.go for main file")
	// fmt.Println("start at port:", os.Getenv("PORT"))

	db := repo.NewPostgreSQL()

	expenseRepo := repo.NewRepo(db)
	uc := usecase.New(expenseRepo)
	hdl := api.NewHandler(uc)

	e := api.New(hdl)

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

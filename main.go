package main

import (
	"github.com/srinivasrdy247/Employee_assignment_go/store"
	"github.com/srinivasrdy247/Employee_assignment_go/server"
	"github.com/srinivasrdy247/Employee_assignment_go/db"
	"github.com/srinivasrdy247/Employee_assignment_go/handler"
	//"github.com/srinivasrdy247/Employee_assignment_go/model"
)

func main(){

	r := router.New()
	
	v1 := r.Group("/api")

	connection := db.Connect()
	db.CreateNew(connection)
	functions := store.NewCon(connection) //store.NewCon(*sqlx.db)
	h := handler.NewHandler(functions)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8587"))	
}
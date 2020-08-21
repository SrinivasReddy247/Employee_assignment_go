package db

import (
  "fmt"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
  )
  
  var schema=`
  CREATE TABLE IF NOT EXISTS Employee(
    Id SERIAL PRIMARY KEY,
    Name varchar(20) NOT NULL unique,
    Password bytea NOT NULL,
    Email varchar(30) NOT NULL unique,
    Status boolean NOT NULL,
    DOB date NOT NULL
);`
  	

  func Connect() (*sqlx.DB){
	  connection, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=admin sslmode=disable")
	  if err != nil {
        log.Fatalln(err)
    }
	  return connection
  }

  func CreateNew(connection *sqlx.DB) {
    _, err := connection.Exec(schema)
    if err != nil {
      log.Fatalln(err)
    }
    fmt.Print("close")
  }
  
 /*res,_ := db.Query("SELECT EXISTS(SELECT FROM information_schema.tables WHERE table_schema = 'schema' AND table_name = 'employee');")
    
    if res != nil{
     
    }*/
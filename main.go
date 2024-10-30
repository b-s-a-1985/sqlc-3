package main

import (
	"context"
	"database/sql"
	"os"
	"sqlc3/person"

	"github.com/joho/godotenv" // .env library
	_ "github.com/lib/pq"
)

func main() {

	// load connection string
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file" + err.Error())
	}
	connStr := os.Getenv("CONNSTR")

	// connect database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ctx := context.Background()
	query := person.New(db)

	// create schema
	err = query.CreateSchema(ctx)
	if err != nil {
		panic("CreateSchema failed - " + err.Error())
	}

	// select active schema
	err = query.SetSchema(ctx)
	if err != nil {
		panic("SetSchema failed - " + err.Error())
	}

	// create table
	err = query.CreateTable(ctx)
	if err != nil {
		panic("CreateTable failed - " + err.Error())
	}

	// insert data
	// err = query.InsertDepartment(ctx, "Human Resources")
	// if err != nil {
	// 	panic("InsertDepartment (Human Resources) failed - " + err.Error())
	// }
	err = query.InsertDepartment(ctx, "Finance")
	if err != nil {
		panic("InsertDepartment (Finance) failed - " + err.Error())
	}
	err = query.InsertDepartment(ctx, "IT")
	if err != nil {
		panic("InsertDepartment (IT) failed - " + err.Error())
	}
} // main

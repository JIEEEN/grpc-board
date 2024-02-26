package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dbType := os.Getenv("DB_TYPE")
	dbDetail := os.Getenv("DB_DETAIL")

	var err error

	DB, err = sql.Open(dbType, dbDetail)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	DB.SetConnMaxLifetime(time.Minute * 5)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	log.Printf("Connect success to db: %v", dbType)

	dbName := createDB()
	log.Printf("Success to create database: %s", dbName)

	createTable()
	log.Print("Success to create tables")
}

func createDB() string {
	dbName := os.Getenv("DB_DATABASE")
	_, err := DB.Exec("create database if not exists " + dbName)
	if err != nil {
		log.Fatalf("Cannot create database %s: %v", dbName, err)
	}

	_, err = DB.Exec("use " + dbName)
	if err != nil {
		log.Fatalf("Cannot use database %s: %v", dbName, err)
	}

	return dbName
}

func createTable() {
	_, err := DB.Exec(`
		create table if not exists users
		(
			id int not null auto_increment primary key,
			email varchar(255) unique not null,
			password varchar(255) not null,
			nickname varchar(50),
			create_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp on update current_timestamp,
			deleted_at timestamp
		)
	`)
	if err != nil {
		log.Fatalf("Cannot create table users: %v", err)
	}

	_, err = DB.Exec(`
		create table if not exists boards
		(
			id int not null auto_increment primary key,
			user_id int,
			nickname varchar(50) unique not null,
			title varchar(255) not null,
			contents text,
			create_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp on update current_timestamp,
			deleted_at timestamp,
			foreign key (user_id) references users(id)
		)
	`)
	if err != nil {
		log.Fatalf("Cannot create table boards: %v", err)
	}

	_, err = DB.Exec(`
		create table if not exists replies
		(
			id int not null auto_increment primary key,
			user_id int,
			nickname varchar(255),
			board_id int,
			contents text,
			create_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp on update current_timestamp,
			deleted_at timestamp,
			foreign key (user_id) references users(id),
			foreign key (board_id) references boards(id)
		)
	`)
	if err != nil {
		log.Fatalf("Cannot create table replies: %v", err)
	}
}

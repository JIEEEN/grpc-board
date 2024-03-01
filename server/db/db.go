package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dbType := os.Getenv("DB_TYPE")
	dbDetail := os.Getenv("DB_DETAIL")

	db, err := sql.Open(dbType, dbDetail)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	log.Printf("Connect success to db: %v", dbType)

	dbName := createDB(db)
	log.Printf("Success to create database: %s", dbName)

	createTable(db)
	log.Print("Success to create tables")

	return db, nil
}

func createDB(db *sql.DB) string {
	dbName := os.Getenv("DB_DATABASE")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ctx, "create database if not exists "+dbName)
	if err != nil {
		log.Fatalf("Cannot create database %s: %v", dbName, err)
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error occurs when fetching rows: %v", err)
	}
	log.Printf("Rows Affected: %d\n", no)

	_, err = db.ExecContext(ctx, "use "+dbName)
	if err != nil {
		log.Fatalf("Cannot use database %s: %v", dbName, err)
	}

	return dbName
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
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

	_, err = db.Exec(`
		create table if not exists board_groups
		(
			id int not null auto_increment primary key,
			name varchar(50) not null,
			user_id int,
			foreign key (user_id) references users(id)
		)
	`)
	if err != nil {
		log.Fatalf("Cannot create table groups: %v", err)
	}

	_, err = db.Exec(`
		create table if not exists boards
		(
			id int not null auto_increment primary key,
			user_id int,
			title varchar(255) not null,
			contents text,
			create_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp on update current_timestamp,
			deleted_at timestamp,
			group_id int,
			foreign key (user_id) references users(id),
			foreign key (group_id) references board_groups(id)
		)
	`)
	if err != nil {
		log.Fatalf("Cannot create table boards: %v", err)
	}

	_, err = db.Exec(`
		create table if not exists group_users
		(
			id int not null auto_increment primary key,
			group_id int,
			user_id int,
			foreign key (group_id) references board_groups(id),
			foreign key (user_id) references users(id)
		)
	`)
	if err != nil {
		log.Fatalf("Cannot create table grou_users: %v", err)
	}

	_, err = db.Exec(`
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

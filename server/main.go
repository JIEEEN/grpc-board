package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/grpc"

	"github.com/JIEEEN/grpc-board/server/db"
	"github.com/JIEEEN/grpc-board/server/env"
	"github.com/JIEEEN/grpc-board/server/user"
)

var _db *sql.DB

func init() {
	env.InitEnv()
	_db, _ = db.InitDB()
}

func main() {
	defer _db.Close()
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot connect to Server: %v", err)
	}
	s := grpc.NewServer()

	userServer := user.NewServer(_db)

	pb.RegisterUserServiceServer(s, userServer)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

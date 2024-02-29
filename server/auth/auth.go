package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
)

type server struct {
	pb.UnimplementedAuthServiceServer
	randomString string
	db           *sql.DB
}

func NewServer(db *sql.DB) *server {
	return &server{
		db: db,
	}
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	email := in.GetEmail()
	password := in.GetPassword()

	var r_email, r_password string
	err := s.db.QueryRow("select email from users where email = ?", email).Scan(&r_email)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.LoginResponse{
			Result:      false,
			AccessToken: "",
		}, err
	}
	if r_email == "" {
		err = errors.New("no Email data on db")
		log.Printf("No Email data on db: %v", err)
		return &pb.LoginResponse{
			Result:      false,
			AccessToken: "",
		}, err
	}

	err = s.db.QueryRow("select password from users where email = ?", email).Scan(&r_password)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.LoginResponse{
			Result:      false,
			AccessToken: "",
		}, err
	}

	if ComparePassword(r_password, password) {
		return &pb.LoginResponse{
			Result:      true,
			AccessToken: "",
		}, nil
	} else {
		log.Printf("Password mismatch: %v", err)
		err := errors.New("password mismatch")
		return &pb.LoginResponse{
			Result:      false,
			AccessToken: "",
		}, err
	}
}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	email := in.GetEmail()       // unique
	nickname := in.GetNickname() // unique

	var err error
	var r_email, r_nickname string
	_ = s.db.QueryRow("select email from users where email = ?", email).Scan(&r_email)
	if r_email != "" {
		log.Printf("Email already exists")
		err = errors.New("email already exists")
		return &pb.RegisterResponse{
			Result: false,
		}, err
	}
	_ = s.db.QueryRow("select nickname from users where nickname = ?", nickname).Scan(&r_nickname)
	if r_nickname != "" {
		log.Printf("Nickname already exists")
		err = errors.New("nickname already exists")
		return &pb.RegisterResponse{
			Result: false,
		}, err
	}

	s.randomString, err = SendMail(email)
	if err != nil {
		log.Printf("Cannot send mail: %v", err)
		return &pb.RegisterResponse{
			Result: false,
		}, err
	}

	return &pb.RegisterResponse{
		Result: true,
	}, nil
}

func (s *server) MailVerification(ctx context.Context, in *pb.MailVerificationRequest) (*pb.MailVerificationResponse, error) {
	vcode := in.GetVcode()
	email := in.GetEmail()
	password := in.GetPassword()
	nickname := in.GetNickname()
	if nickname == "" {
		nickname = GenerateUsername()
	}

	if s.randomString != vcode {
		err := errors.New("codes don't match")
		log.Printf("Codes don't match")
		return &pb.MailVerificationResponse{
			Result: false,
		}, err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Printf("Error occur while hash password: %v", err)
		return &pb.MailVerificationResponse{
			Result: false,
		}, err
	}
	_, err = s.db.Exec("insert into users (email, password, nickname) values (?, ?, ?)", email, hashedPassword, nickname)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.MailVerificationResponse{
			Result: false,
		}, err
	}

	return &pb.MailVerificationResponse{
		Result: true,
	}, nil
}

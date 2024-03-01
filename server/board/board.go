package board

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedBoardServiceServer
	db *sql.DB
}

func NewServer(db *sql.DB) *server {
	return &server{
		db: db,
	}
}

func (s *server) BoardList(ctx context.Context, in *emptypb.Empty) (*pb.BoardListResponse, error) {
	boards := []*pb.Board{}

	rows, err := s.db.Query("select * from boards")
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.BoardListResponse{}, err
	}
	defer rows.Close()

	var r *pb.Board
	for rows.Next() {
		err = rows.Scan(r)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.BoardListResponse{}, err
		}

		boards = append(boards, r)
	}

	return &pb.BoardListResponse{
		Boardlist: boards,
	}, nil
}

func (s *server) CreateBoard(ctx context.Context, in *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	board := &pb.Board{
		Id:       in.GetBoard().GetId(),
		Title:    in.GetBoard().GetTitle(),
		Contents: in.GetBoard().GetContents(),
	}

	_, err := s.db.Exec("insert into boards (id, title, contents) values (?, ?, ?, ?)",
		board.Id, board.Title, board.Contents)

	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.CreateBoardResponse{
			Result: false,
		}, err
	}

	return &pb.CreateBoardResponse{
		Result: true,
	}, nil
}

func (s *server) ReadBoard(ctx context.Context, in *pb.ReadBoardRequest) (*pb.ReadBoardResponse, error) {
	id := in.GetId()

	var r *pb.Board
	err := s.db.QueryRow("select *from boards where id = ?", id).Scan(r)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.ReadBoardResponse{}, nil
	}

	return &pb.ReadBoardResponse{Board: r}, nil
}

func (s *server) UpdateBoard(ctx context.Context, in *pb.UpdateBoardRequest) (*pb.UpdateBoardResponse, error) {
	board := in.GetBoard()

	id := board.GetId()

	var r_title, r_contents string
	err := s.db.QueryRow("select title, contents from boards where id = ?", id).Scan(&r_title, &r_contents)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UpdateBoardResponse{
			Result: false,
		}, err
	}

	_, err = s.db.Exec("update ignore boards set title = ?, contents = ? where id = ?", board.GetTitle(), board.GetContents(), id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UpdateBoardResponse{
			Result: false,
		}, err
	}

	return &pb.UpdateBoardResponse{
		Result: true,
	}, nil
}

func (s *server) DeleteBoard(ctx context.Context, in *pb.DeleteBoardRequest) (*pb.DeleteBoardResponse, error) {
	id := in.GetId()

	_, err := s.db.Exec("delete from boards where id = ?", id)
	if err != nil {
		log.Printf("Failed to delete item from boards: %v", err)
		return &pb.DeleteBoardResponse{
			Result: false,
		}, err
	}

	return &pb.DeleteBoardResponse{
		Result: true,
	}, nil
}

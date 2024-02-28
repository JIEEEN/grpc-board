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

	rows, err := s.db.Query("select id, nickname, title, contents from boards")
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		return &pb.BoardListResponse{}, err
	}
	defer rows.Close()

	var r_id, r_nickname, r_title, r_contents string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_nickname, &r_title, &r_contents)
		if err != nil {
			log.Fatalf("Failed to get row data: %v", err)
			return &pb.BoardListResponse{}, err
		}

		board := &pb.Board{
			Id:       r_id,
			Nickname: r_nickname,
			Title:    r_title,
			Contents: r_contents,
		}

		boards = append(boards, board)
	}

	return &pb.BoardListResponse{
		Boardlist: boards,
	}, nil
}

func (s *server) CreateBoard(ctx context.Context, in *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	board := &pb.Board{
		Id:       in.GetBoard().GetId(),
		Nickname: in.GetBoard().GetNickname(),
		Title:    in.GetBoard().GetTitle(),
		Contents: in.GetBoard().GetContents(),
	}

	_, err := s.db.Exec("insert into boards (id, nickname, title, contents) values (?, ?, ?, ?)",
		board.Id, board.Nickname, board.Title, board.Contents)

	var success bool
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		success = false
		return &pb.CreateBoardResponse{
			Result: success,
		}, err
	}

	success = true
	return &pb.CreateBoardResponse{
		Result: success,
	}, nil
}

func (s *server) ReadBoard(ctx context.Context, in *pb.ReadBoardRequest) (*pb.ReadBoardResponse, error) {
	id := in.GetId()

	var r_id, r_nickname, r_title, r_contents string
	err := s.db.QueryRow("select id, nickname, title, contents from boards where id = ?", id).Scan(&r_id, &r_nickname, &r_title, &r_contents)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		return &pb.ReadBoardResponse{}, nil
	}

	board := pb.Board{
		Id:       r_id,
		Nickname: r_nickname,
		Title:    r_title,
		Contents: r_contents,
	}

	return &pb.ReadBoardResponse{Board: &board}, nil
}

func (s *server) UpdateBoard(ctx context.Context, in *pb.UpdateBoardRequest) (*pb.UpdateBoardResponse, error) {
	board := in.GetBoard()

	id := board.GetId()

	var success bool
	var r_title, r_contents string
	err := s.db.QueryRow("select title, contents from boards where id = ?", id).Scan(&r_title, &r_contents)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		success = false
		return &pb.UpdateBoardResponse{
			Result: success,
		}, err
	}

	_, err = s.db.Exec("update ignore boards set title = ?, contents = ? where id = ?", board.GetTitle(), board.GetContents(), id)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		success = false
		return &pb.UpdateBoardResponse{
			Result: success,
		}, err
	}

	success = true
	return &pb.UpdateBoardResponse{
		Result: success,
	}, nil
}

func (s *server) DeleteBoard(ctx context.Context, in *pb.DeleteBoardRequest) (*pb.DeleteBoardResponse, error) {
	id := in.GetId()

	var success bool
	_, err := s.db.Exec("delete from boards where id = ?", id)
	if err != nil {
		log.Fatalf("Failed to delete item from boards: %v", err)
		success = false
		return &pb.DeleteBoardResponse{
			Result: success,
		}, err
	}

	success = true
	return &pb.DeleteBoardResponse{
		Result: success,
	}, nil
}

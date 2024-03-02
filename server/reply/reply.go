package reply

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
)

type server struct {
	pb.UnimplementedReplyServiceServer
	db *sql.DB
}

func NewServer(db *sql.DB) *server {
	return &server{
		db: db,
	}
}

func (s *server) GetReplyUser(ctx context.Context, in *pb.GetReplyUserRequest) (*pb.GetReplyUserResponse, error) {
	userId := in.GetUserId()
	replies := []*pb.Reply{}

	rows, err := s.db.Query("select id, user_id, nickname, board_id, contents from replies where user_id = ?", userId)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.GetReplyUserResponse{}, err
	}
	defer rows.Close()

	var r_id, r_userId, r_nickname, r_boardId, r_contents string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_userId, &r_nickname, &r_boardId, &r_contents)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.GetReplyUserResponse{}, err
		}

		reply := &pb.Reply{
			Id:       r_id,
			UserId:   r_userId,
			Nickname: r_nickname,
			BoardId:  r_boardId,
			Contents: r_contents,
		}

		replies = append(replies, reply)
	}

	return &pb.GetReplyUserResponse{
		Replies: replies,
	}, nil
}

func (s *server) GetReplyNickname(ctx context.Context, in *pb.GetReplyNicknameRequest) (*pb.GetReplyNicknameResponse, error) {
	nickname := in.GetNickname()
	replies := []*pb.Reply{}

	rows, err := s.db.Query("select id, user_id, nickname, board_id, contents from replies where nickname = ?", nickname)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.GetReplyNicknameResponse{}, err
	}
	defer rows.Close()

	var r_id, r_userId, r_nickname, r_boardId, r_contents string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_userId, &r_nickname, &r_boardId, &r_contents)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.GetReplyNicknameResponse{}, err
		}

		reply := &pb.Reply{
			Id:       r_id,
			UserId:   r_userId,
			Nickname: r_nickname,
			BoardId:  r_boardId,
			Contents: r_contents,
		}

		replies = append(replies, reply)
	}

	return &pb.GetReplyNicknameResponse{
		Replies: replies,
	}, nil
}

func (s *server) CreateReply(ctx context.Context, in *pb.CreateReplyRequest) (*pb.CreateReplyResponse, error) {
	newReply := in.GetReplyRequest()

	_, err := s.db.Exec("insert into replies (user_id, nickname, board_id, contents) values (?, ?, ?, ?)",
		newReply.GetUserId(), newReply.GetNickname(), newReply.GetBoardId(), newReply.GetContents())
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.CreateReplyResponse{
			Result: false,
		}, err
	}

	return &pb.CreateReplyResponse{
		Result: true,
	}, nil
}

func (s *server) GetAllReplyBoard(ctx context.Context, in *pb.AllReplyBoardRequest) (*pb.AllReplyBoardResponse, error) {
	boardId := in.GetBoardId()
	replies := []*pb.Reply{}

	rows, err := s.db.Query("select id, user_id, nickname, board_id, contents from replies where board_id = ?", boardId)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.AllReplyBoardResponse{}, err
	}
	defer rows.Close()

	var r_id, r_userId, r_nickname, r_boardId, r_contents string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_userId, &r_nickname, &r_boardId, &r_contents)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.AllReplyBoardResponse{}, err
		}

		reply := &pb.Reply{
			Id:       r_id,
			UserId:   r_userId,
			Nickname: r_nickname,
			BoardId:  r_boardId,
			Contents: r_contents,
		}

		replies = append(replies, reply)
	}

	return &pb.AllReplyBoardResponse{
		Replies: replies,
	}, nil
}

func (s *server) UpdateReply(ctx context.Context, in *pb.UpdateReplyRequest) (*pb.UpdateReplyResponse, error) {
	reply := in.GetReply()

	_, err := s.db.Exec("update ignore replies set nickname = ?, contents = ? where id = ?", reply.GetNickname(), reply.GetContents(), reply.GetId())
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UpdateReplyResponse{
			Result: false,
		}, err
	}

	return &pb.UpdateReplyResponse{
		Result: true,
	}, nil
}

func (s *server) DeleteReply(ctx context.Context, in *pb.DeleteReplyRequest) (*pb.DeleteReplyResponse, error) {
	id := in.GetId()

	_, err := s.db.Exec("delete from replies where id = ?", id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.DeleteReplyResponse{
			Result: false,
		}, err
	}

	return &pb.DeleteReplyResponse{
		Result: true,
	}, nil
}

package group

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/JIEEEN/grpc-board/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type boardServer struct {
	pb.UnimplementedBoardGroupServiceServer
	db *sql.DB
}

func NewBoardServer(db *sql.DB) *boardServer {
	return &boardServer{
		db: db,
	}
}

func (s *boardServer) GetAllBoardGroups(ctx context.Context, in *emptypb.Empty) (*pb.BoardGroupsInfo, error) {
	boardGroups := []*pb.BoardGroup{}

	rows, err := s.db.Query("select id, name from board_groups")
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.BoardGroupsInfo{}, err
	}
	defer rows.Close()

	var r_id, r_name string
	for rows.Next() {
		err = rows.Scan(&r_id, &r_name)
		if err != nil {
			log.Printf("Failed to get row data: %v", err)
			return &pb.BoardGroupsInfo{}, err
		}

		boardGroup := &pb.BoardGroup{
			Id:   r_id,
			Name: r_name,
		}

		boardGroups = append(boardGroups, boardGroup)
	}

	return &pb.BoardGroupsInfo{
		BoardGroups: boardGroups,
	}, nil
}

func (s *boardServer) GetBoardId(ctx context.Context, in *pb.GetBoardGroupRequest) (*pb.GetBoardGroupResponse, error) {
	id := in.GetId()

	var r_id, r_name string
	err := s.db.QueryRow("select id, name from board_groups where id = ?", id).Scan(&r_id, &r_name)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.GetBoardGroupResponse{}, err
	}

	boardGroup := &pb.BoardGroup{
		Id:   r_id,
		Name: r_name,
	}

	return &pb.GetBoardGroupResponse{
		BoardGroup: boardGroup,
	}, nil
}

func (s *boardServer) CreateBoardGroup(ctx context.Context, in *pb.CreateBoardGroupRequest) (*pb.CreateBoardGroupResponse, error) {
	boardGroup := in.GetBoardGroupRequest()
	name := boardGroup.GetName()

	_, err := s.db.Exec("insert into board_groups (name) values (?)", name)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.CreateBoardGroupResponse{}, err
	}

	return &pb.CreateBoardGroupResponse{
		Result: true,
	}, nil
}

func (s *boardServer) UpdateBoardGroup(ctx context.Context, in *pb.UpdateBoardGroupRequest) (*pb.UpdateBoardGroupResponse, error) {
	boardGroup := in.GetBoardGroup()
	id := boardGroup.GetId()
	name := boardGroup.GetName()

	_, err := s.db.Exec("update board_groups set name = ? where id = ?", name, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.UpdateBoardGroupResponse{}, err
	}

	return &pb.UpdateBoardGroupResponse{
		Result: true,
	}, nil
}

func (s *boardServer) DeleteBoardGroup(ctx context.Context, in *pb.DeleteBoardGroupRequest) (*pb.DeleteBoardGroupResponse, error) {
	id := in.GetId()

	_, err := s.db.Exec("delete from board_groups where id = ?", id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return &pb.DeleteBoardGroupResponse{}, err
	}

	return &pb.DeleteBoardGroupResponse{
		Result: true,
	}, nil
}

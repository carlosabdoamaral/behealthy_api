package persistence

import (
	"context"
	"database/sql"

	"github.com/carlosabdoamaral/behealthy_api/common"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

func UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.Status, error) {
	var (
		db    *sql.DB = common.Database
		query string  = AccountQueriesBuilder.UpdatePassword(req.GetEmail(), req.GetNewPassword())
	)

	_, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return &pb.Status{
		Message: "password updated sucessfully",
	}, nil
}

func SoftDelete(ctx context.Context, req *pb.UpdateAccountStatus) (*pb.Status, error) {
	var (
		db    *sql.DB = common.Database
		query string  = AccountQueriesBuilder.SoftDelete(req.GetEmail())
	)

	_, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return &pb.Status{
		Message: "account soft deleted sucessfully",
	}, nil
}

func Restore(ctx context.Context, req *pb.UpdateAccountStatus) (*pb.Status, error) {
	var (
		db    *sql.DB = common.Database
		query string  = AccountQueriesBuilder.Restore(req.GetEmail())
	)

	_, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return &pb.Status{
		Message: "account restored sucessfully",
	}, nil
}

func Block(ctx context.Context, req *pb.UpdateAccountStatus) (*pb.Status, error) {
	var (
		db    *sql.DB = common.Database
		query string  = AccountQueriesBuilder.Block(req.GetEmail())
	)

	_, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return &pb.Status{
		Message: "account blocked sucessfully",
	}, nil
}

func Unblock(ctx context.Context, req *pb.UpdateAccountStatus) (*pb.Status, error) {
	var (
		db    *sql.DB = common.Database
		query string  = AccountQueriesBuilder.Unblock(req.GetEmail())
	)

	_, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return &pb.Status{
		Message: "account unblocked sucessfully",
	}, nil
}

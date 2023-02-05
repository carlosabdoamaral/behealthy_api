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

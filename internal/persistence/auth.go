package persistence

import (
	"context"
	"database/sql"
	"errors"

	"github.com/carlosabdoamaral/behealthy_api/common"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

func SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.AccountDetails, error) {
	var (
		db    *sql.DB            = common.Database
		query string             = AuthQueriesBuilder.SignIn(req)
		res   *pb.AccountDetails = &pb.AccountDetails{}
	)

	res, err := ScanSingleAccountDetails(db.Query(query))
	if err != nil {
		return nil, err
	}

	if res.GetAccountId() == 0 {
		return nil, errors.New("account not found")
	}

	return res, nil
}

func SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.AccountDetails, error) {
	// Implement me
	return nil, nil
}

func ValidateTwoFactorCode(ctx context.Context, req *pb.TwoFactorCode) (*pb.Status, error) {
	var (
		db    *sql.DB = common.Database
		query string  = AuthQueriesBuilder.ValidateTwoFactorCode(req)
	)

	resp := 0
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&resp)
		if err != nil {
			return nil, err
		}
	}

	if resp == 0 {
		err = errors.New("two factor code is wrong")
		return nil, err
	}

	return &pb.Status{
		Message: "two factor code is valid",
	}, nil
}

func RefreshTwoFactorCode(ctx context.Context, req *pb.RefreshTwoFactorCodeRequest) (*pb.Status, error) {
	var (
		db    = common.Database
		query = AuthQueriesBuilder.RefreshTwoFactorCode(req)
	)

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return &pb.Status{
		Message: "updated successfully",
	}, nil
}

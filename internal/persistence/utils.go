package persistence

import (
	"database/sql"

	"github.com/carlosabdoamaral/behealthy_api/internal/responses"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

func ScanSingleAccountDetails(rows *sql.Rows, _ error) (*pb.AccountDetails, error) {
	resp := &responses.AccountDetails{}
	for rows.Next() {
		err := rows.Scan(
			&resp.AccountId,
			&resp.FullName,
			&resp.UserName,
			&resp.Email,
			&resp.Password,
			&resp.BirthDate,
			&resp.Role,
			&resp.CreatedAt,
			&resp.UpdatedAt,
			&resp.SoftDeleted,
			&resp.IsBlocked,
			&resp.TwoFactorCode,
		)

		if err != nil {
			return nil, err
		}
	}

	return responses.NewAccountDetailsFromJSONToProto(resp), nil
}

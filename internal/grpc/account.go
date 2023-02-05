package grpc

import (
	"context"

	"github.com/carlosabdoamaral/behealthy_api/internal/persistence"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

func (s *AccountServer) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.Status, error) {
	return persistence.UpdatePassword(ctx, req)
}

func (*AccountServer) SoftDelete(ctx context.Context, req *pb.UpdateAccountStatus) (*pb.Status, error) {
	return persistence.SoftDelete(ctx, req)
}

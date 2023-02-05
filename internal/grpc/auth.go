package grpc

import (
	"context"

	"github.com/carlosabdoamaral/behealthy_api/internal/persistence"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
)

func (s *AuthServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.AccountDetails, error) {
	return persistence.SignIn(ctx, req)
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.AccountDetails, error) {
	return persistence.SignUp(ctx, req)
}

func (s *AuthServer) ValidateTwoFactorCode(ctx context.Context, req *pb.TwoFactorCode) (*pb.Status, error) {
	return persistence.ValidateTwoFactorCode(ctx, req)
}

func (s *AuthServer) RefreshTwoFactorCode(ctx context.Context, req *pb.RefreshTwoFactorCodeRequest) (*pb.Status, error) {
	return persistence.RefreshTwoFactorCode(ctx, req)
}

package service

import (
	pb "auth/genproto/user"
	"auth/storage/postgres"
	"context"
	"log/slog"
)

type UserManagementService interface {
	RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	// ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error)
	// RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
	Logout(ctx context.Context, token *pb.LogoutRequest) (*pb.Message, error)
	GetUserProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error)
	// SaveToken(ctx context.Context, req *pb.Token) (*pb.Void, error)
}

type UserService struct {
	pb.UnimplementedUsersServer
	authRepo postgres.IUserStorages
	Logger   *slog.Logger
}

func NewUserService(authRepo postgres.IUserStorages, logger *slog.Logger) *UserService {
	return &UserService{authRepo: authRepo, Logger: logger}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp, err := s.authRepo.RegisterUser(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := s.authRepo.LoginUser(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
// func (s *UserService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
// 	resp, err := s.authRepo.ValidateToken(ctx, req)
// 	if err != nil {
// 		s.Logger.Error(err.Error())
// 		return nil, err
// 	}
// 	return resp, nil
// }
// func (s *UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
// 	resp, err := s.authRepo.RefreshToken(ctx, req)
// 	if err != nil {
// 		s.Logger.Error(err.Error())
// 		return nil, err
// 	}
// 	return resp, nil
// }

func (s *UserService) Logout(ctx context.Context, token *pb.LogoutRequest) (*pb.Message, error) {
	resp, err := s.authRepo.Logout(ctx, token)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
func (s *UserService) GetUserProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	resp, err := s.authRepo.GetUserProfile(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error) {
	resp, err := s.authRepo.UpdateUserProfile(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
// func (s *UserService) SaveToken(ctx context.Context, req *pb.Token) (*pb.Void, error) {
// 	resp, err := s.authRepo.SaveToken(ctx, req)
// 	if err != nil {
// 		s.Logger.Error(err.Error())
// 		return nil, err
// 	}
// 	return resp, nil
// }

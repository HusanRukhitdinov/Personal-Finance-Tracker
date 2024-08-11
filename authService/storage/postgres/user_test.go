package postgres

import (
	pb "auth/genproto/user"
	"context"
	"fmt"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	db, err := ConnectionDb()

	if err != nil {
		panic(err)
	}

	a := NewUserRepository(db)
	
	req := pb.RegisterRequest{
		Username: "aasdsdb32ba",
		Email:    "aaccsf32sz3a",
		Password: "aaff233efa",
		Role:     "admirsdn",
		Fullname: "aae4354eea",
	}
	ctx := context.Background()

	resp, err := a.RegisterUser(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func TestLoginUser(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	a := NewUserRepository(db)

	req := pb.LoginRequest{
		Email:    "stsdring",
		Password: "$2a$10$scH/zSZHB1BaWixK4I48xO/OUsryDGN1IWQkw0rqzQ9VGeVOAza0G",
	}
	ctx := context.Background()

	resp, err := a.LoginUser(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestLogout(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	a := NewUserRepository(db)

	req := pb.LogoutRequest{
		Token: "Votlp9rgIy1DOPsApVlQQvahiJGWIBaJtWxvpTEVZCgRz499MAjdIgjB0Y44yUgd",
	}
	ctx := context.Background()

	res, err := a.Logout(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetUserProfile(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	a := NewUserRepository(db)

	ctx := context.Background()
	req := &pb.ProfileRequest{
		Id: "71b3a22c-be2c-49e5-84c5-6ed7e94f6a70",
	}

	resp, err := a.GetUserProfile(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

// func TestValidateToken(t *testing.T) {
// 	db, err := ConnectionDb()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	a := NewUserRepository(db)

// 	ctx := context.Background()
// 	req := &pb.ValidateTokenRequest{
// 		Token: "",
// 	}

// 	resp, err := a.ValidateToken(ctx, req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp)
// }

// func TestRefreshToken(t *testing.T) {
// 	db, err := ConnectionDb()
// 	if err != nil {
// 		panic(err)
// 	}
// 	a := NewUserRepository(db)

// 	req := pb.RefreshTokenRequest{
// 		Token: "",
// 	}
// 	ctx := context.Background()

// 	res, err := a.RefreshToken(ctx, &req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(res)
// }

func TestUpdateUserProfile(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	a := NewUserRepository(db)

	req := pb.ProfileUpdateRequest{
		Username: "aaa",
		Email:    "aaa",
	}
	ctx := context.Background()

	res, err := a.UpdateUserProfile(ctx, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

package postgres

import (
	pb "auth/genproto/user"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IUserStorages interface {
	RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, token *pb.LogoutRequest) (*pb.Message, error)
	GetUserProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error)
}
type UserRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewUserRepository(db *sql.DB) IUserStorages {
	return &UserRepository{Db: db}
}

func (u *UserRepository) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	var userID string
	resp := pb.RegisterResponse{}
	userQuery :=
		`INSERT INTO users (username,email,password,role,fullname)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id,username,email,password,role,fullname,created_at`

	err = u.Db.QueryRow(userQuery, req.Username, req.Email, passwordHash, req.Role, req.Fullname).Scan(
		&userID, &resp.Username, &resp.Email, &resp.Password, &resp.Role, &resp.Fullname, &resp.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (u *UserRepository) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	query := `SELECT id,username, role FROM users WHERE email = $1`

	var id, username, role string

	err = u.Db.QueryRowContext(ctx, query, req.Email).Scan(&id, &username, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return nil, fmt.Errorf("email yoki parol noto'g'ri")
		}
		return nil, fmt.Errorf("so'rov bajarilmasligi: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.New("password is incorrect")
		}
		return nil, err
	}

	return &pb.LoginResponse{
		UserId:   id,
		Username: username,
		Role:     role,
	}, nil

}

func (a *UserRepository) Logout(ctx context.Context, token *pb.LogoutRequest) (*pb.Message, error) {
	_, err := a.Db.Exec(`
	update 
		token 
	set 
		deleted_at=$1 
	where 
		token=$2`, time.Now(), token.Token)
	if err != nil {
		return nil, err
	}
	return &pb.Message{Message: "Logout successfully"}, nil
}

func (u *UserRepository) GetUserProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	resp := &pb.ProfileResponse{}

	q := `
	select 
		username, 
		email, 
		fullname, 
		created_at 
	from 
		users
	where 
		id = $1
	`

	rows, err := u.Db.Query(q, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&resp.Username,
			&resp.Email,
			&resp.Fullname,
			&resp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (u *UserRepository) UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error) {
	resp := pb.ProfileUpdateResponse{}

	q := `
	update 
		users 
	set 
		username=$1
	where 
		email=$2
	returning
		id, 
		username, 
		email, 
		fullname, 
		updated_at 
	`

	err := u.Db.QueryRow(q, req.Username, req.Email).Scan(
		&resp.Id,
		&resp.Username,
		&resp.Email,
		&resp.Fullname,
		&resp.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}



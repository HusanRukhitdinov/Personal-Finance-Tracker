package service

import (
	pb "budget/genproto/budgeting_service"
	"budget/storage"
	"context"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type TransactionService struct {
	storage storage.IStorage
	log     slog.Logger
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(storage storage.IStorage, log slog.Logger) *TransactionService {
	return &TransactionService{
		storage: storage,
		log:     log,
	}
}

func (service *TransactionService) CreateTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.Transaction, error) {
	return service.storage.Transactions().CreateTransaction(ctx, in)

}
func (service *TransactionService) UpdateTransaction(ctx context.Context, in *pb.Transaction) (*pb.Transaction, error) {
	return service.storage.Transactions().UpdateTransaction(ctx, in)
}
func (service *TransactionService) GetTransaction(ctx context.Context, in *pb.PrimaryKey) (*pb.Transaction, error) {
	return service.storage.Transactions().GetTransaction(ctx, in)
}
func (service *TransactionService) GetListTransactions(ctx context.Context, in *pb.GetListRequest) (*pb.Transactions, error) {
	return service.storage.Transactions().GetAllTransaction(ctx, in)
}
func (service *TransactionService) DeleteTransaction(ctx context.Context, in *pb.PrimaryKey) (*emptypb.Empty, error) {
	return service.storage.Transactions().DeleteTransaction(ctx, in)
}
func (service *TransactionService) GetUserSpending(ctx context.Context, in *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	return service.storage.Transactions().GetUserTotalSpend(ctx, in)
}
func (service *TransactionService) GetUserIncome(ctx context.Context, in *pb.GetUserMoneyRequest) (*pb.GetUserMoneyResponse, error) {
	return service.storage.Transactions().GetUserTotalIncome(ctx, in)
}

package service

import (
	pb "budget/genproto/budgeting_service"
	"budget/storage"
	"context"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

type BudgetService struct {
	storage storage.IStorage
	log     slog.Logger
	pb.UnimplementedBudgetServiceServer
}

func NewBudgetService(storage storage.IStorage, log slog.Logger) *BudgetService {
	return &BudgetService{
		storage: storage,
		log:     log,
	}
}

func (service *BudgetService) CreateBudget(ctx context.Context, in *pb.BudgetRequest) (*pb.Budget, error) {
	return service.storage.Budgets().CreateBudget(ctx, in)

}
func (service *BudgetService) UpdateBudget(ctx context.Context, in *pb.Budget) (*pb.Budget, error) {
	return service.storage.Budgets().UpdateBudget(ctx, in)
}
func (service *BudgetService) GetBudget(ctx context.Context, in *pb.PrimaryKey) (*pb.Budget, error) {
	return service.storage.Budgets().GetBudget(ctx, in)
}
func (service *BudgetService) GetListBudgets(ctx context.Context, in *pb.GetListRequest) (*pb.Budgets, error) {
	return service.storage.Budgets().GetAllBudget(ctx, in)
}
func (service *BudgetService) DeleteBudget(ctx context.Context, in *pb.PrimaryKey) (*emptypb.Empty, error) {
	return service.storage.Budgets().DeleteBudget(ctx, in)
}

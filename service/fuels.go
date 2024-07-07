package service

import (
	"context"

	pb "github.com/Military/military-service/genprotos/militaries"
	"github.com/Military/military-service/storage"
)

type FuelService struct {
	Fuel storage.StorageI
	pb.UnimplementedFuelServiceServer
}

func NewFuelService(fuel storage.StorageI) *FuelService {
	return &FuelService{
		Fuel: fuel,
	}
}

func (f *FuelService) Create(ctx context.Context, req *pb.FuelReq) (*pb.Void, error) {
	return f.Fuel.Fuel().Create(ctx, req)
}

func (f *FuelService) Update(ctx context.Context, req *pb.Fuel) (*pb.Void, error) {
	return f.Fuel.Fuel().Update(ctx, req)
}

func (f *FuelService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return f.Fuel.Fuel().Delete(ctx, req)
}

func (f *FuelService) Get(ctx context.Context, req *pb.ById) (*pb.Fuel, error) {
	return f.Fuel.Fuel().Get(ctx, req)
}

func (f *FuelService) GetAll(ctx context.Context, req *pb.FuelReq) (*pb.Fuel, error) {
	return f.Fuel.Fuel().GetAll(ctx, req)
}
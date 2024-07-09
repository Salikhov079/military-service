package service


import (
	"context"

	pb "github.com/Military/military-service/genprotos/militaries"
	"github.com/Military/military-service/storage"
)

type TechniqueService struct {
	Technique storage.StorageI
	pb.UnimplementedTechniqueServiceServer
}

func NewTechniqueService(technique storage.StorageI) *TechniqueService {
	return &TechniqueService{
		Technique: technique,
	}
}

func (t *TechniqueService) Create(ctx context.Context, req *pb.TechniqueReq) (*pb.Void, error) {
	return t.Technique.Technique().Create(ctx, req)
}

func (t *TechniqueService) Update(ctx context.Context, req *pb.Technique) (*pb.Void, error) {
	return t.Technique.Technique().Update(ctx, req)
}

func (t *TechniqueService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return t.Technique.Technique().Delete(ctx, req)
}

func (t *TechniqueService) Get(ctx context.Context, req *pb.ById) (*pb.Technique, error) {
	return t.Technique.Technique().Get(ctx, req)
}

func (t *TechniqueService) GetAll(ctx context.Context, req *pb.TechniqueReq) (*pb.AllTechnique, error) {
	return t.Technique.Technique().GetAll(ctx, req)
}

func (t *TechniqueService) Add(ctx context.Context, req *pb.TechniqueAddSub) (*pb.Void, error) {
	return t.Technique.Technique().Add(ctx, req)
}

func (t *TechniqueService) Sub(ctx context.Context, req *pb.TechniqueAddSub) (*pb.Void, error) {
	return t.Technique.Technique().Sub(ctx, req)
}
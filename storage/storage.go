package storage

import (
	"context"

	pb "github.com/Military/military-service/genprotos/militaries"
)

type StorageI interface {
	Bullet() BulletI
	Fuel() FuelI
	Technique() TechniqueI
}

type BulletI interface {
	Create(ctx context.Context, req *pb.BulletReq) (*pb.Void, error)
	Update(ctx context.Context, req *pb.Bullet) (*pb.Void, error)
	Delete(ctx context.Context, req *pb.ById) (*pb.Void, error)
	Get(ctx context.Context, req *pb.ById) (*pb.Bullet, error)
	GetAll(ctx context.Context, req *pb.BulletReq) (*pb.AllBullets, error)
}

type FuelI interface {
	Create(ctx context.Context, req *pb.FuelReq) (*pb.Void, error)
	Update(ctx context.Context, req *pb.Fuel) (*pb.Void, error)
	Delete(ctx context.Context, req *pb.ById) (*pb.Void, error)
	Get(ctx context.Context, req *pb.ById) (*pb.Fuel, error)
	GetAll(ctx context.Context, req *pb.FuelReq) (*pb.AllFuels, error)
}

type TechniqueI interface {
	Create(ctx context.Context, req *pb.TechniqueReq) (*pb.Void, error)
	Update(ctx context.Context, req *pb.Technique) (*pb.Void, error)
	Delete(ctx context.Context, req *pb.ById) (*pb.Void, error)
	Get(ctx context.Context, req *pb.ById) (*pb.Technique, error)
	GetAll(ctx context.Context, req *pb.TechniqueReq) (*pb.AllTechnique, error)
}

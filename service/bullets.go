package service

import (
	"context"

	pb "github.com/Military/military-service/genprotos/militaries"
	"github.com/Military/military-service/storage"
)

type BulletService struct {
	Bullet storage.StorageI
	pb.UnimplementedBulletServiceServer
}

func NewBulletService(bullet storage.StorageI) *BulletService {
	return &BulletService{
		Bullet: bullet,
	}
}

func (b *BulletService) Create(ctx context.Context, req *pb.BulletReq) (*pb.Void, error) {
	return b.Bullet.Bullet().Create(ctx, req)
}

func (b *BulletService) Update(ctx context.Context, req *pb.Bullet) (*pb.Void, error) {
	return b.Bullet.Bullet().Update(ctx, req)
}

func (b *BulletService) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	return b.Bullet.Bullet().Delete(ctx, req)
}

func (b *BulletService) Get(ctx context.Context, req *pb.ById) (*pb.Bullet, error) {
	return b.Bullet.Bullet().Get(ctx, req)
}

func (b *BulletService) GetAll(ctx context.Context, req *pb.BulletReq) (*pb.AllBullets, error) {
	return b.Bullet.Bullet().GetAll(ctx, req)
}

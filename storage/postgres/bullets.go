package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/Military/military-service/genprotos/militaries"
	"github.com/Military/military-service/helper"
	"github.com/google/uuid"
)

type Bullet struct {
	db *sql.DB
}

func NewBullet(db *sql.DB) *Bullet {
	return &Bullet{db: db}
}

func (b *Bullet) Create(ctx context.Context, req *pb.BulletReq) (*pb.Void, error) {
	id := uuid.New().String()
	query := `
	INSERT INTO bullets(id, caliber, type, quantity)
	VALUES($1, $2, $3, $4)
	`
	_, err := b.db.ExecContext(ctx, query, id, req.Caliber, req.Type, req.Quantity)
	if err != nil {
		log.Printf("Error while creating bullet: %v", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (b *Bullet) Update(ctx context.Context, req *pb.Bullet) (*pb.Void, error) {
	query := `
	UPDATE bullets
	SET caliber = $2, type = $3, quantity = $4
	WHERE id = $1
	`
	_, err := b.db.ExecContext(ctx, query, req.Id, req.Caliber, req.Type, req.Quantity)
	if err != nil {
		log.Fatal("error while updating bullet")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (b *Bullet) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE from bullets
		SET deleted_at = $1
	WHERE id = $2
	`
	_, err := b.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Fatal("error while deleting bullet")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (b *Bullet) Get(ctx context.Context, req *pb.ById) (*pb.Bullet, error) {
	query := `
	SELECT id, caliber, type, quantity
	FROM bullets
	WHERE id = $1
	`
	_, err := b.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		log.Fatal("error while getting bullet")
		return nil, err
	}
	return &pb.Bullet{}, nil
}

func (b *Bullet) GetAll(ctx context.Context, req *pb.BulletReq) (*pb.AllBullets, error){
	query := `
	SELECT id, caliber, type, quantity
	FROM bullets
	`
	param := make(map[string]interface{})
	filter := `where deleted_at = 0`
	if req.Caliber != 0 {
		param["caliber"] = req.Caliber
		filter += ` and caliber = :caliber`
	}
	if req.Type != "" {
		param["type"] = req.Type
		filter += ` and type = :type`
	}
	if req.Quantity != 0 {
		param["quantity"] = req.Quantity
		filter += ` and quantity = :quantity`
	}
	query += filter

	query, arr := helper.ReplaceQueryParams(query, param)
	rows, err := b.db.QueryContext(ctx, query, arr...)
	if err != nil {
		log.Fatal("error while getting all bullets")
		return nil, err
	}
	defer rows.Close()

	var bullets []*pb.Bullet
	for rows.Next() {
		var bullet pb.Bullet
		err = rows.Scan(&bullet.Id, &bullet.Caliber, &bullet.Type, &bullet.Quantity)
		if err != nil {
			log.Fatal("error while scanning bullet")
			return nil, err
		}
		bullets = append(bullets, &bullet)
	}
	return &pb.AllBullets{Bullets: bullets}, nil
	



}

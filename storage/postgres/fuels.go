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

type Fuel struct {
	db *sql.DB
}

func NewFuel(db *sql.DB) *Fuel {
	return &Fuel{db: db}
}

func (f *Fuel) Create(ctx context.Context, req *pb.FuelReq) (*pb.Void, error) {
	id := uuid.New().String()
	query := `
	INSERT INTO fuels(id, type, quantity)
	VALUES($1, $2, $3)
	`
	_, err := f.db.ExecContext(ctx, query, id, req.Type, req.Quantity)
	if err != nil {
		log.Fatal("error while creating fuel")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (f *Fuel) Update(ctx context.Context, req *pb.Fuel) (*pb.Void, error) {
	query := `
	UPDATE fuels
	SET type = $2, quantity = $3
	WHERE id = $1
	`
	_, err := f.db.ExecContext(ctx, query, req.Id, req.Type, req.Quantity)
	if err != nil {
		log.Fatal("error while updating fuel")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (f *Fuel) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE fuels
	SET deleted_at = $1
	WHERE id = $2
	`
	_, err := f.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Fatal("error while deleting fuel")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (f *Fuel) Get(ctx context.Context, req *pb.ById) (*pb.Fuel, error) {
	query := `
	SELECT id, type, quantity
	FROM fuels
	WHERE id = $1
	`
	var feul pb.Fuel
	err := f.db.QueryRowContext(ctx, query, req.Id).Scan(&feul.Id, &feul.Type, &feul.Quantity)
	if err != nil {
		log.Fatal("error while getting fuel")
		return nil, err
	}
	return &feul, nil
}

func (f *Fuel) GetAll(ctx context.Context, req *pb.FuelReq) (*pb.AllFuels, error) {

	query := `
	SELECT id, type, quantity
	FROM fuels 
	`
	param := make(map[string]interface{})
	filter := `where deleted_at = 0`
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
	rows, err := f.db.QueryContext(ctx, query, arr...)
	if err != nil {
		log.Fatal("error while getting all bullets")
		return nil, err
	}
	defer rows.Close()

	var fuels []*pb.Fuel
	for rows.Next() {
		var fuel pb.Fuel
		err = rows.Scan(&fuel.Id, &fuel.Type, &fuel.Quantity)
		if err != nil {
			log.Fatal("error while scanning bullet")
			return nil, err
		}
		fuels = append(fuels, &fuel)
	}
	return &pb.AllFuels{Fuels: fuels}, nil

}

func (b *Fuel) Add(ctx context.Context, req *pb.FuelAddSub) (*pb.Void, error) {
	query := `
	UPDATE fuels
	SET quantity = quantity + $2
	WHERE type = $1
	`
	_, err := b.db.ExecContext(ctx, query, req.Name, req.Quantity)
	if err != nil {
		log.Fatal("error while updating fuel")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (b *Fuel) Sub(ctx context.Context, req *pb.FuelAddSub) (*pb.Void, error) {
	query := `
	UPDATE fuels
	SET quantity = quantity - $2
	WHERE type = $1
	`
	_, err := b.db.ExecContext(ctx, query, req.Name, req.Quantity)
	if err != nil {
		log.Fatal("error while updating fuel")
		return nil, err
	}
	return &pb.Void{}, nil
}
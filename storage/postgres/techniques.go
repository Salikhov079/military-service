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

type Technique struct {
	db *sql.DB
}

func NewTechnique(db *sql.DB) *Technique {
	return &Technique{db: db}
}

func (t *Technique) Create(ctx context.Context, req *pb.TechniqueReq) (*pb.Void, error) {
	id := uuid.New().String()
	query := `
	INSERT INTO techniques(id, model, type, quantity)
	VALUES($1, $2, $3, $4)
	`
	_, err := t.db.ExecContext(ctx, query, id, req.Model, req.Type, req.Quantity)
	if err != nil {
		log.Fatal("error while creating technique")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (t *Technique) Update(ctx context.Context, req *pb.Technique) (*pb.Void, error) {
	query := `
	UPDATE techniques
	SET model = $2, type = $3, quantity = $4
	WHERE id = $1
	`
	_, err := t.db.ExecContext(ctx, query, req.Id, req.Model, req.Type, req.Quantity)
	if err != nil {
		log.Fatal("error while updating technique")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (t *Technique) Delete(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE techniques
	SET deleted_at = $1
	WHERE id = $2
	`
	_, err := t.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		log.Fatal("error while deleting technique")
		return nil, err
	}
	return &pb.Void{}, nil
}

func (t *Technique) Get(ctx context.Context, req *pb.ById) (*pb.Technique, error) {
	query := `
	SELECT id, model, type, quantity
	FROM techniques
	WHERE id = $1
	`
	_, err := t.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		log.Fatal("error while getting technique")
		return nil, err
	}
	return &pb.Technique{}, nil
}


func (t *Technique) GetAll(ctx context.Context, req *pb.TechniqueReq) (*pb.AllTechnique, error) {
	query := `
	SELECT id, model, type, quantity
	FROM techniques
	`
	param := make(map[string]interface{})
	filter := `WHERE deleted_at = 0`
	if req.Type != "" {
		param["type"] = req.Type
		filter += ` AND type = :type`
	}
	if req.Quantity != 0 {
		param["quantity"] = req.Quantity
		filter += ` AND quantity = :quantity`
	}
	if req.Model != "" {
		param["model"] = req.Model
		filter += ` AND model = :model`
	}
	query += filter

	query, args := helper.ReplaceQueryParams(query, param)
	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Error while getting all techniques: %v", err)
		return nil, err
	}
	defer rows.Close()

	var techniques []*pb.Technique
	for rows.Next() {
		var technique pb.Technique
		if err = rows.Scan(&technique.Id, &technique.Model, &technique.Type, &technique.Quantity); err != nil {
			log.Printf("Error while scanning technique: %v", err)
			return nil, err
		}
		techniques = append(techniques, &technique)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating rows: %v", err)
		return nil, err
	}

	return &pb.AllTechnique{
		Techniques: techniques,
	}, nil
}
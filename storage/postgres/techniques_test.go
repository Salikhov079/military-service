package postgres

import (
	"context"
	"database/sql"
	"testing"

	pb "github.com/Military/military-service/genprotos/militaries"
)

func TestCreateTechnique(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	technique := NewTechnique(db)
	_, err = technique.Create(context.Background(), &pb.TechniqueReq{
		Model:    "T-72",
		Type:     "weapon",
		Quantity: 10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestCreate passed")

}

func TestUpdateTechnique(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	technique := NewTechnique(db)
	_, err = technique.Update(context.Background(), &pb.Technique{
		Id:       "e79735af-cea1-4701-b9f5-a0b83e49ee37",
		Model:    "T-90",
		Type:     "weapon",
		Quantity: 15,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestUpdate passed")
}

func TestDeleteTechnique(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	technique := NewTechnique(db)
	_, err = technique.Delete(context.Background(), &pb.ById{
		Id: "e79735af-cea1-4701-b9f5-a0b83e49ee37",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestDelete passed")
}

func TestGetTechnique(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	technique := NewTechnique(db)
	_, err = technique.Get(context.Background(), &pb.ById{
		Id: "e79735af-cea1-4701-b9f5-a0b83e49ee37",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestGet passed")
}

func TestGetAllTechnique(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	technique := NewTechnique(db)
	_, err = technique.GetAll(context.Background(), &pb.TechniqueReq{
		Type: "weapon",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestGetAll passed")
}

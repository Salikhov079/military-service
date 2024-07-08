package postgres

import (
	"context"
	"database/sql"
	"testing"

	pb "github.com/Military/military-service/genprotos/militaries"
)

func TestCreateFuel(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	f := NewFuel(db)
	_, err = f.Create(context.Background(), &pb.FuelReq{
		Type:     "diesel",
		Quantity: 1000,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestCreate passed")

}

func TestUpdateFuel(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	f := NewFuel(db)
	_, err = f.Update(context.Background(), &pb.Fuel{
		Id:       "702a0418-5208-4cd5-bd13-4b0c413728df",
		Type:     "petrol",
		Quantity: 1500,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestUpdate passed")
}

func TestDeleteFuel(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	f := NewFuel(db)
	_, err = f.Delete(context.Background(), &pb.ById{
		Id: "702a0418-5208-4cd5-bd13-4b0c413728df",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestDelete passed")
}

func TestGetFuel(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	f := NewFuel(db)
	_, err = f.Get(context.Background(), &pb.ById{
		Id: "702a0418-5208-4cd5-bd13-4b0c413728df",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestGet passed")
}

func TestGetAllFuel(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	f := NewFuel(db)
	_, err = f.GetAll(context.Background(), &pb.FuelReq{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestGetAll passed")
}

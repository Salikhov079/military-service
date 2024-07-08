package postgres

import (
	"context"
	"database/sql"
	"testing"

	pb "github.com/Military/military-service/genprotos/militaries"
)

func TestCreate(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b := NewBullet(db)
	_, err = b.Create(context.Background(), &pb.BulletReq{
		Caliber:  20.5,
		Type:     "weapon",
		Quantity: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestCreate passed")

}

func TestUpdate(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b := NewBullet(db)
	_, err = b.Update(context.Background(), &pb.Bullet{
		Id:       "aa3b71c4-1f9d-498d-8435-b5ac84a854ee",
		Caliber:  30.5,
		Type:     "weapon",
		Quantity: 7,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestUpdate passed")
}

func TestDelete(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b := NewBullet(db)
	_, err = b.Delete(context.Background(), &pb.ById{
		Id: "aa3b71c4-1f9d-498d-8435-b5ac84a854ee",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestDelete passed")
}

func TestGet(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b := NewBullet(db)
	_, err = b.Get(context.Background(), &pb.ById{
		Id: "aa3b71c4-1f9d-498d-8435-b5ac84a854ee",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestGet passed")
}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=20005 dbname=military")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b := NewBullet(db)
	_, err = b.GetAll(context.Background(), &pb.BulletReq{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestGetAll passed")
}

package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Military/military-service/config"
	"github.com/Military/military-service/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
	Bullets storage.BulletI
	Fuels storage.FuelI
	Technoques storage.TechniqueI
}

func DbConn() (storage.StorageI, error) {

	// Get postgres connection data from .env file
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		uint16(cfg.PostgresPort),
		cfg.PostgresDatabase,
	)
	// Connecting to postgres
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
		return nil, err
	}
	return &Storage{Db: db}, nil
}


func(s *Storage) Bullet() storage.BulletI{
	if s.Bullets == nil {
		s.Bullets = NewBullet(s.Db)
	}
	return s.Bullets
}

func(s *Storage) Fuel() storage.FuelI{
	if s.Fuels == nil {
		s.Fuels = NewFuel(s.Db)
	}
	return s.Fuels
}

func(s *Storage) Technique() storage.TechniqueI{
	if s.Technoques == nil {
		s.Technoques = NewTechnique(s.Db)
	}
	return s.Technoques
}



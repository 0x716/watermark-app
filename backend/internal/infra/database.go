package infra

import (
	"database/sql"
	"fmt"

	"github.com/0x716/watermark-app/internal/config"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func InitDB() error {
	cfg := config.GlobalConfig.Database

	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", config.GlobalConfig.Database.Engine, config.GlobalConfig.Database.User, config.GlobalConfig.Database.Password, config.GlobalConfig.Database.Host, config.GlobalConfig.Database.Port, config.GlobalConfig.Database.Name, config.GlobalConfig.Database.Sslmode)

	var err error
	DB, err = sql.Open(cfg.Title, connStr)
	if err != nil {
		return fmt.Errorf("DB Connect Failed: %v", err)
	}

	return nil
}

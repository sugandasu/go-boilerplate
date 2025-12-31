package migration

import (
	"time"

	"github.com/sugandasu/ruru/sunjumig"
	"gorm.io/gorm"
)

func Init20250213144137() sunjumig.Migration {
	return sunjumig.Migration{
		ID:        0,
		Name:      "20250213144137_create_roles_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250213144137_create_roles_table_up,
		Down:      mig_20250213144137_create_roles_table_down,
	}
}

func mig_20250213144137_create_roles_table_up(tx *gorm.DB) error {
	err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS roles (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) UNIQUE,
		level INT,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		created_by VARCHAR(255),
		updated_at TIMESTAMPTZ NULL,
		updated_by VARCHAR(255),
		deleted_at TIMESTAMPTZ NULL,
		deleted_by VARCHAR(255)
	);`).Error

	return err
}

func mig_20250213144137_create_roles_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS roles;`).Error

	return err
}

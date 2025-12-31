package migration

import (
	"time"

	"github.com/sugandasu/ruru/sunjumig"
	"gorm.io/gorm"
)

func Init20250214172046() sunjumig.Migration {
	return sunjumig.Migration{
		ID:        0,
		Name:      "20250214172046_create_users_table",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20250214172046_create_users_table_up,
		Down:      mig_20250214172046_create_users_table_down,
	}
}

func mig_20250214172046_create_users_table_up(tx *gorm.DB) error {
	err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		role_id VARCHAR(255) NULL REFERENCES roles(id) ON DELETE RESTRICT ON UPDATE RESTRICT,
		name VARCHAR(255),
		username VARCHAR(255) UNIQUE,
		email VARCHAR(255) UNIQUE,
		password VARCHAR(255),
		phone_number VARCHAR(255) UNIQUE,
		status VARCHAR(255),
		created_at TIMESTAMPTZ DEFAULT NOW(),
		created_by VARCHAR(255),
		updated_at TIMESTAMPTZ NULL,
		updated_by VARCHAR(255) NULL,
		deleted_at TIMESTAMPTZ NULL,
		deleted_by VARCHAR(255) NULL
	);`).Error

	return err
}

func mig_20250214172046_create_users_table_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS users;`).Error

	return err
}

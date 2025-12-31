package migration

import (
	"time"

	"github.com/sugandasu/ruru/sunjumig"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Init21000218144856() sunjumig.Migration {
	return sunjumig.Migration{
		ID:        0,
		Name:      "21000218144856_insert_admin_users",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_21000218144856_insert_admin_users_up,
		Down:      mig_21000218144856_insert_admin_users_down,
	}
}

func mig_21000218144856_insert_admin_users_up(tx *gorm.DB) error {
	now := time.Now()

	roleId := "01KBA6RMARZ28MT6MWTNQE68F4"
	userID := "01KBA6QN20AQFMXQJ88C97FPW0"

	err := tx.Exec(`INSERT INTO roles (id, name, level, created_by) VALUES (?, ?, ?, ?);`,
		roleId,
		"admin",
		0,
		"default",
	).Error
	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), 0)
	err = tx.Exec(`
	INSERT INTO users (id, name, username, email, password, phone_number, role_id, status, created_by, created_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		userID,
		"admin",
		"admin",
		"youremail@gmail.com",
		string(password),
		"081234567890",
		roleId,
		"active",
		"default",
		now,
	).Error
	if err != nil {
		return err
	}

	return err
}

func mig_21000218144856_insert_admin_users_down(tx *gorm.DB) error {
	err := tx.Exec(`DELETE FROM users WHERE username = ?`, "admin").Error
	if err != nil {
		return err
	}
	err = tx.Exec(`DELETE FROM roles WHERE name = ?`, "admin").Error
	if err != nil {
		return err
	}

	return nil
}

package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sugandasu/go-boilerplate/config"
	"github.com/sugandasu/go-boilerplate/migration"
	"github.com/sugandasu/ruru/nibirudb"
	"github.com/sugandasu/ruru/sunjumig"
	"gorm.io/gorm"
)

func initDB(cmd *cobra.Command) (db *gorm.DB, migrations []sunjumig.Migration) {
	conn, _ := cmd.Flags().GetString("connection")
	if conn == "" || conn == "default" {
		cfg := config.Load()
		db = nibirudb.NewDatabaseConnection(&cfg.DB).DB(context.Background())
		migrations = append(migrations,
			migration.Init20250213144137(),
			migration.Init20250214172046(),
			migration.Init21000218144856(),
		)
	}

	return
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migration file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := sunjumig.Create(name); err != nil {
			log.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator, err := sunjumig.Init(initDB(cmd))
		if err != nil {
			log.Println("Unable to fetch migrator", err.Error())
			return
		}

		log.Println("Running migration...")
		err = migrator.Up()
		if err != nil {
			log.Println("Unable to run `up` migrations", err.Error())
			return
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator, err := sunjumig.Init(initDB(cmd))
		if err != nil {
			log.Println("Unable to fetch migrator", err.Error())
			return
		}

		err = migrator.Down()
		if err != nil {
			log.Println("Unable to run `down` migrations")
			return
		}
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrator, err := sunjumig.Init(initDB(cmd))
		if err != nil {
			log.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			log.Println("Unable to fetch migration status")
			return
		}
	},
}

func initMigrate() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Fatalf("unknown command %q", strings.Join(args, " "))
			}

			_ = cmd.Help()
		},
	}

	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")
	_ = migrateCreateCmd.MarkFlagRequired("name")

	migrateUpCmd.Flags().StringP("connection", "c", "", "DB connection for the migration")
	migrateDownCmd.Flags().StringP("connection", "c", "", "DB connection for the migration")
	migrateStatusCmd.Flags().StringP("connection", "c", "", "DB connection for the migration")

	migrateCmd.AddCommand(migrateCreateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateStatusCmd)

	return migrateCmd
}

func init() {
	rootCmd.AddCommand(initMigrate())
}

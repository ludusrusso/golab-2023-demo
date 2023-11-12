package dbchema

import (
	"embed"
	"fmt"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate(dbURL string) error {
	u, err := url.Parse(dbURL)
	if err != nil {
		return fmt.Errorf("invalid database url: %w", err)
	}
	db := dbmate.New(u)
	db.FS = fs
	db.MigrationsDir = []string{"migrations"}
	db.AutoDumpSchema = false

	err = db.CreateAndMigrate()
	if err != nil {
		return fmt.Errorf("cannot migrate database: %w", err)
	}

	return nil
}

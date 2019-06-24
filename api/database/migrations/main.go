package migrations

import (
	"app/models"
)

func Run() {

	(models.Token{}).Migrate()
	(models.Kanji{}).Migrate()
}

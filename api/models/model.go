package models

type Model interface {
	Migrate()
	Seed()
}

func Prepare(model Model) {

	model.Migrate()
	model.Seed()
}

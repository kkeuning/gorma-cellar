package design

import (
	cellar "github.com/goadesign/goa-cellar/design"
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("Cellar", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("Account", func() {
			RendersTo(cellar.Account)
			Description("Cellar Account")
			HasMany("Bottles", "Bottle")
		})

		Model("Bottle", func() {
			BuildsFrom(func() {
				Payload("bottle", "create")
				Payload("bottle", "update")
			})
			RendersTo(cellar.Bottle)
			Description("Bottle Model")
		})

	})
})

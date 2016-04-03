package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma-cellar/app"
	"github.com/goadesign/gorma-cellar/models"
	"github.com/jinzhu/gorm"
)

// ErrDatabaseError is the error returned when a db query fails.
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("account")}
}

// Create runs the create action.
func (c *AccountController) Create(ctx *app.CreateAccountContext) error {
	a := models.Account{}
	a.Name = ctx.Payload.Name
	ra, err := adb.Add(ctx.Context, &a)
	if err != nil {
		return ErrDatabaseError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.AccountHref(ra.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	err := adb.Delete(ctx.Context, ctx.AccountID)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	account, err := adb.OneAccount(ctx.Context, ctx.AccountID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	account.Href = app.AccountHref(account.ID)
	return ctx.OK(account)
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	m, err := adb.Get(ctx.Context, ctx.AccountID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	m.Name = ctx.Payload.Name
	err = adb.Update(ctx, &m)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

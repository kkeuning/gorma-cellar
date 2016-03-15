package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma-cellar/app"
)

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
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	// TBD: implement
	return nil
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	// TBD: implement
	res := &app.Account{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	// TBD: implement
	return nil
}

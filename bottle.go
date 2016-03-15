package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/gorma-cellar/app"
	"golang.org/x/net/websocket"
	"io"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("bottle")}
}

// Create runs the create action.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *BottleController) Delete(ctx *app.DeleteBottleContext) error {
	// TBD: implement
	return nil
}

// List runs the list action.
func (c *BottleController) List(ctx *app.ListBottleContext) error {
	// TBD: implement
	res := app.BottleCollection{}
	return ctx.OK(res)
}

// Rate runs the rate action.
func (c *BottleController) Rate(ctx *app.RateBottleContext) error {
	// TBD: implement
	return nil
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// TBD: implement
	res := &app.Bottle{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *BottleController) Update(ctx *app.UpdateBottleContext) error {
	// TBD: implement
	return nil
}

// Watch runs the watch action.
func (c *BottleController) Watch(ctx *app.WatchBottleContext) error {
	c.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (c *BottleController) WatchWSHandler(ctx *app.WatchBottleContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// TBD: implement
		ws.Write([]byte("watch bottle"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
	}
}

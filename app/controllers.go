//************************************************************************//
// API "cellar": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/gorma-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/gorma-cellar
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AccountController is the controller interface for the Account actions.
type AccountController interface {
	goa.Muxer
	Create(*CreateAccountContext) error
	Delete(*DeleteAccountContext) error
	List(*ListAccountContext) error
	Show(*ShowAccountContext) error
	Update(*UpdateAccountContext) error
}

// MountAccountController "mounts" a Account resource controller on the given service.
func MountAccountController(service *goa.Service, ctrl AccountController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/cellar/accounts", ctrl.MuxHandler("preflight", handleAccountOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/cellar/accounts/:accountID", ctrl.MuxHandler("preflight", handleAccountOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateAccountContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateAccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("POST", "/cellar/accounts", ctrl.MuxHandler("Create", h, unmarshalCreateAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Create", "route", "POST /cellar/accounts")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("DELETE", "/cellar/accounts/:accountID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "Delete", "route", "DELETE /cellar/accounts/:accountID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("GET", "/cellar/accounts", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "List", "route", "GET /cellar/accounts")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowAccountContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("GET", "/cellar/accounts/:accountID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Account", "action", "Show", "route", "GET /cellar/accounts/:accountID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateAccountContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateAccountPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleAccountOrigin(h)
	service.Mux.Handle("PUT", "/cellar/accounts/:accountID", ctrl.MuxHandler("Update", h, unmarshalUpdateAccountPayload))
	service.LogInfo("mount", "ctrl", "Account", "action", "Update", "route", "PUT /cellar/accounts/:accountID")
}

// handleAccountOrigin applies the CORS response headers corresponding to the origin.
func handleAccountOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createAccountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateAccountPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAccountPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateAccountPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// BottleController is the controller interface for the Bottle actions.
type BottleController interface {
	goa.Muxer
	Create(*CreateBottleContext) error
	Delete(*DeleteBottleContext) error
	List(*ListBottleContext) error
	Rate(*RateBottleContext) error
	Show(*ShowBottleContext) error
	Update(*UpdateBottleContext) error
	Watch(*WatchBottleContext) error
}

// MountBottleController "mounts" a Bottle resource controller on the given service.
func MountBottleController(service *goa.Service, ctrl BottleController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("preflight", handleBottleOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("preflight", handleBottleOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/cellar/accounts/:accountID/bottles/:bottleID/actions/rate", ctrl.MuxHandler("preflight", handleBottleOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/cellar/accounts/:accountID/bottles/:bottleID/watch", ctrl.MuxHandler("preflight", handleBottleOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateBottleContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateBottlePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("POST", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("Create", h, unmarshalCreateBottlePayload))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Create", "route", "POST /cellar/accounts/:accountID/bottles")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteBottleContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("DELETE", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Delete", "route", "DELETE /cellar/accounts/:accountID/bottles/:bottleID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListBottleContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("GET", "/cellar/accounts/:accountID/bottles", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "List", "route", "GET /cellar/accounts/:accountID/bottles")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRateBottleContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RateBottlePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Rate(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("PUT", "/cellar/accounts/:accountID/bottles/:bottleID/actions/rate", ctrl.MuxHandler("Rate", h, unmarshalRateBottlePayload))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Rate", "route", "PUT /cellar/accounts/:accountID/bottles/:bottleID/actions/rate")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowBottleContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("GET", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Show", "route", "GET /cellar/accounts/:accountID/bottles/:bottleID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateBottleContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateBottlePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("PATCH", "/cellar/accounts/:accountID/bottles/:bottleID", ctrl.MuxHandler("Update", h, unmarshalUpdateBottlePayload))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Update", "route", "PATCH /cellar/accounts/:accountID/bottles/:bottleID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewWatchBottleContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Watch(rctx)
	}
	h = handleBottleOrigin(h)
	service.Mux.Handle("GET", "/cellar/accounts/:accountID/bottles/:bottleID/watch", ctrl.MuxHandler("Watch", h, nil))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Watch", "route", "GET /cellar/accounts/:accountID/bottles/:bottleID/watch")
}

// handleBottleOrigin applies the CORS response headers corresponding to the origin.
func handleBottleOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateBottlePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createBottlePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalRateBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalRateBottlePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &rateBottlePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateBottlePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateBottlePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateBottlePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Health(*HealthHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/cellar/_ah/health", ctrl.MuxHandler("preflight", handleHealthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthHealthContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	h = handleHealthOrigin(h)
	service.Mux.Handle("GET", "/cellar/_ah/health", ctrl.MuxHandler("Health", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Health", "route", "GET /cellar/_ah/health")
}

// handleHealthOrigin applies the CORS response headers corresponding to the origin.
func handleHealthOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger.json", "public/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

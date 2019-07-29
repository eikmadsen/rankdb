// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "rankdb": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/Vivino/rankdb/api/design
// --out=$(GOPATH)/src/github.com/Vivino/rankdb/api
// --version=v1.3.1

package app

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"github.com/goadesign/goa/encoding/msgpack"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(msgpack.NewEncoder, "application/msgpack")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(msgpack.NewDecoder, "application/msgpack")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// BackupController is the controller interface for the Backup actions.
type BackupController interface {
	goa.Muxer
	Delete(*DeleteBackupContext) error
	Status(*StatusBackupContext) error
}

// MountBackupController "mounts" a Backup resource controller on the given service.
func MountBackupController(service *goa.Service, ctrl BackupController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteBackupContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("DELETE", "/backup/:backup_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Backup", "action", "Delete", "route", "DELETE /backup/:backup_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStatusBackupContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Status(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("GET", "/backup/:backup_id", ctrl.MuxHandler("status", h, nil))
	service.LogInfo("mount", "ctrl", "Backup", "action", "Status", "route", "GET /backup/:backup_id", "security", "jwt")
}

// ElementsController is the controller interface for the Elements actions.
type ElementsController interface {
	goa.Muxer
	Create(*CreateElementsContext) error
	Delete(*DeleteElementsContext) error
	DeleteMulti(*DeleteMultiElementsContext) error
	Get(*GetElementsContext) error
	GetAround(*GetAroundElementsContext) error
	GetMulti(*GetMultiElementsContext) error
	Put(*PutElementsContext) error
	PutMulti(*PutMultiElementsContext) error
}

// MountElementsController "mounts" a Elements resource controller on the given service.
func MountElementsController(service *goa.Service, ctrl ElementsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Element)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h, "api:update")
	service.Mux.Handle("POST", "/lists/:list_id/elements", ctrl.MuxHandler("create", h, unmarshalCreateElementsPayload))
	service.LogInfo("mount", "ctrl", "Elements", "action", "Create", "route", "POST /lists/:list_id/elements", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:delete")
	service.Mux.Handle("DELETE", "/lists/:list_id/elements/:element_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Elements", "action", "Delete", "route", "DELETE /lists/:list_id/elements/:element_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteMultiElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteMulti(rctx)
	}
	h = handleSecurity("jwt", h, "api:delete")
	service.Mux.Handle("DELETE", "/lists/:list_id/elements", ctrl.MuxHandler("delete-multi", h, nil))
	service.LogInfo("mount", "ctrl", "Elements", "action", "DeleteMulti", "route", "DELETE /lists/:list_id/elements", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("GET", "/lists/:list_id/elements/:element_id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Elements", "action", "Get", "route", "GET /lists/:list_id/elements/:element_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetAroundElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*MultiElement)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.GetAround(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("POST", "/lists/:list_id/elements/:element_id/around", ctrl.MuxHandler("get-around", h, unmarshalGetAroundElementsPayload))
	service.LogInfo("mount", "ctrl", "Elements", "action", "GetAround", "route", "POST /lists/:list_id/elements/:element_id/around", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetMultiElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*MultiElement)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.GetMulti(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("POST", "/lists/:list_id/elements/find", ctrl.MuxHandler("get-multi", h, unmarshalGetMultiElementsPayload))
	service.LogInfo("mount", "ctrl", "Elements", "action", "GetMulti", "route", "POST /lists/:list_id/elements/find", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPutElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Element)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Put(rctx)
	}
	h = handleSecurity("jwt", h, "api:update")
	service.Mux.Handle("PUT", "/lists/:list_id/elements/:element_id", ctrl.MuxHandler("put", h, unmarshalPutElementsPayload))
	service.LogInfo("mount", "ctrl", "Elements", "action", "Put", "route", "PUT /lists/:list_id/elements/:element_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPutMultiElementsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(PutMultiElementsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.PutMulti(rctx)
	}
	h = handleSecurity("jwt", h, "api:update")
	service.Mux.Handle("PUT", "/lists/:list_id/elements", ctrl.MuxHandler("put-multi", h, unmarshalPutMultiElementsPayload))
	service.LogInfo("mount", "ctrl", "Elements", "action", "PutMulti", "route", "PUT /lists/:list_id/elements", "security", "jwt")
}

// unmarshalCreateElementsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateElementsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &element{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalGetAroundElementsPayload unmarshals the request body into the context request data Payload field.
func unmarshalGetAroundElementsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &multiElement{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalGetMultiElementsPayload unmarshals the request body into the context request data Payload field.
func unmarshalGetMultiElementsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &multiElement{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalPutElementsPayload unmarshals the request body into the context request data Payload field.
func unmarshalPutElementsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &element{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalPutMultiElementsPayload unmarshals the request body into the context request data Payload field.
func unmarshalPutMultiElementsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	var payload PutMultiElementsPayload
	if err := service.DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload
	return nil
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Health(*HealthHealthContext) error
	Root(*RootHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("GET", "/health", ctrl.MuxHandler("health", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Health", "route", "GET /health", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRootHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Root(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("root", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Root", "route", "GET /", "security", "jwt")
}

// JWTController is the controller interface for the JWT actions.
type JWTController interface {
	goa.Muxer
	JWT(*JWTJWTContext) error
}

// MountJWTController "mounts" a JWT resource controller on the given service.
func MountJWTController(service *goa.Service, ctrl JWTController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewJWTJWTContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.JWT(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("POST", "/jwt", ctrl.MuxHandler("jwt", h, nil))
	service.LogInfo("mount", "ctrl", "JWT", "action", "JWT", "route", "POST /jwt", "security", "jwt")
}

// ListsController is the controller interface for the Lists actions.
type ListsController interface {
	goa.Muxer
	Clone(*CloneListsContext) error
	Create(*CreateListsContext) error
	Delete(*DeleteListsContext) error
	Get(*GetListsContext) error
	GetPercentile(*GetPercentileListsContext) error
	GetRange(*GetRangeListsContext) error
	GetAll(*GetAllListsContext) error
	Reindex(*ReindexListsContext) error
	Repair(*RepairListsContext) error
	Verify(*VerifyListsContext) error
}

// MountListsController "mounts" a Lists resource controller on the given service.
func MountListsController(service *goa.Service, ctrl ListsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCloneListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RankList)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Clone(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/lists/:list_id/clone", ctrl.MuxHandler("clone", h, unmarshalCloneListsPayload))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Clone", "route", "PUT /lists/:list_id/clone", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RankList)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("POST", "/lists", ctrl.MuxHandler("create", h, unmarshalCreateListsPayload))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Create", "route", "POST /lists", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("DELETE", "/lists/:list_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Delete", "route", "DELETE /lists/:list_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("GET", "/lists/:list_id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Get", "route", "GET /lists/:list_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetPercentileListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetPercentile(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("GET", "/lists/:list_id/percentile", ctrl.MuxHandler("get-percentile", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "GetPercentile", "route", "GET /lists/:list_id/percentile", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetRangeListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetRange(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("GET", "/lists/:list_id/range", ctrl.MuxHandler("get-range", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "GetRange", "route", "GET /lists/:list_id/range", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetAllListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetAll(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("GET", "/lists", ctrl.MuxHandler("get_all", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "GetAll", "route", "GET /lists", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReindexListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Reindex(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/lists/:list_id/reindex", ctrl.MuxHandler("reindex", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Reindex", "route", "PUT /lists/:list_id/reindex", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRepairListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Repair(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/lists/:list_id/repair", ctrl.MuxHandler("repair", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Repair", "route", "PUT /lists/:list_id/repair", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVerifyListsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Verify(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/lists/:list_id/verify", ctrl.MuxHandler("verify", h, nil))
	service.LogInfo("mount", "ctrl", "Lists", "action", "Verify", "route", "PUT /lists/:list_id/verify", "security", "jwt")
}

// unmarshalCloneListsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCloneListsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &rankList{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreateListsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateListsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &rankList{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// MultilistController is the controller interface for the Multilist actions.
type MultilistController interface {
	goa.Muxer
	Backup(*BackupMultilistContext) error
	Create(*CreateMultilistContext) error
	Delete(*DeleteMultilistContext) error
	Get(*GetMultilistContext) error
	Put(*PutMultilistContext) error
	Reindex(*ReindexMultilistContext) error
	Restore(*RestoreMultilistContext) error
	Verify(*VerifyMultilistContext) error
}

// MountMultilistController "mounts" a Multilist resource controller on the given service.
func MountMultilistController(service *goa.Service, ctrl MultilistController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewBackupMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*MultiListBackup)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Backup(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/xlist/backup", ctrl.MuxHandler("backup", h, unmarshalBackupMultilistPayload))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Backup", "route", "PUT /xlist/backup", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListPayloadQL)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h, "api:update")
	service.Mux.Handle("POST", "/xlist/elements", ctrl.MuxHandler("create", h, unmarshalCreateMultilistPayload))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Create", "route", "POST /xlist/elements", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:delete")
	service.Mux.Handle("DELETE", "/xlist/elements/:element_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Delete", "route", "DELETE /xlist/elements/:element_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:read")
	service.Mux.Handle("GET", "/xlist/elements/:element_id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Get", "route", "GET /xlist/elements/:element_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPutMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListPayloadQL)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Put(rctx)
	}
	h = handleSecurity("jwt", h, "api:update")
	service.Mux.Handle("PUT", "/xlist/elements", ctrl.MuxHandler("put", h, unmarshalPutMultilistPayload))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Put", "route", "PUT /xlist/elements", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReindexMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListQL)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Reindex(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/xlist/reindex", ctrl.MuxHandler("reindex", h, unmarshalReindexMultilistPayload))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Reindex", "route", "PUT /xlist/reindex", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRestoreMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Restore(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("POST", "/xlist/restore", ctrl.MuxHandler("restore", h, nil))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Restore", "route", "POST /xlist/restore", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVerifyMultilistContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ListQL)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Verify(rctx)
	}
	h = handleSecurity("jwt", h, "api:manage")
	service.Mux.Handle("PUT", "/xlist/verify", ctrl.MuxHandler("verify", h, unmarshalVerifyMultilistPayload))
	service.LogInfo("mount", "ctrl", "Multilist", "action", "Verify", "route", "PUT /xlist/verify", "security", "jwt")
}

// unmarshalBackupMultilistPayload unmarshals the request body into the context request data Payload field.
func unmarshalBackupMultilistPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &multiListBackup{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreateMultilistPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateMultilistPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listPayloadQL{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalPutMultilistPayload unmarshals the request body into the context request data Payload field.
func unmarshalPutMultilistPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listPayloadQL{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalReindexMultilistPayload unmarshals the request body into the context request data Payload field.
func unmarshalReindexMultilistPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listQL{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalVerifyMultilistPayload unmarshals the request body into the context request data Payload field.
func unmarshalVerifyMultilistPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &listQL{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// StaticController is the controller interface for the Static actions.
type StaticController interface {
	goa.Muxer
	goa.FileServer
}

// MountStaticController "mounts" a Static resource controller on the given service.
func MountStaticController(service *goa.Service, ctrl StaticController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/doc/*filepath", ctrl.MuxHandler("preflight", handleStaticOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/swagger/*filepath", ctrl.MuxHandler("preflight", handleStaticOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/doc/*filepath", "api/public/swagger")
	h = handleStaticOrigin(h)
	service.Mux.Handle("GET", "/doc/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Static", "files", "api/public/swagger", "route", "GET /doc/*filepath")

	h = ctrl.FileHandler("/api/swagger/*filepath", "api/swagger")
	h = handleStaticOrigin(h)
	service.Mux.Handle("GET", "/api/swagger/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Static", "files", "api/swagger", "route", "GET /api/swagger/*filepath")

	h = ctrl.FileHandler("/doc/", "api/public/swagger/index.html")
	h = handleStaticOrigin(h)
	service.Mux.Handle("GET", "/doc/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Static", "files", "api/public/swagger/index.html", "route", "GET /doc/")

	h = ctrl.FileHandler("/api/swagger/", "api/swagger/index.html")
	h = handleStaticOrigin(h)
	service.Mux.Handle("GET", "/api/swagger/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Static", "files", "api/swagger/index.html", "route", "GET /api/swagger/")
}

// handleStaticOrigin applies the CORS response headers corresponding to the origin.
func handleStaticOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

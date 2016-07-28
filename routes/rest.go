package routes

import (
	"net/http"

	"golang.org/x/net/context"
)

type CtxHandleFunc func(rw http.ResponseWriter, r *http.Request, ctx context.Context)

type RESTCtxHandler interface {
	Init(rw http.ResponseWriter, r *http.Request, next CtxHandleFunc)
	HandleGet(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandleGetList(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandlePost(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandlePut(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandlePatch(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandleDelete(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandleCopy(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandleHead(rw http.ResponseWriter, r *http.Request, ctx context.Context)
	HandleOptions(rw http.ResponseWriter, r *http.Request, ctx context.Context)
}

// -----------------------
// default handler. embed this in a handler to follow the RESTCtxHandler interface without needing to implement all the
// verbs

type DefaultRESTCtxHandler struct{}

//Init creates the context passed to each handler. Overriding this allows population of the context for all
//verbs
func (h *DefaultRESTCtxHandler) Init(rw http.ResponseWriter, r *http.Request, next CtxHandleFunc) {
	next(rw, r, context.Background())
}
func (h *DefaultRESTCtxHandler) HandleGet(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandleGetList(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandlePost(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandlePut(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandlePatch(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandleDelete(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandleCopy(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandleHead(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTCtxHandler) HandleOptions(rw http.ResponseWriter, r *http.Request, ctx context.Context) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}

type RESTHandler interface {
	HandleGet(rw http.ResponseWriter, r *http.Request)
	HandleGetList(rw http.ResponseWriter, r *http.Request)
	HandlePost(rw http.ResponseWriter, r *http.Request)
	HandlePut(rw http.ResponseWriter, r *http.Request)
	HandlePatch(rw http.ResponseWriter, r *http.Request)
	HandleDelete(rw http.ResponseWriter, r *http.Request)
	HandleCopy(rw http.ResponseWriter, r *http.Request)
	HandleHead(rw http.ResponseWriter, r *http.Request)
	HandleOptions(rw http.ResponseWriter, r *http.Request)
}

// -----------------------
// default handler. embed this in a handler to follow the RESTHandler interface without needing to implement all the
// verbs

type DefaultRESTHandler struct{}

func (h *DefaultRESTHandler) HandleGet(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandleGetList(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandlePost(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandlePut(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandlePatch(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandleDelete(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandleCopy(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandleHead(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}
func (h *DefaultRESTHandler) HandleOptions(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusNotImplemented)
}

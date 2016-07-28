package routes

import "net/http"

func InitCtx(next RESTCtxHandler) RESTHandler {
	return &CtxInitMiddleware{Next: next}
}

//BackgroundCtxMiddleware converts a RESTCtxHandler into a RESTHandler
type CtxInitMiddleware struct {
	Next RESTCtxHandler
}

func (h *CtxInitMiddleware) HandleGet(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandleGet)
}
func (h *CtxInitMiddleware) HandleGetList(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandleGetList)
}
func (h *CtxInitMiddleware) HandlePost(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandlePost)
}
func (h *CtxInitMiddleware) HandlePut(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandlePut)
}
func (h *CtxInitMiddleware) HandlePatch(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandlePatch)
}
func (h *CtxInitMiddleware) HandleDelete(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandleDelete)
}
func (h *CtxInitMiddleware) HandleCopy(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandleCopy)
}
func (h *CtxInitMiddleware) HandleHead(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandleHead)
}
func (h *CtxInitMiddleware) HandleOptions(rw http.ResponseWriter, r *http.Request) {
	h.Next.Init(rw, r, h.Next.HandleOptions)
}

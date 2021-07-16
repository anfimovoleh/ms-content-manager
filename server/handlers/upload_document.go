package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type UploadDocumentHandler struct {
	log *zap.Logger
}

func NewUploadDocumentHandler(log *zap.Logger) *UploadDocumentHandler {
	return &UploadDocumentHandler{log: log}
}

func (h UploadDocumentHandler) Handle(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("started")
	_, _ = w.Write([]byte("upload document"))
	h.log.Debug("finished")
}

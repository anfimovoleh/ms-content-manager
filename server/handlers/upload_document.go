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
	// Maximum upload of 10 MB files
	err := r.ParseMultipartForm(1000 << 20)
	if err != nil {
		h.log.With(zap.Error(err)).Error("failed to parse multipart form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get header for filename, size and headers
	file, header, err := r.FormFile("myFile")
	if err != nil {
		h.log.With(zap.Error(err)).Error("failed to parse multipart form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	h.log.With(
		zap.String("file_name", header.Filename),
		zap.Int64("file_size", header.Size),
		zap.Any("mime_header", header.Header),
	).Debug("uploaded file")
}

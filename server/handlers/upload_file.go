package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type UploadFileHandler struct {
	maxFileSize int64

	log *zap.Logger
}

func NewUploadDocumentHandler(log *zap.Logger) *UploadFileHandler {
	return &UploadFileHandler{log: log}
}

func (h UploadFileHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	err := r.ParseMultipartForm(h.maxFileSize)
	if err != nil {
		h.log.With(zap.Error(err)).Error("failed to parse multipart form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileName := "file"
	// Get header for filename, size and headers
	file, header, err := r.FormFile(fileName)
	if err != nil {
		h.log.With(zap.Error(err)).Error("failed to parse multipart form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = file.Close()
	if err != nil {
		h.log.With(zap.String("file_name", fileName)).Error("failed to close file")
	}

	h.log.With(
		zap.String("file_name", header.Filename),
		zap.Int64("file_size", header.Size),
		zap.Any("mime_header", header.Header),
	).Debug("uploaded file")
}

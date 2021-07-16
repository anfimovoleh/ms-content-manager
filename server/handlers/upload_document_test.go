package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anfimovoleh/ms-content-manager/config"

	"go.uber.org/zap"
)

func TestUploadDocumentHandler_Handle(t *testing.T) {
	type fields struct {
		log *zap.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				log: config.New().Log(),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/upload", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UploadDocumentHandler{
				log: tt.fields.log,
			}

			u.Handle(tt.args.w, tt.args.r)
		})
	}
}

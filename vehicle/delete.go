package vehicle

import (
	"net/http"
	"strconv"

	"github.com/Lzrtn/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var string_id string = r.PathValue("id")
	id, _ := strconv.ParseInt(string_id, 10, 64)
	rep, _ := d.store.Vehicle().Delete(r.Context(), id)

	if rep {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(rw, "404", http.StatusInternalServerError)
	}
}

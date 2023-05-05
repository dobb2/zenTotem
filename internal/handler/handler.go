package handler

import (
	"encoding/json"
	"github.com/dobb2/zenTotem/internal/crypto"
	"github.com/dobb2/zenTotem/internal/entity"
	"github.com/dobb2/zenTotem/internal/storage"
	"github.com/rs/zerolog"
	"net/http"
)

type UserHandler struct {
	storage storage.Storer
	cache   storage.Cacher
	logger  zerolog.Logger
}

func New(storage storage.Storer, cache storage.Cacher, logger zerolog.Logger) UserHandler {
	return UserHandler{
		storage: storage,
		cache:   cache,
		logger:  logger,
	}
}

func (u UserHandler) IncrementVal(w http.ResponseWriter, r *http.Request) {
	var elem entity.Element
	if err := json.NewDecoder(r.Body).Decode(&elem); err != nil {
		u.logger.Debug().Err(err).Msg("invalid json")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if elem.Key == "" || elem.Value == 0 {
		u.logger.Debug().Msg("The json does not match the correct type")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "The json does not match the correct type", http.StatusBadRequest)
		return
	}

	outValue, err := u.cache.Increment(elem)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(outValue)
	if err != nil {
		http.Error(w, "problem marshal incremented value to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (u UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.logger.Debug().Err(err).Msg("invalid json")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age == 0 {
		u.logger.Debug().Msg("The json does not match the correct type")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "The json does not match the correct type", http.StatusBadRequest)
		return
	}

	idUser, err := u.storage.Create(user)
	if err != nil {
		u.logger.Error().Err(err).Msg("cannot create user")
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(idUser)
	if err != nil {
		http.Error(w, "problem marshal metric to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (u UserHandler) PostSign(w http.ResponseWriter, r *http.Request) {
	var getSign entity.Sign
	if err := json.NewDecoder(r.Body).Decode(&getSign); err != nil {
		u.logger.Debug().Err(err).Msg("invalid json")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if getSign.Text == "" {
		u.logger.Debug().Msg("The json does not match the correct type")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "The json does not match the correct type", http.StatusBadRequest)
		return
	}

	var outSign entity.Sign
	outSign.Hex = crypto.SignHMAC(getSign.Text, getSign.Key)

	out, err := json.Marshal(outSign)
	if err != nil {
		u.logger.Error().Err(err).Msg("cannot marsha signt struct to json")
		http.Error(w, "problem marshal sign to json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

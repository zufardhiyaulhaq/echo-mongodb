package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mongodb_client "github.com/zufardhiyaulhaq/echo-mongodb/pkg/mongodb"
)

type Server struct {
	settings settings.Settings
	client   mongodb_client.Interface
}

func NewServer(settings settings.Settings, client mongodb_client.Interface) Server {
	return Server{
		settings: settings,
		client:   client,
	}
}

func (e Server) ServeHTTP() {
	handler := NewHandler(e.settings, e.client)

	r := mux.NewRouter()

	r.HandleFunc("/mongodb/{key}", handler.Handle)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!"))
	})
	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!"))
	})

	err := http.ListenAndServe(":"+e.settings.HTTPPort, r)
	if err != nil {
		log.Fatal().Err(err)
	}
}

type Handler struct {
	settings settings.Settings
	client   mongodb_client.Interface
}

func NewHandler(settings settings.Settings, client mongodb_client.Interface) Handler {
	return Handler{
		settings: settings,
		client:   client,
	}
}

func (h Handler) Handle(w http.ResponseWriter, req *http.Request) {
	value := mux.Vars(req)["key"]
	docID := primitive.NewObjectID()
	echo := types.Echo{
		ID:   docID,
		Echo: value,
	}

	err := h.client.InsertEcho(echo)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	read, err := h.client.GetEcho(docID)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(read.ID.String() + ":" + read.Echo))
}

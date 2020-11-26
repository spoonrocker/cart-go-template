package http

import (
	"cartapi/cartapi"
	"cartapi/cartapi/logger"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Server struct {
	router *Router
}

type ErrorCode int

const (
	EntityNotFound ErrorCode = iota
	InvalidCartItem
	InvalidParameter
	UnprocessableBody
	UnknownError
)

type ErrorResponse struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func NewServer(router Router) Server {
	return Server{router: &router}
}

func (s *Server) Start() {
	if err := http.ListenAndServe(cartapi.Config.Server.Addr, s.router); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func writeJsonResponse(body interface{}, statusCode int, w http.ResponseWriter, r *http.Request, action string) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.WithReqIdAndAction(log.Error().Stack().Err(err), r, action).
			Msg("failed to write response")
	}
}

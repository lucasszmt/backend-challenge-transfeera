package app

import (
	"github.com/lucasszmt/transfeera-challenge/app/v1/handler"
	"github.com/lucasszmt/transfeera-challenge/app/v1/routes"
)

func (s *Server) router() {
	routes.ReceiverRoutes(s.app, handler.NewReceiverHandler(s.receiverService))
}

package infrastructure

import "letargico/about/domain"

type handler struct {
	aboutService domain.AboutGateway
}

func NewAboutHandler(aboutService domain.AboutGateway) AboutHandler {
	return &handler{aboutService: aboutService}
}

package infrastructure

import "github.com/panxomon/letargico/about/domain"

type handler struct {
	aboutService domain.Service
}

func NewAboutHandler(aboutService domain.Service) AboutHandler {
	return &handler{aboutService: aboutService}
}

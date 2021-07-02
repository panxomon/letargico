package domain

type AboutGateway interface {
	Find() (*About, error)
}

func NewAboutService(aboutrepo AboutRepository) AboutGateway {
	return &service{aboutrepo: aboutrepo}
}

func (s *service) Find() (*About, error) {
	return s.aboutrepo.Find()
}

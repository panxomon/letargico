package about

func NewAboutService(aboutrepo Repository) AboutGateway {
	return &service{aboutrepo: aboutrepo}
}

type AboutGateway interface {
	Find() (*About, error)
}

func (s *service) Find() (*About, error) {
	return s.aboutrepo.Find()
}

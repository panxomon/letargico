package domain

type About struct {
	Message string `json:"message"`
}

//servicio que se expone
type Service interface {
	Find() (*About, error)
}

//repositorio de datos
type Repository interface {
	Find() (*About, error)
}

//struct interna
type service struct {
	aboutrepo Repository
}

func NewAboutService(aboutrepo Repository) Service {
	return &service{aboutrepo: aboutrepo}
}

func (s *service) Find() (*About, error) {
	return s.aboutrepo.Find()
}

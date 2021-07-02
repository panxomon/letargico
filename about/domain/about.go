package domain

type About struct {
	Message string `json:"message"`
}

type service struct {
	aboutrepo AboutRepository
}

type AboutRepository interface {
	Find() (*About, error)
}

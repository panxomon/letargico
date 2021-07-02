package about

type About struct {
	Message string `json:"message"`
}

type service struct {
	aboutrepo Repository
}

type Repository interface {
	Find() (*About, error)
}

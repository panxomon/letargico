package infraestructure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/panxomon/letargico/about/domain"
)

type AboutHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	aboutService domain.Service
}

func NewAboutHandler(aboutService domain.Service) AboutHandler {
	return &handler{aboutService: aboutService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p, err := h.aboutService.Find()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&p)
}

type fileRepository struct {
	file string
}

func NewDirRepository() (domain.Repository, error) {

	repo := &fileRepository{file: "files/about.md"}

	return repo, nil

}

func (f *fileRepository) Find() (*domain.About, error) {

	about := &domain.About{}

	content, err := ioutil.ReadFile(f.file)

	if err != nil {
		return nil, err
	}

	about.Message = string(content)

	return about, nil
}

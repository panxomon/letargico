package infrastructure

import (
	"io/ioutil"
	"letargico/about/domain"
)

type fileRepository struct {
	file string
}

func NewDirRepository() (domain.AboutRepository, error) {

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

package article

type Articles []Article

type Article struct {
	Titulo   string  `json:"titulo"`
	Articulo string  `json:"articulo"`
	Fecha    string  `json:"fecha"`
	Autor    []Autor `json:"autor"`
}

type Autor struct {
	Nombre string `json:"nombre"`
	Alias  string `json:"alias"`
}

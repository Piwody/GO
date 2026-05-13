package repository

import "integrador/internal/domain"

type MemoryRepo struct {
	db map[string]*domain.Image
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		db: make(map[string]*domain.Image),
	}
}

func (r *MemoryRepo) Guardar(img *domain.Image, id string) {
	r.db[id] = img
}

func (r *MemoryRepo) Listar() []*domain.Image {

	lista := make([]*domain.Image, 0)

	for _, img := range r.db {
		lista = append(lista, img)
	}
	return lista
}

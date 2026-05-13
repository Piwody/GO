package domain

import "time"

type Image struct {
	id        string
	nombre    string
	formato   string
	tags      []string
	createdAt time.Time
}

func NewImage(id, nombre, formato string, tags []string) *Image {
	return &Image{
		id:        id,
		nombre:    nombre,
		formato:   formato,
		tags:      tags,
		createdAt: time.Now(),
	}
}

//Set
func (i *Image) GetNombre() string  { return i.nombre }
func (i *Image) GetFormato() string { return i.formato }

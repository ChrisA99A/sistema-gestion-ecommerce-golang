package models

import "errors"

type Usuario struct {
	id     int
	nombre string
	correo string
	rol    string
}

func NewUsuario(id int, nombre string, correo string, rol string) (*Usuario, error) {
	if nombre == "" {
		return nil, errors.New("el nombre del usuario no puede estar vacío")
	}

	if correo == "" {
		return nil, errors.New("el correo del usuario no puede estar vacío")
	}

	if rol != "cliente" && rol != "administrador" {
		return nil, errors.New("el rol debe ser cliente o administrador")
	}

	return &Usuario{
		id:     id,
		nombre: nombre,
		correo: correo,
		rol:    rol,
	}, nil
}

func (u Usuario) ID() int {
	return u.id
}

func (u Usuario) Nombre() string {
	return u.nombre
}

func (u Usuario) Rol() string {
	return u.rol
}

package productos

import "context"

type Service interface { // punto 1.b
	GetAll() ([]Product, error)
	Store(nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	UpdateNameAndPrice(id int, nombre string, precio float64) (Product, error)
	Delete(id int) (Product, error)
}
type service struct { // punto 1.c
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, color, codigo, fecha_creacion, stock, precio, publicado)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	return s.repository.Update(id, nombre, color, codigo, fecha_creacion, stock, precio, publicado)
}

func (s *service) Delete(id int) (Product, error) {
	return s.repository.Delete(id)
}
func (s *service) UpdateNameAndPrice(id int, nombre string, precio float64) (Product, error) {
	return s.repository.UpdateNameAndPrice(id, nombre, precio)
}

func NewService(r Repository) Service { // punto 1.d
	return &service{
		repository: r,
	}
}

// Implementación de Storage en MySQL

type ServiceMySQL interface { // punto 1.b
	GetAll() ([]Product, error)
	GetByName(name string) (Product, error)
	GetById(id int) (Product, error)
	Store(nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	UpdateWithContext(ctx context.Context, id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	GetFullData() ([]Product_Warehouse, error)
}
type serviceMySQL struct { // punto 1.c
	repository RepositoryMySQL
}

func (s *serviceMySQL) GetAll() ([]Product, error) {
	productos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func (s *serviceMySQL) GetByName(name string) (Product, error) {
	producto, err := s.repository.GetByName(name)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}

func (s *serviceMySQL) GetById(id int) (Product, error) {
	producto, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}

func (s *serviceMySQL) Store(nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	producto, err := s.repository.Store(nombre, color, codigo, fecha_creacion, stock, precio, publicado)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}

func (s *serviceMySQL) Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	producto, err := s.repository.Update(id, nombre, color, codigo, fecha_creacion, stock, precio, publicado)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}

func (s *serviceMySQL) UpdateWithContext(ctx context.Context, id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	producto, err := s.repository.UpdateWithContext(ctx, id, nombre, color, codigo, fecha_creacion, stock, precio, publicado)
	if err != nil {
		return Product{}, err
	}
	return producto, nil
}

func (s *serviceMySQL) GetFullData() ([]Product_Warehouse, error) {
	productos, err := s.repository.GetFullData()
	if err != nil {
		return nil, err
	}
	return productos, nil
}

func NewServiceMySQL(r RepositoryMySQL) ServiceMySQL { // punto 1.d
	return &serviceMySQL{
		repository: r,
	}
}

package usuarios

// type Service interface {
// 	GetAll() ([]UsuarioRepository, error)
// 	Store(Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error)
// 	Update(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error)
// 	Delete(id int) (UsuarioRepository, error)
// 	UpdateLastNameAndAge(id int, apellido string, edad int) (UsuarioRepository, error)
// }
// type service struct {
// 	repository Repository
// }

// func (s *service) GetAll() ([]UsuarioRepository, error) {
// 	ps, err := s.repository.GetAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ps, nil
// }

// func (s *service) Store(Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error) {
// 	lastID, err := s.repository.LastID()
// 	if err != nil {
// 		return UsuarioRepository{}, err
// 	}

// 	lastID++

// 	usuario, err := s.repository.Store(lastID, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion)
// 	if err != nil {
// 		return UsuarioRepository{}, err
// 	}

// 	return usuario, nil
// }

// func (s *service) Update(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error) {
// 	return s.repository.Update(id, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion)
// }

// func (s *service) Delete(id int) (UsuarioRepository, error) {
// 	return s.repository.Delete(id)
// }
// func (s *service) UpdateLastNameAndAge(id int, apellido string, edad int) (UsuarioRepository, error) {
// 	return s.repository.UpdateLastNameAndAge(id, apellido, edad)
// }

// func NewService(r Repository) Service {
// 	return &service{
// 		repository: r,
// 	}
// }

type Service interface {
	GetAll() ([]Usuario, error)
	Store(Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error)
	Update(Id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error)
	Delete(id int) (Usuario, error)
	UpdateNameEdad(id int, name string, edad int) (Usuario, error)
}
type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Usuario, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion)
	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}
func (s *service) Update(Id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error) {

	return s.repository.Update(Id, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion)
}

func (s *service) UpdateNameEdad(Id int, Nombre string, Edad int) (Usuario, error) {
	return s.repository.UpdateNameEdad(Id, Nombre, Edad)
}

func (s *service) Delete(id int) (Usuario, error) {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

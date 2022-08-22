package transactions

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error)
	Update(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error)
	Delete(id int) (Transaccion, error)
	UpdateCodeAndAmount(id int, codigo_transaccion string, monto float64) (Transaccion, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Transaccion, error) {
	ts, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (s *service) Store(codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Transaccion{}, err
	}

	lastID++

	transaction, err := s.repository.Store(lastID, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion)

	if err != nil {
		return Transaccion{}, err
	}

	return transaction, nil
}

func (s *service) Update(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error) {
	return s.repository.Update(id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion)
}

func (s *service) Delete(id int) (Transaccion, error) {
	return s.repository.Delete(id)
}

func (s *service) UpdateCodeAndAmount(id int, codigo_transaccion string, monto float64) (Transaccion, error) {
	return s.repository.UpdateCodeAndAmount(id, codigo_transaccion, monto)

}

//constructor
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

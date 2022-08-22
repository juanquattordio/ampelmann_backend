package transactions

import "fmt"

type Transaccion struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

var ts []Transaccion
var lastId int

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error)
	LastID() (int, error)
	Update(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error)
	Delete(id int) (Transaccion, error)
	UpdateCodeAndAmount(id int, codigo_transaccion string, monto float64) (Transaccion, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Transaccion, error) {
	return ts, nil
}

func (r *repository) LastID() (int, error) {
	return lastId, nil
}

func (r *repository) Store(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error) {
	t := Transaccion{id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion}
	ts = append(ts, t)
	lastId = t.Id
	return t, nil
}

func (r *repository) Update(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error) {
	t := Transaccion{id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion}
	if len(ts) > 0 {
		updated := false
		for i := range ts {
			if ts[i].Id == id {
				t.Id = id
				ts[i] = t
				updated = true
			}
		}
		if !updated {
			return Transaccion{}, fmt.Errorf("Transacciono %d no encontrado", id)
		}
	} else {
		return Transaccion{}, fmt.Errorf("No hay elementos para actualizar")
	}
	return t, nil
}

func (r *repository) UpdateCodeAndAmount(id int, codigo_transaccion string, monto float64) (Transaccion, error) {
	var t Transaccion
	if len(ts) > 0 {
		updated := false
		for i := range ts {
			if ts[i].Id == id {
				t = ts[i]
				t.Id = id
				t.CodigoTransaccion = codigo_transaccion
				t.Monto = monto
				ts[i] = t
				updated = true
			}
		}
		if !updated {
			return Transaccion{}, fmt.Errorf("Transaccion %d - codigo %s no encontrado", id, codigo_transaccion)
		}
	} else {
		return Transaccion{}, fmt.Errorf("No hay elementos para actualizar")
	}
	return t, nil
}

func (r *repository) Delete(id int) (Transaccion, error) {
	t := Transaccion{}
	var index int
	if len(ts) > 0 {
		deleted := false
		for i := range ts {
			if ts[i].Id == id {
				t = ts[i]
				index = i
				deleted = true
			}
		}
		if !deleted {
			return Transaccion{}, fmt.Errorf("Transaccion %d no encontrada", id)
		}
	} else {
		return Transaccion{}, fmt.Errorf("No hay elementos para eliminar")
	}
	ts = append(ts[:index], ts[index+1:]...) // aca pisa el slice de Transaccionos extrayendo el eliminado
	return t, nil
}

func NewRepository() Repository {
	return &repository{}
}

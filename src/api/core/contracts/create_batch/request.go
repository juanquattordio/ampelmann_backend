package create_batch

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomFechaOrigen time.Time

type Request struct {
	IdReceta         *int64             `form:"id_receta" json:"id_receta" binding:"required"`
	FechaOrigen      *CustomFechaOrigen `form:"fecha" json:"fecha"`
	LitrosProducidos *float64           `form:"litros_producidos" json:"litros_producidos" binding:"required"`
}

// Implement Marshaler and Unmarshaler interface
func (j *CustomFechaOrigen) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = CustomFechaOrigen(t)
	return nil
}

func (j CustomFechaOrigen) MarshalJSON() ([]byte, error) {
	miJson, err := json.Marshal(time.Time(j))
	return miJson, err
}

// Maybe a Format function for printing your date
func (j CustomFechaOrigen) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

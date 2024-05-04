package datatypes

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (c NullString) MarshalJSON() ([]byte, error) {
	if c.String == "" && !c.Valid {
		return json.Marshal("")
	}
	return json.Marshal(c.String)
}

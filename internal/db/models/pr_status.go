package models

import (
	"database/sql/driver"
	"fmt"
)

type PRStatus string

const (
	PRStatusOpen   PRStatus = "OPEN"
	PRStatusMerged PRStatus = "MERGED"
)

func (PRStatus) Values() []string {
	return []string{
		string(PRStatusOpen),
		string(PRStatusMerged),
	}
}

func (s PRStatus) GormDataType() string {
	return "varchar(20)"
}

func (s PRStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s PRStatus) Scan(value interface{}) error {
	if value == nil {
		s = PRStatusOpen
		return nil
	}

	switch v := value.(type) {
	case []byte:
		s = PRStatus(string(v))
	case string:
		s = PRStatus(v)
	default:
		return fmt.Errorf("cannot scan %T into PRStatus", value)
	}

	// Валидация значения
	if !s.IsValid() {
		return fmt.Errorf("invalid PRStatus value: %s", s)
	}

	return nil
}

func (s PRStatus) IsValid() bool {
	switch s {
	case PRStatusOpen, PRStatusMerged:
		return true
	default:
		return false
	}
}

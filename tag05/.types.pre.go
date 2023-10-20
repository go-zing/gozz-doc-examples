package tag05

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type (
	UserStruct struct {
		Id        string
		Name      string
		Address   string
		CreatedAt time.Time
		UpdatedAt time.Time
		Friends   []struct {
			Id        string
			Name      string
			Address   string
			CreatedAt time.Time
			UpdatedAt time.Time
		}
	}

	UserMap map[string]struct {
		Id        string
		Name      string
		Address   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UserSlice []struct {
		Id        string
		Name      string
		Address   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UserInterface interface {
		User() struct {
			Id        string
			Name      string
			Address   string
			CreatedAt time.Time
			UpdatedAt time.Time
		}
	}

	UserFunc func(struct {
		Id        string
		Name      string
		Address   string
		CreatedAt time.Time
		UpdatedAt time.Time
	})
)

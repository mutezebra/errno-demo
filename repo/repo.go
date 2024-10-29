package repo

import "github.com/pkg/errors"

func Create(id int64) error {
	if id == 1 {
		return errors.Errorf("the id have exist,id: %d", id)
	}
	return nil
}

package service

import (
	"github.com/pkg/errors"

	"github.com/mutezebra/error-demo/pkg/errno"
	"github.com/mutezebra/error-demo/pkg/pack"
	"github.com/mutezebra/error-demo/pkg/snowflake"
	"github.com/mutezebra/error-demo/repo"
)

func Process(req string) (data interface{}, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if len(req) == 0 {
		return nil, errno.New(errno.ReqTooShort, "req is too short")
	}

	var id int64
	id, err = snowflake.GenerateID()
	if err != nil {
		return nil, errors.WithMessage(err, "get new id failed")
	}

	if err = repo.Create(id); err != nil {
		return nil, errors.WithMessage(err, "create by id failed")
	}

	data = errno.Success
	return data, nil
}

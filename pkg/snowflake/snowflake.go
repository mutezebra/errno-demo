package snowflake

import "github.com/pkg/errors"

var count int64

func GenerateID() (int64, error) {
	if count == 1 {
		return 0, errors.New("generate id failed")
	}
	count++
	return count, nil
}

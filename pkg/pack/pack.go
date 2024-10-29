package pack

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mutezebra/error-demo/pkg/errno"
)

func LogError(err error) {
	if err == nil {
		return
	}

	var e errno.Errno
	if errors.As(err, &e) {
		log.Printf("Code: %d, Message: %s", e.Code(), e.Error())
		return
	}
	log.Printf("Error: %+v", err)
}

func SendError(w http.ResponseWriter, r *http.Request, err error) {
	var e errno.Errno

	if !errors.As(err, &e) {
		e = errno.ServerError
	}

	// 如果不是errno的错误,那就说明是我们自己的问题
	resp := map[string]interface{}{
		"code": e.Code(),
		"msg":  e.Error(),
	}
	// 响应
	_ = resp
}

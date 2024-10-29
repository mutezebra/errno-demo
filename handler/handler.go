package handler

import (
	"net/http"

	"github.com/mutezebra/error-demo/pkg/pack"
	"github.com/mutezebra/error-demo/service"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// ...something
	req := ""

	resp, err := service.Process(req)
	if err != nil {
		pack.SendError(w, r, err)
		return
	}

	// 处理正常的响应
	_ = resp
}

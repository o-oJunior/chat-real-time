package router

import (
	"net/http"
)

var userRouters = []Router{
	{
		URI:            "/createUser",
		Method:         http.MethodPost,
		Function:       func(w http.ResponseWriter, r *http.Request) {},
		Authentication: false,
	},
}

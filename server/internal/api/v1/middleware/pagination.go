package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePagination(ctx *gin.Context) (int, int, int) {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
		logger.Warn("(Page) Página inválida, será setado novo valor: %d", page)
	}
	limitStr := ctx.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
		logger.Warn("(Limit) Limite inválido, será setado novo valor: %d", limit)
	}
	offset := (page - 1) * limit
	return page, limit, offset
}

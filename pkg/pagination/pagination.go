package gin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yrcs/yuneto/third_party/pagination"
	"gorm.io/gorm"
)

func PackPagingData(ctx *gin.Context, req *pagination.PagingRequest) {
	pInt, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	psInt, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "1000"))
	page := uint32(pInt)
	pageSize := uint32(psInt)
	query := ctx.QueryMap("query")
	req.Page = &page
	req.PageSize = &pageSize
	req.Query = query
	orderBy := ctx.QueryMap("orderBy")
	for k, v := range orderBy {
		ov, _ := strconv.Atoi(v)
		switch ov {
		case 1:
			req.OrderBy[k] = pagination.Order_ASC
		case 2:
			req.OrderBy[k] = pagination.Order_DESC
		}
	}
}

func Paginate(req *pagination.PagingRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("DeletedAt IS NULL").Where(req.GetQuery())
		for k, v := range req.OrderBy {
			db = db.Order(k + " " + pagination.Order_name[int32(v.Number())])
		}
		offset := (req.GetPage() - 1) * req.GetPageSize()
		return db.Limit(int(req.GetPageSize())).Offset(int(offset))
	}
}

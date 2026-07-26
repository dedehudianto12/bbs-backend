package http

import (
	"net/http"
	"strconv"
)

type PaginatedResponse struct {
	Items any    `json:"items"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

func ParsePagination(r *http.Request) (page, limit int, search, sort string) {
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}
	search = r.URL.Query().Get("search")
	sort = r.URL.Query().Get("sort")
	if sort != "asc" {
		sort = "desc"
	}
	return
}

func SuccessPaginated(w http.ResponseWriter, status int, items any, total, page, limit int, sort string) {
	resp := Response{
		Data: PaginatedResponse{
			Items: items,
			Total: total,
			Page:  page,
			Limit: limit,
			Sort:  sort,
		},
		Error: nil,
	}
	writeJSON(w, status, resp)
}

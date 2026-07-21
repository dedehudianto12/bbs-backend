package category

type Category struct {
	Slug      string `json:"slug"`
	Label     string `json:"label"`
	Group     string `json:"group"`
	SortOrder int    `json:"sortOrder"`
}

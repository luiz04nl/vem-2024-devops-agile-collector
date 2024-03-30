package shared

type Repository struct {
	Name            string `db:"name"`
	URL             string `db:"url"`
	StarsTotalCount int    `db:"starsTotalCount"`
}

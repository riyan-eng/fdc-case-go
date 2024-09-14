package datastruct

type ExampleList struct {
	UUID      string `db:"uuid" json:"id"`
	Name      any    `db:"nama" json:"name"`
	Detail    any    `db:"detail" json:"detail"`
	TotalRows int    `db:"total_rows" json:"-"`
}

type ExampleDetail struct {
	UUID   string `db:"uuid" json:"id"`
	Name   any    `db:"nama" json:"name"`
	Detail any    `db:"detail" json:"detail"`
}

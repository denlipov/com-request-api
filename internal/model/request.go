package model

// Request - request entity.
type Request struct {
	ID  uint64 `db:"id"`
	Foo uint64 `db:"foo"`
}

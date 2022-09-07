package do

import "time"

type Dict struct {
	Id          uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ParentId    uint64
	Name        string
	Value       string
	DictCode    string
	HasChildren bool
}

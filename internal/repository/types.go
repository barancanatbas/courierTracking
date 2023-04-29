package repository

type FindType int

const (
	FindAll FindType = iota
	FindOne
)

type AdvanceQuery struct {
	Wheres   []Where
	Joins    []string
	Preloads []string
	Select   string
	FindType FindType
}

type Where struct {
	Column   string
	Operator Operator
	Value    interface{}
}

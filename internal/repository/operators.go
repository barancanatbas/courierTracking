package repository

type Operator string

const (
	EqOperator      Operator = "="
	NotEqOperator   Operator = "<>"
	GtOperator      Operator = ">"
	LtOperator      Operator = "<"
	GteOperator     Operator = ">="
	LteOperator     Operator = "<="
	LikeOperator    Operator = "like"
	NotLikeOperator Operator = "not like"
	InOperator      Operator = "in"
	NotInOperator   Operator = "not in"
)

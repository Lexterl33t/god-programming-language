package ast

type Program struct {
	Body BlockState
}

type BlockState struct {
	List []interface{}
}

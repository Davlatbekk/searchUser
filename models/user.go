package models

type User struct {
	Id       int
	Name     string
	Surname  string
	Birthday string
}

type GetListRequest struct {
	Offset, Limit    int
	Search           string
	FromDate, ToDate string
}

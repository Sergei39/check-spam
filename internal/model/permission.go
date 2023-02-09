package model

type Permission string

const (
	ALLOWED     = Permission("allowed")
	NOT_ALLOWED = Permission("not_allowed")
)

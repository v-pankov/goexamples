package request

// Context is a part of usecase request model consideted valid
type Context struct {
}

type Model struct {
	UserName UserName
}

type UserName string

func (un UserName) String() string {
	return string(un)
}

package request

type Context struct {
}

type Model struct {
	UserName UserName
}

type UserName string

func (un UserName) String() string {
	return string(un)
}

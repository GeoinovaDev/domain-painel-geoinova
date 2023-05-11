package app

type IQueryOption interface {
	Apply(interface{})
}

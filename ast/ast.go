package ast

type File struct {
	Module Module
	Go     Go
}

type Module struct {
	Name string
}

type Go struct {
	Version string
}

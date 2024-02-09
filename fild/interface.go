package fild

type IFild interface {
	Init(different int) (bool, Fild)
	Copy() Fild
	Soluve() (bool, Fild)
}

package Interface

type AbstractFactory interface {
	CreateChair() IChair
	CreateTable() ITable
	CreateSofa() ISofa
}
package concreates

import (
	"rgunindi/abstract-factory-go/Products/Modern"
)

type ModernFactory struct{}

func (m *ModernFactory) CreateChair() modern.Chair {
	return modern.Chair{}
}

func (m *ModernFactory) CreateSofa() modern.Sofa {
	return modern.Sofa{}
}

func (m *ModernFactory) CreateTable() modern.Table {
	return modern.Table{}
}

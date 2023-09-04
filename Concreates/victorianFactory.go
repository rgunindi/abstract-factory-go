package concreates

import (
	"rgunindi/abstract-factory-go/Products/Victorian"
)

type VictorianFactory struct{}

func (v *VictorianFactory) CreateChair() victorian.Chair {
	return victorian.Chair{}
}

func (v *VictorianFactory) CreateSofa() victorian.Sofa {
	return victorian.Sofa{}
}

func (v *VictorianFactory) CreateTable() victorian.Table {
	return victorian.Table{}
}

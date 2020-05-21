package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	txtOp := new(ExportTextFileOperate)
	api := txtOp.FactoryMethod()
	api.Export() // export to xxx.txt

	dbOp := new(ExportDBOperate)
	api = dbOp.FactoryMethod()
	api.Export() // export to db
}

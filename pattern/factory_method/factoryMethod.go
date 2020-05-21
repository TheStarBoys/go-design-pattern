package factory_method

import "fmt"

// ExportFileApi 工厂方法所创建的对象的接口
type ExportFileApi interface {
	Export()
}

// ExportTextFile 导出到文件，实现了ExportFileApi
type ExportTextFile struct {}

func (e *ExportTextFile) Export() {
	fmt.Println("export to xxx.txt")
}

type ExportDB struct {}

func (e *ExportDB) Export() {
	fmt.Println("export to db")
}

// ExportFileOperate 创建器接口，声明工厂方法
type ExportFileOperate interface {
	FactoryMethod() ExportFileApi
}

// ExportTextFileOperate 具体的创建器实现对象
type ExportTextFileOperate struct {}

func (e *ExportTextFileOperate) FactoryMethod() ExportFileApi {
	return new(ExportTextFile)
}

// ExportDBOperate 具体的创建器实现对象
type ExportDBOperate struct {}

func (e *ExportDBOperate) FactoryMethod() ExportFileApi {
	return new(ExportDB)
}
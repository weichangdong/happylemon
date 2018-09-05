package models

type TestModel struct {
	*BaseModels
}

func (self *TestModel) ReadInfo() map[string]interface{} {
	res, _ := MyDb.Table("user").First()
	return res
}

package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

// ClearCounter 清除Counter
func (imp *CounterInterfaceImp) ClearCounter(id int32) error {
	cli := db.Get()
	return cli.Table(imp.TableName()).Delete(&model.CounterModel{Id: id}).Error
}

// UpsertCounter 更新/写入counter
func (imp *CounterInterfaceImp) UpsertCounter(counter *model.CounterModel) error {
	cli := db.Get()
	return cli.Table(imp.TableName()).Save(counter).Error
}

// GetCounter 查询Counter
func (imp *CounterInterfaceImp) GetCounter(id int32) (*model.CounterModel, error) {
	var err error
	var counter = new(model.CounterModel)

	cli := db.Get()
	err = cli.Table(imp.TableName()).Where("id = ?", id).First(counter).Error

	return counter, err
}

func (imp *CounterInterfaceImp) TableName() string { return "Counters"}

// AddContactor 更新/写入Contactor
func (imp *ContactorInterfaceImp) AddContactor(Contactor *model.ContactorModel) error {
	cli := db.Get()
	return cli.Table(imp.TableName()).Save(Contactor).Error
}

// GetContactors 查询Contactors
func (imp *ContactorInterfaceImp) GetContactors() (*[]model.ContactorModel, error) {
	var err error
	var Contactor = []model.ContactorModel{}

	cli := db.Get()
	err = cli.Table(imp.TableName()).Find(&Contactor).Error

	return &Contactor, err
}

func (imp *ContactorInterfaceImp) TableName() string { return "contactors"}
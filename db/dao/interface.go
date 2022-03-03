package dao

import (
	"wxcloudrun-golang/db/model"
)

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetCounter(id int32) (*model.CounterModel, error)
	UpsertCounter(counter *model.CounterModel) error
	ClearCounter(id int32) error
	TableName() string
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}


// ContactorInterface 联系人数据模型接口
type ContactorInterface interface {
	GetContactors() (*[]model.ContactorModel, error)
	AddContactor(Contactor *model.ContactorModel) error
	TableName() string
}

// ContactorInterfaceImp 联系人数据模型实现
type ContactorInterfaceImp struct{}

// Counter 实现实例
var Counter CounterInterface = &CounterInterfaceImp{}
var Contactor ContactorInterface = &ContactorInterfaceImp{}
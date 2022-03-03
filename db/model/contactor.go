package model

import "time"

// ContactorModel 通讯录模型
type ContactorModel struct {
	Id        int32     `gorm:"column:Id" json:"Id"`
	Name      string    `gorm:"column:Name" json:"Name"`
	Gender    string    `gorm:"column:Gender" json:"Gender"`
	Nation    string    `gorm:"column:Nation" json:"Nation"`
	StartYear string    `gorm:"column:StartYear" json:"StartYear"`
	College   string    `gorm:"column:College" json:"College"`
	Phone     string    `gorm:"column:Phone" json:"Phone"`
	QQ        string    `gorm:"column:QQ" json:"QQ"`
	Wechat    string    `gorm:"column:Wechat" json:"Wechat"`
	Country   string    `gorm:"column:Country" json:"Country"`
	Prov      string    `gorm:"column:Prov" json:"Prov"`
	City      string    `gorm:"column:City" json:"City"`
	Address   string    `gorm:"column:Address" json:"Address"`
	Company   string    `gorm:"column:Company" json:"Company"`
	Title     string    `gorm:"column:Title" json:"Title"`
	Author    string    `gorm:"column:Author" json:"Author"`
	CreatedAt time.Time `gorm:"column:CreatedAt" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt" json:"UpdatedAt"`
}

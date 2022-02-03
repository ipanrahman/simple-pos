package model

type CasbinRule struct {
	ID    uint   `gorm:"column:id;primaryKey;autoIncrement"`
	PType string `gorm:"column:ptype;size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"column:v0;size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"column:v1;size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"column:v2;size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"column:v3;size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"column:v4;size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"column:v5;size:512;uniqueIndex:unique_index"`
}

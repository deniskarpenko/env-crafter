package models

type Tag struct {
	TagId   int    `gorm:"column:tag_id:primary_key;autoIncrement;" json:"tag_id"`
	ImageId int    `gorm:"column:image_id;not null;" json:"image_id"`
	TagName string `gorm:"column:tag_name;size:25"`
}

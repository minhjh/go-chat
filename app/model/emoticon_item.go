package model

import "time"

type EmoticonItem struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`  // 表情包详情ID
	EmoticonId int       `gorm:"column:emoticon_id;default:0" json:"emoticon_id"` // 表情分组ID
	UserId     int       `gorm:"column:user_id;default:0" json:"user_id"`         // 用户ID（0：代码系统表情包）
	Describe   string    `gorm:"column:describe" json:"describe"`                 // 表情描述
	Url        string    `gorm:"column:url" json:"url"`                           // 图片链接
	FileSuffix string    `gorm:"column:file_suffix" json:"file_suffix"`           // 文件后缀名
	FileSize   int       `gorm:"column:file_size;default:0" json:"file_size"`     // 文件大小（单位字节）
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`             // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`             // 更新时间
}

package models

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 公共模型
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`    // id主键
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // 逻辑删除标记
	CreatedAt time.Time      `json:"created_at"`              // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`              // 修改时间
}

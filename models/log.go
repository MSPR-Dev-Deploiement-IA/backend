package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Path     string `json:"path"`
	Method   string `json:"method"`
	Status   int    `json:"status"`
	Latency  int64  `json:"latency"`
	IPID     uint   `json:"ip_id" gorm:"type:int;not null"` // Foreign key
	ClientIP IP     `json:"client_ip" gorm:"foreignKey:IPID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

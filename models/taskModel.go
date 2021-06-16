package models

import (
	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	USER_ID                uint
	TASK_NAME              string `gorm:"type:varchar(255); not null" json:"taskName"`
	ESTIMATED_POMS         string `gorm:"type:varchar(255); not null" json:"estimatedPoms"`
	POM_SECONDS            string `gorm:"type:int; default: 0" json:"pomSeconds"`
	BREAK_SECONDS          string `gorm:"type:int; default: 0" json:"breakSeconds"`
	COMPLETED_POMS         string `gorm:"type:int; default: 0" json:"completedPoms"`
	COMPLETED_SMALL_BREAKS string `gorm:"type:int; default: 0" json:"completedSmallBreaks"`
	COMPLETED_BIG_BREAKS   string `gorm:"type:int; default: 0" json:"completedBigBreaks"`
	IS_COMPLETE            string `gorm:"type:boolean; default: false" json:"isComplete"`
	START_DATE             string `gorm:"type:date;" json:"startDate"`
	COMPLETE_DATE          string `gorm:"type:date;" json:"completeDate"`
}

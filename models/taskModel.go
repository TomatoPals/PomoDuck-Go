package models

import (
	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	USER_ID                uint
	TASK_NAME              string `json:"taskName"`
	ESTIMATED_POMS         string `json:"estimatedPoms"`
	POM_SECONDS            string `json:"pomSeconds"`
	BREAK_SECONDS          string `json:"breakSeconds"`
	COMPLETED_POMS         string `json:"completedPoms"`
	COMPLETED_SMALL_BREAKS string `json:"completedSmallBreaks"`
	COMPLETED_BIG_BREAKS   string `json:"completedBigBreaks"`
	IS_COMPLETE            string `json:"isComplete"`
	START_DATE             string `json:"startDate"`
	COMPLETE_DATE          string `json:"completeDate"`
}

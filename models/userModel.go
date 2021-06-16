package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FIRST_NM                  string `gorm:"type:varchar(255); not null" json:"firstName"`
	LAST_NM                   string `gorm:"type:varchar(255); not null" json:"lastName"`
	EMAIL                     string `gorm:"type:varchar(255); not null" json:"email"`
	PASSWORD                  string `gorm:"type:varchar(255); not null" json:"password"`
	COUNTRY                   string `json:"country"`
	ALIAS                     string `json:"alias"`
	ALIAS_IMG                 string `json:"aliasImg"`
	DISPLAY_PREF              string `gorm:"type:int; default: 1" json:"displayPref"`
	TOTAL_POM_SECONDS         string `gorm:"type:int; default: 0" json:"totalPomSeconds"`
	TOTAL_SMALL_BREAK_SECONDS string `gorm:"type:int; default: 0" json:"totalSmallBreakSeconds"`
	TOTAL_BIG_BREAK_SECONDS   string `gorm:"type:int; default: 0" json:"totalBigBreakSeconds"`
	POM_TIME                  string `gorm:"type:int; default: 25" json:"pomTime"`
	SMALL_BREAK_TIME          string `gorm:"type:int; default: 5" json:"smallBreakTime"`
	BIG_BREAK_TIME            string `gorm:"type:int; default: 15" json:"bigBreakTime"`
	Tasks                     []Tasks
}

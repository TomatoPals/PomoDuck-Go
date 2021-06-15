package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FIRST_NM                  string `json:"firstName"`
	LAST_NM                   string `json:"lastName"`
	EMAIL                     string `json:"email"`
	PASSWORD                  string `json:"password"`
	COUNTRY                   string `json:"country"`
	ALIAS                     string `json:"alias"`
	ALIAS_IMG                 string `json:"aliasImg"`
	DISPLAY_PREF              string `json:"displayPref"`
	TOTAL_POM_SECONDS         string `json:"totalPomSeconds"`
	TOTAL_SMALL_BREAK_SECONDS string `json:"totalSmallBreakSeconds"`
	TOTAL_BIG_BREAK_SECONDS   string `json:"totalBigBreakSeconds"`
	POM_TIME                  string `json:"pomTime"`
	SMALL_BREAK_TIME          string `json:"smallBreakTime"`
	BIG_BREAK_TIME            string `json:"bigBreakTime"`
	Tasks                     []Tasks
}

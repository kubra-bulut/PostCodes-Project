package models

import (
	"fmt"
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type PostCode struct {
	ID          uint           `json:"id" gorm:"primaryKey;comment:ID"`
	Code        string         `json:"code" gorm:"index:post_code_idx,sort:desc;comment:Posta Kodu"`
	City        string         `json:"city" gorm:"index:city_country_town_idx,sort:desc;comment:Şehir"`
	County      string         `json:"county" gorm:"index:city_county_town_idx,sort:desc;comment:İlçe"`
	Town        string         `json:"town" gorm:"index:city_county_town_idx,sort:desc;comment:Semt, Bucak, Belde"`
	District    string         `json:"district" gorm:"comment:Mahalle"`
	HiFromHook  string         `json:"hi_from_hook" gorm:"-"`
	CountryCode string         `json:"country_code" gorm:"default:Turkiye;comment:Ülke Kodu"`
	Location    datatypes.JSON `json:"location,omitempty"`
	Tags        pq.StringArray `gorm:"type:text[]"`
	UploadDate  time.Time      `json:"upload_date" gorm:"type:date;comment:Yükleme Tarihi"`
	CreatedAt   time.Time      `json:"created_at" gorm:"comment:Oluşturulma Tarihi"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"comment:Güncellenme Tarihi"`
}

type Location struct {
	X, Y float64
}

func (p PostCode) TableName() string {
	return "post_codes"
}
func (p *PostCode) BeforeCreate(tx *gorm.DB) (err error) {
	p.UploadDate = time.Now()
	return
}

func (p *PostCode) AfterFind(tx *gorm.DB) (err error) {
	p.HiFromHook = fmt.Sprintf("Hi from hook %s", p.Code)
	return
}

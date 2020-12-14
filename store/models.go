package store

import (
	"time"
)

// GeoCoordinate db struct
type GeoCoordinate struct {
	ID        uint `gorm:"autoIncrement:true;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Lat       float64 `gorm:"primaryKey" json:"lat"`
	Lon       float64 `gorm:"primaryKey" json:"lon"`
	Postcode  string  `json:"postcode"`
}

// CSVUpload struct to handle progress
type CSVUpload struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Records   int `json:"-"`
	Bulks     int
	Counts    int
	Status    string
	Reference string
}

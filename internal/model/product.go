package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Product struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SKU          string         `gorm:"unique;not null" json:"sku"`
	Name         string         `gorm:"not null" json:"name"`
	Slug         string         `gorm:"unique;not null" json:"slug"`
	Category     string         `json:"category"`
	Description  string         `json:"description"`
	Price        float64        `gorm:"type:numeric(12,2);not null" json:"price"`
	SalePrice    *float64       `gorm:"type:numeric(12,2)" json:"sale_price"`
	Currency     string         `gorm:"default:'VND'" json:"currency"`
	Stock        int            `gorm:"default:0" json:"stock"`
	Sold         int            `gorm:"default:0" json:"sold"`
	Thumbnail    string         `json:"thumbnail"`
	Images       datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"images"`
	Sizes        datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"sizes"`
	Colors       datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"colors"`
	Tags         datatypes.JSON `gorm:"type:jsonb;default:'[]'" json:"tags"`
	Rating       float32        `gorm:"default:0" json:"rating"`
	ReviewCount  int            `gorm:"default:0" json:"review_count"`
	IsBestSeller bool           `gorm:"default:false" json:"is_best_seller"`
	IsNew        bool           `gorm:"default:true" json:"is_new"`
	Status       string         `gorm:"default:'active'" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

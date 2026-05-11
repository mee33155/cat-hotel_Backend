package database

import (
	"log"

	"backend-go/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("cat_hotel.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&models.Room{}, &models.Booking{})

	var count int64
	DB.Model(&models.Room{}).Count(&count)
	if count == 0 {
		rooms := []models.Room{
			{
				ID:          1,
				Name:        "Standard Room",
				Description: "ห้องขนาดเล็ก เหมาะกับแมวตัวเล็ก พร้อมที่นอนนุ่มและแม่เหล็กสำหรับเล่น บรรยากาศอบอุ่นเหมือนอยู่บ้าน",
				Price:       500,
				Capacity:    1,
				Size:        "12 ตร.ม.",
				Amenities:   `["ที่นอนนุ่ม","แม่เหล็กเล่น","ถาดทราย","ชามอาหารและน้ำ","ตู้กระจกดูแมว"]`,
			},
			{
				ID:          2,
				Name:        "VIP Suite",
				Description: "ห้องกว้างขวาง มีระเบียงส่องแดดให้น้องแมวนอนอาบแดด พร้อมเครื่องกรองอากาศและกล้องวงจรปิดดูได้ตลอด 24 ชม.",
				Price:       1200,
				Capacity:    2,
				Size:        "24 ตร.ม.",
				Amenities:   `["ระเบียงส่องแดด","เครื่องกรองอากาศ","กล้องวงจรปิด 24 ชม.","ที่นอนหรู","ถาดทรายอัตโนมัติ","ชามอาหารอัตโนมัติ","ของเล่นหลากชนิด"]`,
			},
		}
		DB.Create(&rooms)
	}
}

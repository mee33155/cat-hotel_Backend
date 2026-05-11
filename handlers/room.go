package handlers

import (
    "backend-go/database"
    "backend-go/models"

    "github.com/gofiber/fiber/v2"
)

func GetRooms(c *fiber.Ctx) error {
    var rooms []models.Room
    database.DB.Find(&rooms)
    return c.JSON(rooms)
}

func GetRoomByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var room models.Room
    if err := database.DB.First(&room, id).Error; err != nil {
        return c.Status(404).SendString("Room not found")
    }
    return c.JSON(room)
}

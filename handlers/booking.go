package handlers

import (
    "time"

    "backend-go/database"
    "backend-go/models"

    "github.com/gofiber/fiber/v2"
    "github.com/stripe/stripe-go/v74"
    "github.com/stripe/stripe-go/v74/checkout/session"
)

func GetBookings(c *fiber.Ctx) error {
    var bookings []models.Booking
    database.DB.Preload("Room").Find(&bookings)
    return c.JSON(bookings)
}

func CreateCheckoutSession(c *fiber.Ctx) error {
    type Request struct {
        CustomerName string `json:"customer_name"`
        CatName      string `json:"cat_name"`
        RoomID       uint   `json:"room_id"`
        CheckIn      string `json:"check_in"`
        CheckOut     string `json:"check_out"`
    }

    var req Request
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).SendString("Invalid input")
    }

    checkIn, err1 := time.Parse("2006-01-02", req.CheckIn)
    checkOut, err2 := time.Parse("2006-01-02", req.CheckOut)
    if err1 != nil || err2 != nil {
        return c.Status(400).SendString("Invalid date format")
    }

    nights := int(checkOut.Sub(checkIn).Hours() / 24)
    if nights < 1 {
        nights = 1
    }

    var room models.Room
    database.DB.First(&room, req.RoomID)

    params := &stripe.CheckoutSessionParams{
        PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
        LineItems: []*stripe.CheckoutSessionLineItemParams{
            {
                PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
                    Currency: stripe.String("thb"),
                    ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
                        Name: stripe.String(room.Name + " - Booking for " + req.CatName),
                    },
                    UnitAmount: stripe.Int64(int64(room.Price * float64(nights) * 100)),
                },
                Quantity: stripe.Int64(1),
            },
        },
        Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
        SuccessURL: stripe.String("http://localhost:3000/success?session_id={CHECKOUT_SESSION_ID}"),
        CancelURL:  stripe.String("http://localhost:3000/cancel"),
    }

    s, err := session.New(params)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    booking := models.Booking{
        CustomerName: req.CustomerName,
        CatName:      req.CatName,
        RoomID:       req.RoomID,
        CheckIn:      checkIn,
        CheckOut:     checkOut,
        SessionID:    s.ID,
        Status:       "pending",
    }
    database.DB.Create(&booking)

    return c.JSON(fiber.Map{"url": s.URL})
}

func SuccessPayment(c *fiber.Ctx) error {
    sessionID := c.Query("session_id")
    database.DB.Model(&models.Booking{}).Where("session_id = ?", sessionID).Update("status", "paid")
    return c.SendString("Payment Successful!")
}

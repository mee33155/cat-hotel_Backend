package models

type Room struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    ImageURL    string  `json:"image_url"`
    Capacity    int     `json:"capacity"`
    Size        string  `json:"size"`
    Amenities   string  `json:"amenities"` // JSON array stored as string
}

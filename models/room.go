package models

import "encoding/json"

type StringArray string

func (s StringArray) MarshalJSON() ([]byte, error) {
	var arr []string
	if err := json.Unmarshal([]byte(s), &arr); err != nil {
		return []byte("[]"), nil
	}
	return json.Marshal(arr)
}

type Room struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	ImageURLs   StringArray `json:"image_urls"`
	Capacity    int         `json:"capacity"`
	Size        string      `json:"size"`
	Amenities   StringArray `json:"amenities"`
}

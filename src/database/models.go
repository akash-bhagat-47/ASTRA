package db

type DataLine struct {
	City      string  `json:"city"`
	Temp      float32 `json:"temp"`
	Humidity  float32 `json:"humidity"`
	Timestamp string  `json:"timestamp"`
}

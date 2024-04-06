package db

type DataLine struct {
	City      string  `json:"city"`
	Temp      float32 `json:"temp"`
	Humidity  float32 `json:"humidity"`
	Timestamp string  `json:"timestamp"`
}

// curl -X POST http://localhost:8080/albums/add \
// -H "Content-Type: application/json" \
// -d '{
//     "city": "Jaipur",
//     "temp": 35.5,
//     "humidity": 60,
//     "timestamp": "2024-04-06T14:30:00Z"
// }'

package entities

type KY026 struct {
	ID              int    `json:"id"`
	FechaActivacion string `json:"fecha_activacion"`
	Estado          int    `json:"estado"`
}

type MQ2 struct {
	ID              int    `json:"id"`
	FechaActivacion string `json:"fecha_activacion"`
	Estado          int    `json:"estado"`
}

type ESP32 struct {
	ID      int   `json:"id"`
	KY026ID int   `json:"ky026_id"`
	MQ2ID   int   `json:"mq2_id"`
	KY026   KY026 `json:"ky026"`
	MQ2     MQ2   `json:"mq2"`
}

package repositories

import "ApiMulti/src/domain/entities"

type SensorRepository interface {
	SaveKY026(sensor *entities.KY026) error
	SaveMQ2(sensor *entities.MQ2) error
	GetKY026ByID(id int) (*entities.KY026, error)
	GetMQ2ByID(id int) (*entities.MQ2, error)
	SaveESP32(esp32 *entities.ESP32) error
	GetAllKY026() ([]*entities.KY026, error)
	GetAllMQ2() ([]*entities.MQ2, error)
}

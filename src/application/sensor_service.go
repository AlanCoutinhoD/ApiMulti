package application

import (
	"ApiMulti/src/core"
	"ApiMulti/src/domain/entities"
	"ApiMulti/src/domain/repositories"
	"encoding/json"
	"os"
	"time"
)

type SensorService struct {
	repo repositories.SensorRepository
}

func NewSensorService(repo repositories.SensorRepository) *SensorService {
	return &SensorService{repo: repo}
}

func (s *SensorService) ProcessKY026Reading(estado int) error {
	sensor := &entities.KY026{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveKY026(sensor); err != nil {
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		return err
	}

	return core.PublishMessage(os.Getenv("RABBITMQ_QUEUE_KY026"), message)
}

func (s *SensorService) ProcessMQ2Reading(estado int) error {
	sensor := &entities.MQ2{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveMQ2(sensor); err != nil {
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		return err
	}

	return core.PublishMessage(os.Getenv("RABBITMQ_QUEUE_MQ2"), message)
}

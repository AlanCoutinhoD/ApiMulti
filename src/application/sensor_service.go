package application

import (
	"ApiMulti/src/core"
	"ApiMulti/src/domain/entities"
	"ApiMulti/src/domain/repositories"
	"encoding/json"
	"log"
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
		log.Printf("Error saving KY026 to database: %v", err)
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("Error marshaling KY026 message: %v", err)
		return err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE_KY026")
	if err := core.PublishMessage(queueName, message); err != nil {
		log.Printf("Error publishing KY026 message: %v", err)
		return err
	}

	log.Printf("Successfully processed KY026 reading: %+v", sensor)
	return nil
}

func (s *SensorService) ProcessMQ2Reading(estado int) error {
	sensor := &entities.MQ2{
		FechaActivacion: time.Now().Format("2006-01-02 15:04:05"),
		Estado:          estado,
	}

	if err := s.repo.SaveMQ2(sensor); err != nil {
		log.Printf("Error saving MQ2 to database: %v", err)
		return err
	}

	// Publish to RabbitMQ
	message, err := json.Marshal(sensor)
	if err != nil {
		log.Printf("Error marshaling MQ2 message: %v", err)
		return err
	}

	queueName := os.Getenv("RABBITMQ_QUEUE_MQ2")
	if err := core.PublishMessage(queueName, message); err != nil {
		log.Printf("Error publishing MQ2 message: %v", err)
		return err
	}

	log.Printf("Successfully processed MQ2 reading: %+v", sensor)
	return nil
}

package repositories

import (
	"ApiMulti/src/domain/entities"
	"database/sql"
	"errors"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) SaveKY026(sensor *entities.KY026) error {
	query := "INSERT INTO KY_026 (fecha_activacion, estado) VALUES (?, ?)"
	result, err := r.db.Exec(query, sensor.FechaActivacion, sensor.Estado)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	sensor.ID = int(id)
	return nil
}

func (r *MySQLRepository) SaveMQ2(sensor *entities.MQ2) error {
	query := "INSERT INTO MQ_2 (fecha_activacion, estado) VALUES (?, ?)"
	result, err := r.db.Exec(query, sensor.FechaActivacion, sensor.Estado)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	sensor.ID = int(id)
	return nil
}

func (r *MySQLRepository) GetKY026ByID(id int) (*entities.KY026, error) {
	sensor := &entities.KY026{}
	query := "SELECT idKY_026, fecha_activacion, estado FROM KY_026 WHERE idKY_026 = ?"
	err := r.db.QueryRow(query, id).Scan(&sensor.ID, &sensor.FechaActivacion, &sensor.Estado)
	if err == sql.ErrNoRows {
		return nil, errors.New("sensor not found")
	}
	if err != nil {
		return nil, err
	}
	return sensor, nil
}

func (r *MySQLRepository) GetMQ2ByID(id int) (*entities.MQ2, error) {
	sensor := &entities.MQ2{}
	query := "SELECT idMQ_2, fecha_activacion, estado FROM MQ_2 WHERE idMQ_2 = ?"
	err := r.db.QueryRow(query, id).Scan(&sensor.ID, &sensor.FechaActivacion, &sensor.Estado)
	if err == sql.ErrNoRows {
		return nil, errors.New("sensor not found")
	}
	if err != nil {
		return nil, err
	}
	return sensor, nil
}

func (r *MySQLRepository) SaveESP32(esp32 *entities.ESP32) error {
	query := "INSERT INTO ESP32 (idKY_026, idMQ_2) VALUES (?, ?)"
	result, err := r.db.Exec(query, esp32.KY026ID, esp32.MQ2ID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	esp32.ID = int(id)
	return nil
}

func (r *MySQLRepository) GetAllKY026() ([]*entities.KY026, error) {
	query := "SELECT idKY_026, fecha_activacion, estado FROM KY_026 ORDER BY fecha_activacion DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.KY026
	for rows.Next() {
		sensor := &entities.KY026{}
		if err := rows.Scan(&sensor.ID, &sensor.FechaActivacion, &sensor.Estado); err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}

func (r *MySQLRepository) GetAllMQ2() ([]*entities.MQ2, error) {
	query := "SELECT idMQ_2, fecha_activacion, estado FROM MQ_2 ORDER BY fecha_activacion DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.MQ2
	for rows.Next() {
		sensor := &entities.MQ2{}
		if err := rows.Scan(&sensor.ID, &sensor.FechaActivacion, &sensor.Estado); err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}

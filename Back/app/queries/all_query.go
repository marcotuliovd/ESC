package queries

import (
	"Back/app/models"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from Book model.
type VagaQueries struct {
	*sqlx.DB
}

func (q *VagaQueries) CreateUser(u *models.Clients) (uuid.UUID, error) {
	u.Id = uuid.New()
	queryX := `INSERT INTO clients (id, username, name, phone) VALUES ($1, $2, $3, $4)`
	_, err := q.Exec(queryX, u.Id, u.Username, u.Name, u.Phone)

	if err != nil {
		return u.Id, err
	}

	return u.Id, nil
}

func (q *VagaQueries) CreateVehicle(v *models.Vehicle) (uuid.UUID, error) {
	v.Id = uuid.New()
	queryX := `INSERT INTO vehicle (id, owner, plate, title, type) VALUES ($1, $2, $3, $4, $5)`
	_, err := q.Exec(queryX, v.Id, v.Owner, v.Plate, v.Title, v.Type)

	if err != nil {
		return v.Id, err
	}

	return v.Id, nil
}


func (q *VagaQueries) OccupationSpace(s *models.Space) (uuid.UUID, error) {
	s.Id = uuid.New()
	queryX := `INSERT INTO space (id, type, vehicle_id) VALUES ($1, $2, $3)`
	_, err := q.Exec(queryX, s.Id, s.Type, s.VehicleId)

	if err != nil {
		return s.Id, err
	}

	return s.Id, nil
}

func (q *VagaQueries) CreateHistory(history *models.History) (error) {
	history.Id = uuid.New()
	queryX := `INSERT INTO history (id, amount, vehicle_id, entry, exit, type) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := q.Exec(queryX, history.Id, history.Amount, history.VehicleId, history.Entry, history.Exit, history.Type)
	if err != nil {
		return err
	}
	return nil
}

func (q *VagaQueries) GetVagaByVehicleId(VehicleId uuid.UUID) (*models.Space, error) {
	query := `SELECT id, vehicle_id, type, created_at from public.space WHERE vehicle_id = $1`
	occupation := &models.Space{}
	err := q.Get(occupation, query, VehicleId)
	if err != nil {
		return nil, err
	}
	return occupation, nil
}

func (q *VagaQueries) DeleteOccupation(SpaceId uuid.UUID) (error) {
	query := `Delete from space WHERE id = $1`
	_, err := q.Exec(query, SpaceId)
	if err != nil {
		return err
	}
	return nil
}

func (q *VagaQueries) ServiceReport(init time.Time, finish time.Time) ([]models.HistoryReturn,error) {
	query := `SELECT id, amount, vehicle_id, entry, exit, type, created_at
		FROM public.history
		WHERE created_at >= $1 AND created_at <= $2;`
	occupation := []models.HistoryReturn{}
	err := q.Select(&occupation, query, init, finish)
	if err != nil {
		return nil, err
	}
	return occupation, nil
}
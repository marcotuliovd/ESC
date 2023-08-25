package queries

import (
	"Back/app/models"

	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from Book model.
type VagaQueries struct {
	*sqlx.DB
}


// CreateBook method for creating book by given Book object.
func (q *VagaQueries) EntradaDeVeiculo(b *models.Vaga) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b., b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

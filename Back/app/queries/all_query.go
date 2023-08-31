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
	query := `UPDATE public.vagas VALUES dono = $2, telefone = $3, tipo_de-veiculo = $4, placa = $5, modelo = $6, $7, $8 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, b.Id, b.Dono, b.Telefone, b.TipoVeiculo, b.Placa, b.Modelo)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

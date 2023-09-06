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
	query := `UPDATE vagas SET dono = $1, telefone = $2, tipo_de_veiculo = $3, placa = $4, modelo = $5 WHERE id = $6`

	// Send query to database.
	_, err := q.Exec(query, b.Dono, b.Telefone, b.TipoVeiculo, b.Placa, b.Modelo, b.Id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

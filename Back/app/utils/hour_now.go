package utils

import (
	"time"
)

func HourNow() time.Time{
	// Configurar o fuso horário desejado (UTC-3)
	fusoHorario := "America/Sao_Paulo"
	loc, _ := time.LoadLocation(fusoHorario)

	// Obter a hora atual no fuso horário desejado
	HourNow := time.Now().In(loc)

	return HourNow
}

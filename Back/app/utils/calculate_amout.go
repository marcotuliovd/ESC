package utils

func CalculateAmount(Type string, time int) (float32) {
	var price float32

	switch Type {
	case "pequeno":
		price = 5.00
	case "medio":
		price = 7.00
	case "grande":
		price = 10.00
	case "moto":
		price = 8.00
	default:
		price = 10.00
	}

    if time >= 600 {
        time = 1440
    }

	if time <= 60 {
		return price
	}

	time = time - 60

	return (float32(time) * 0.05) + price
}

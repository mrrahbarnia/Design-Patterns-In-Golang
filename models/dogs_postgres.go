package models

import (
	"context"
	"fmt"
	"time"
)

func (p *postgresRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn()

	query := `
		SELECT
			id,
			breed,
			weight_low_lbs,
			weight_high_lbs,
			average_weight,
			lifespan,
			description,
			alternate_names,
			geographic_origin
		FROM dog_breeds
	`
	rows, err := p.DB.QueryContext(ctx, query)
	fmt.Printf("Rows: %v\n", rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dogBreeds []*DogBreed

	for rows.Next() {
		var d DogBreed
		if err := rows.Scan(
			&d.ID,
			&d.Breed,
			&d.WeightLowLbs,
			&d.WeightHighLbs,
			&d.AverageWeight,
			&d.LifeSpan,
			&d.Description,
			&d.AlternateNames,
			&d.GeographicOrigin,
		); err != nil {
			return nil, err
		}
		dogBreeds = append(dogBreeds, &d)
	}
	return dogBreeds, nil
}

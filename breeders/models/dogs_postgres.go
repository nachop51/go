package models

import (
	"context"
	"time"
)

func (repo *postgresRepo) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs, (weight_low_lbs + weight_high_lbs) / 2 as average_weight, lifespan, COALESCE(details, ''), COALESCE(alternate_names, ''), COALESCE(geographic_origin, '') FROM dog_breeds ORDER BY breed`

	var dogBreeds []*DogBreed

	rows, err := repo.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		d := &DogBreed{}

		err := rows.Scan(
			&d.Id,
			&d.Breed,
			&d.WeightLowLbs,
			&d.WeightHighLbs,
			&d.AverageWeight,
			&d.LifeSpan,
			&d.Details,
			&d.AlternateNames,
			&d.GeographicOrigin,
		)

		if err != nil {
			return nil, err
		}

		dogBreeds = append(dogBreeds, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dogBreeds, nil
}

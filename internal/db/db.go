package db

type DB interface {
	// Upsert inserts or updates a vector.
	Upsert(vector []float64, id string) error

	// Delete deletes a vector.
	Delete(id string) error

	// Fetch fetches a vector.
	Fetch(id string) ([]float64, error)

	// Query queries the database for the nearest neighbors of the given vector.
	Query(vector []float64, k int) ([]string, error)

	// Update updates the vector of the given id.
	Update(vector []float64, id string) error

	// DescribeIndexStats returns the stats of the index.
	DescribeIndexStats() (map[string]interface{}, error)
}

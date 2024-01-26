package math

import (
	"errors"
	"math"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
// It returns an error if the vectors are of different lengths or are empty.
func CosineSimilarity(vec1, vec2 []float64) (float64, error) {
	if len(vec1) != len(vec2) {
		return 0, errors.New("two vectors must be of the same length")
	}

	var dotProduct float64
	var magnitude1, magnitude2 float64

	for i := range vec1 {
		dotProduct += vec1[i] * vec2[i]

		// keep track of magnitude separately to avoid unnecessary
		// multiplication and addition operations in the same loop
		// to avoid overhead of a dedicated function call
		magnitude1 += vec1[i] * vec1[i]
		magnitude2 += vec2[i] * vec2[i]
	}

	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	if magnitude1 == 0 || magnitude2 == 0 {
		return 0, errors.New("vectors must not be zero vectors")
	}

	cosineSimilarity := dotProduct / (magnitude1 * magnitude2)
	return cosineSimilarity, nil
}

// cosineDistance calculates the cosine distance between two vectors.
// It returns an error if the vectors are of different lengths or are empty.
func CosineDistance(vec1, vec2 []float64) (float64, error) {
	cosineDistance, err := CosineSimilarity(vec1, vec2)
	if err != nil {
		return 0, err
	}
	return 1 - cosineDistance, nil
}

// EuclideanDistance calculates the euclidean distance between two vectors.
func EuclideanDistance(vec1, vec2 []float64) (float64, error) {
	if len(vec1) != len(vec2) {
		return 0, errors.New("two vectors must be of the same length")
	}

	var sum float64

	for i := range vec1 {
		sum += math.Pow(vec1[i]-vec2[i], 2)
	}

	return math.Sqrt(sum), nil
}

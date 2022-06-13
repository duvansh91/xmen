package models

// Human represents a human entity.
type Human struct {
	DNA      []string `bson:"dna"`
	IsMutant bool     `bson:"is_mutant"`
}

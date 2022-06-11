package models

type Human struct {
	DNA      []string `bson:"dna"`
	IsMutant bool     `bson:"is_mutant"`
}

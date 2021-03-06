package models

// Stats represents a stats entity.
type Stats struct {
	CountMutantDNA int16   `json:"count_mutant_dna"`
	CountHumanDNA  int16   `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

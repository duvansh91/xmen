package models

type Stats struct {
	CountMutantDNA int16   `json:"count_mutant_dna"`
	CountHumanDNA  int16   `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}

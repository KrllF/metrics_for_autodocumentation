package entity

type Stat struct {
	UnCovered int64
	InCorrect int64
	Coverage  float64
}

type StructStat struct {
	CoverageStruct   float64
	OkCoverageStruct bool
}

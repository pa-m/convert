package convert

// Matrix sub interface of gonum's one
type Matrix interface {
	Dims() (r, c int)
	At(i, j int) float64
}

// Vector sub interface of gonum's one
type Vector interface {
	Len() int
	AtVec(i int) float64
}

// Dataer sub interface of gorgonia's one
type Dataer interface {
	Data() interface{}
}

package main

type person struct {
	name string
	age  int
}
type worker struct {
	person *person
	position          string
	yearsOfExperience int
	education         string
}
type doctor struct {
	Speciality string `json:"speciality"`
	*worker
}
type teacher struct {
	AcademicDegree string `json:"academicDegree"`
	*worker
}
type artist struct {
	PaintingStyle string `json:"paintingStyle"`
	*worker
}
type developer struct {
	ProgrammingLanguage string `json:"programmingLanguage"`
	*worker
}
type driver struct {
	DrivingCategory string `json:"drivingCategory"`
	*worker
}

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type worker struct {
	person            *person
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

type Worker interface {
	GetNamePosition() string
}

func GetInfo(w Worker) (info string) {
	info = w.GetNamePosition()
	return
}

func main() {

	var worker1 = doctor{
		Speciality: "dentist",
		worker: &worker{
			person: &person{
				name: "Mary",
				age:  54,
			},
			position:          "children dentist",
			yearsOfExperience: 24,
			education:         "Kiev Medical University",
		},
	}

	var worker2 = teacher{
		AcademicDegree: "doctor of science",
		worker: &worker{
			person: &person{
				name: "Alex",
				age:  64,
			},
			position:          "university teacher",
			yearsOfExperience: 40,
			education:         "Kiev  University",
		},
	}
	var worker3 = artist{
		PaintingStyle: "graphic art",
		worker: &worker{
			person: &person{
				name: "Natalie",
				age:  19,
			},
			position:          "graphic artist",
			yearsOfExperience: 1,
			education:         "NUBIP",
		},
	}
	var worker4 = developer{
		ProgrammingLanguage: "Golang",
		worker: &worker{
			person: &person{
				name: "Natalie",
				age:  19,
			},
			position:          "trainee",
			yearsOfExperience: 0,
			education:         "NUBIP",
		},
	}
	var worker5 = driver{
		DrivingCategory: "B",
		worker: &worker{
			person: &person{
				name: "Vasya",
				age:  54,
			},
			position:          "taxi driver",
			yearsOfExperience: 24,
			education:         "KNU ",
		},
	}

	var workers = make(map[int]Worker)
	{
		workers[0] = worker1
		workers[1] = worker2
		workers[2] = worker3
		workers[3] = worker4
		workers[4] = worker5
	}
	//fmt.Println(workers[2])
	for _, worker := range workers {
		fmt.Println(GetInfo(worker))
	}


}

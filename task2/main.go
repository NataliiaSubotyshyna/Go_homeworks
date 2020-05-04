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

	//Операції з вказівниками
	var worker1pointer *doctor = &worker1     //показник на існуючу структуру
	var agePointer *int = &worker2.person.age // показник на поле існуючої структури

	worker1pointer.person.age = 20
	*agePointer = 33

	fmt.Println(worker1.person.age)
	fmt.Println(worker2.person.age)

	//Виклик методів
	fmt.Println(worker1.person.age)
	worker1.person.SetAge()
	fmt.Println(worker1.person.age)
	worker1.person.HappyBirthday()

	//Виклик паніки
	ann := &person{}
	ann = nil
	ann.Hello()

	//інформація про працівників
	fmt.Println("Worker1\n" + worker1.fullInfo() + "\n")
	fmt.Println("Worker2\n" + worker2.fullInfo() + "\n")
	fmt.Println("Worker3\n" + worker3.fullInfo() + "\n")
	fmt.Println("Worker4\n" + worker4.fullInfo() + "\n")
	fmt.Println("Worker5\n" + worker5.fullInfo() + "\n")

}

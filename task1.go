package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}
type worker struct {
	person
	position          string
	yearsOfExperience int
	education         string
}
type doctor struct {
	Speciality string `tag:"speciality"`
	worker
}
type teacher struct {
	AcademicDegree string `tag:"academicDegree"`
	worker
}
type artist struct {
	PaintingStyle string `tag:"paintingStyle"`
	worker
}
type developer struct {
	ProgrammingLanguage string `tag:"programmingLanguage"`
	worker
}
type driver struct {
	DrivingCategory string `tag:"drivingCategory"`
	worker
}

func main() {
	worker1 := doctor{
		Speciality: "dentist",
		worker: worker{
			person: person{
				name: "Mary",
				age:  54,
			},
			position:          "children dentist",
			yearsOfExperience: 24,
			education:         "Kiev Medical University",
		},
	}
	worker2 := teacher{
		AcademicDegree: "doctor of science",
		worker: worker{
			person: person{
				name: "Alex",
				age:  64,
			},
			position:          "university teacher",
			yearsOfExperience: 40,
			education:         "Kiev  University",
		},
	}
	worker3 := artist{
		PaintingStyle: "graphic art",
		worker: worker{
			person: person{
				name: "Natalie",
				age:  19,
			},
			position:          "graphic artist",
			yearsOfExperience: 1,
			education:         "NUBIP",
		},
	}
	worker4 := developer{
		ProgrammingLanguage: "Golang",
		worker: worker{
			person: person{
				name: "Natalie",
				age:  19,
			},
			position:          "trainee",
			yearsOfExperience: 0,
			education:         "NUBIP",
		},
	}
	worker5 := driver{
		DrivingCategory: "B",
		worker: worker{
			person: person{
				name: "Vasya",
				age:  54,
			},
			position:          "taxi driver",
			yearsOfExperience: 24,
			education:         "KNU ",
		},
	}
	fmt.Println("Worker1\n" + worker1.fullInfo() + "\n" +worker1.profInfo() +"\n")
	fmt.Println("Worker2\n" +worker2.fullInfo() + "\n" +worker2.profInfo() +"\n")
	fmt.Println("Worker3\n" +worker3.fullInfo() + "\n" +worker3.profInfo() +"\n")
	fmt.Println("Worker4\n" +worker4.fullInfo() + "\n" +worker4.profInfo() +"\n")
	fmt.Println("Worker5\n" +worker5.fullInfo() + "\n" +worker5.profInfo() +"\n")

}

func (d doctor) profInfo() string{
	return "speciality:" + d.Speciality
}
func (t teacher) profInfo() string{
	return "Academic Degree:" + t.AcademicDegree
}
func (a artist) profInfo() string{
	return "Painting Style:" + a.PaintingStyle
}
func (dev developer) profInfo() string{
	return "Programming Language:" + dev.ProgrammingLanguage
}
func (dr driver) profInfo() string{
	return "Driving Category:" + dr.DrivingCategory
}

func (w worker) fullInfo() string {

	return "name:" + w.name + "\nage:" + strconv.Itoa(w.age) + "\nposition:" + w.position + "\neducation:" + w.education +
		"\nyears of experience:" + strconv.Itoa(w.yearsOfExperience)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Виклик методів
	fmt.Println(worker1.person.age)
	worker1.person.SetAge()
	fmt.Println(worker1.person.age)
	worker1.person.HappyBirthday()

	//Операції з вказівниками
	worker1pointer.person.age = 20
	*agePointer = 33

    fmt.Println(worker1.person.age)
	fmt.Println(worker2.person.age)

	//Виклик паніки
	ann := &person{}
	ann = nil
	ann.Hello()

	fmt.Println("Worker1\n"+ worker1.fullInfo() +"\n")
	fmt.Println("Worker2\n" +worker2.fullInfo() +"\n")
	fmt.Println("Worker3\n" +worker3.fullInfo() +"\n")
	fmt.Println("Worker4\n" +worker4.fullInfo() +"\n")
	fmt.Println("Worker5\n" +worker5.fullInfo() +"\n")

}

func (w worker) fullInfo() string {
	return "name:" + w.person.name + "\nage:" + strconv.Itoa(w.person.age) + "\nposition:" + w.position + "\neducation:" + w.education +
		"\nyears of experience:" + strconv.Itoa(w.yearsOfExperience)
}
func (p *person) Hello()  {
	fmt.Println("Hello,", p.name)
}

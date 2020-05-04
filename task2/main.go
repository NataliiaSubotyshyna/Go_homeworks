package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Виклик методів
	fmt.Println(worker1.age)
	worker1.SetAge()
	fmt.Println(worker1.age)
	worker1.HappyBirthday()

	//Операції з вказівниками
	worker1pointer.age = 20
	*agePointer = 33

    fmt.Println(worker1.age)
	fmt.Println(worker2.age)

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
	return "name:" + w.name + "\nage:" + strconv.Itoa(w.age) + "\nposition:" + w.position + "\neducation:" + w.education +
		"\nyears of experience:" + strconv.Itoa(w.yearsOfExperience)
}
func (p *person) Hello()  {
	fmt.Println("Hello,", p.name)
}

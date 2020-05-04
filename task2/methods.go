package main

import (
	"fmt"
	"strconv"
)

func (w worker) fullInfo() string {
	return "name:" + w.person.name + "\nage:" + strconv.Itoa(w.person.age) + "\nposition:" + w.position + "\neducation:" + w.education +
		"\nyears of experience:" + strconv.Itoa(w.yearsOfExperience)
}
func (p *person) Hello() {
	fmt.Println("Hello,", p.name)
}
func (person *person) SetAge() {
	person.age = person.age + 1
}

func (receiver *person) HappyBirthday() {
	fmt.Println("Happy "+strconv.Itoa(receiver.age)+" birthday,", receiver.name)
}

package main

import (
	"fmt"
"strconv"
)

func (person *person) SetAge(){
	person.age = person.age + 1
}

func (receiver *person) HappyBirthday(){
	fmt.Println("Happy " + strconv.Itoa(receiver.age) + " birthday,", receiver.name)
}


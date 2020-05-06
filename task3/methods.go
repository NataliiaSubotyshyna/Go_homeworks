package main

import (
	"fmt"
	"strconv"
)

func (receiver *person) HappyBirthday() {
	fmt.Println("Happy "+strconv.Itoa(receiver.age)+" birthday,", receiver.name)
}
func (d doctor) GetNamePosition() string {
	return d.person.name + ":" + d.position
}
func (t teacher) GetNamePosition() string {
	return t.person.name+":" + t.position
}
func (dev developer) GetNamePosition() string {
	return dev.person.name+":" + dev.position
}
func (dr driver) GetNamePosition() string {
	return dr.person.name+":" + dr.position
}
func (a artist) GetNamePosition() string {
	return a.person.name+":" + a.position
}
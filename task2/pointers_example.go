package main

	var worker1pointer *doctor = &worker1 //показник на існуючу структуру
	var agePointer *int = &worker2.age // показник на поле існуючої структури
	var tom *person = &person{name: "tom",age: 45} //показник на структуру з ініціалізацією полів

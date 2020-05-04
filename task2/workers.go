package main

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

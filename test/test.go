package test

type student struct {
	Name   string
	age    int
	salary float64
}

func NewStudent(name string, age int, salary float64) *student {
	return &student{
		Name:   name,
		age:    age,
		salary: salary,
	}
}
func (stu student)GetSalary()float64{
	return stu.salary
}
func (stu *student)SetSalary(salary float64)bool{
	if salary<0{
		return false
	}
	stu.salary=salary
	return true
}
func (stu *student)GetAge()int{
	return stu.age
}
func (stu *student) SetAge(age int) bool {
	if age < 0 || age > 130 {
		return false
	}
	stu.age = age
	return true
}

package academy

import (
	"math"
	"sort"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(grades); i++ {
		sum += grades[i]
	}
	return int(math.Round(float64(sum) / float64(len(grades))))
	panic("not implemented")
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	if len(attendance) == 0 {
		return 0
	}
	sum := float64(0)
	for i := 0; i < len(attendance); i++ {
		if attendance[i] == true {
			sum += 1
		}
	}
	return sum / float64(len(attendance))
	panic("not implemented")
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	sort.Ints(s.Grades)
	if s.Project == 1 || s.Grades[0] == 1 || AttendancePercentage(s.Attendance) < 0.6 {
		return 1
	}
	Final := int(math.Round((float64(s.Project) + float64(AverageGrade(s.Grades))) / 2))
	if AttendancePercentage(s.Attendance) < 0.8 {
		return Final - 1
	}
	return Final
	panic("not implemented")
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	m := map[string]uint8{}
	for _, student := range students {
		m[student.Name] = uint8(FinalGrade(student))
	}
	return m
	panic("not implemented")
}

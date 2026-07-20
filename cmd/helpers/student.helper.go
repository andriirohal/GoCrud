package helpers

import (
  "practice_go/cmd/models" 
  
  "github.com/jackc/pgx/v5"
);

func ScanStudent(row pgx.Row) (models.Student, error) {
  var student models.Student

  err := row.Scan(
	&student.Id,
	&student.Name,
	&student.Age,
	&student.Grade,
	&student.Active,
  );

  return student, err
};

func IsAgeValid(age int) bool {
  return age >= 7	
};

func IsGradeValid(grade float64) bool {
  return grade >= 0 && grade <= 100;
};
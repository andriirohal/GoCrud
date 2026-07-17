package services

import (
	"practice_go/cmd/db"
	"practice_go/cmd/helpers"
	"practice_go/cmd/models"
	"context"
	"errors"
	"strings"
);

func CreateStudent(student models.Student) (models.Student, error) {
  name := strings.TrimSpace(student.Name);
	
  if name == "" {
	return models.Student{}, errors.New("Invalid student name");
  };

  if !helpers.IsAgeValid(student.Age) {
    return models.Student{}, errors.New("Invalid student age");
  };

  if !helpers.IsGradeValid(student.Grade) {
	return models.Student{}, errors.New("Invalid student grade");
  };

  newStudent, err := helpers.ScanStudent(db.DB.QueryRow(context.Background(), "INSERT INTO students (name, age, grade) VALUES ($1, $2, $3) RETURNING id, name, age, grade, active",
    name, student.Age, student.Grade, 
    ),
  );

  if err != nil {
    return models.Student{}, err;
  };
  
  return newStudent, nil;
};

func UpdateStudent(id string, student models.UpdateStudent) (models.Student, error) {
  if student.Name != nil {
    name := strings.TrimSpace(*student.Name);
  
    if name == "" {
      return models.Student{}, errors.New("Invalid student name");
    };

    *student.Name = name;
  };

  if student.Age != nil {
    if !helpers.IsAgeValid(*student.Age) {
      return models.Student{}, errors.New("Invalid student age");
    };
  };
  
  if student.Grade != nil {
    if !helpers.IsGradeValid(*student.Grade) {
      return models.Student{}, errors.New("Invalid student grade");
    };
  };
  
  updatedStudent, err := helpers.ScanStudent(db.DB.QueryRow(context.Background(), "UPDATE students SET name = COALESCE($2, name), age = COALESCE($3, age), grade = COALESCE($4, grade), active = COALESCE($5, active) WHERE id = $1 RETURNING id, name, age, grade, active",
    id, student.Name, student.Age, student.Grade, student.Active,
    ),
  );

  if err != nil {
    return models.Student{}, err;
  };

  return updatedStudent, nil;
};

func DeleteStudent(id string) (models.Student, error) {
  deletedStudent, err := helpers.ScanStudent(db.DB.QueryRow(context.Background(), "DELETE from students WHERE id = $1 RETURNING id, name, age, grade, active",
    id,
    ),
  );

  if err != nil {
    return models.Student{}, err;
  };

  return deletedStudent, nil;
};

func GetStudentById(id string) (models.Student, error) {
  student, err := helpers.ScanStudent(db.DB.QueryRow(context.Background(), "SELECT id, name, age, grade, active FROM students WHERE id = $1",
    id,
    ),
  ); 

  if err != nil {
    return models.Student{}, err;
  };

  return student, nil;
};

func GetAllStudents() ([]models.Student, error) {
  rows, err := db.DB.Query(context.Background(), "SELECT id, name, age, grade, active FROM students");

  if err != nil {
    return nil, err;
  };

  defer rows.Close();

  var students []models.Student;

  for rows.Next() {
    var student models.Student;

    err := rows.Scan(
      &student.Id,
      &student.Name,
      &student.Age,
      &student.Grade,
      &student.Active,
    );

    if err != nil {
      return nil, err;
    };

    students = append(students, student);
  };

  if err := rows.Err(); err != nil {
    return nil, err;
  };

  return students, nil;
};
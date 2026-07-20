package controllers

import (
  "GoCrud/cmd/models"
  "GoCrud/cmd/services"
  
  "net/http"
  "github.com/gin-gonic/gin"
);

func CreateStudentController(c *gin.Context) {
  var student models.Student;

  if err := c.ShouldBindJSON(&student); err != nil {
	c.JSON(http.StatusBadRequest, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  createdStudent, err := services.CreateStudent(student);

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  c.JSON(http.StatusOK, createdStudent);
};

func GetStudentByIdController(c *gin.Context) {
  id := c.Param("id");
	
  student, err := services.GetStudentById(id);

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  c.JSON(http.StatusOK, student);
};

func DeleteStudentController(c *gin.Context) {
  id := c.Param("id");
  
  deletedStudent, err := services.DeleteStudent(id);

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  c.JSON(http.StatusOK, deletedStudent);
};

func UpdateStudentController(c *gin.Context) {
  id := c.Param("id");

  var input models.UpdateStudent;

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  updatedStudentGrade, err := services.UpdateStudent(id, input);

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  c.JSON(http.StatusOK, updatedStudentGrade);
};

func GetAllStudentsController(c *gin.Context) {
  students, err := services.GetAllStudents(); 
  
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H {
      "error": err.Error(),
	  });
	  return;
  };

  c.JSON(http.StatusOK, students);
};
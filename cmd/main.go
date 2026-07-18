package main

import (
  "practice_go/cmd/controllers"
	"practice_go/cmd/db"
  
	"log" 
  "github.com/gin-gonic/gin"
);

func main() {
  if err := db.Connect(); err != nil {
    panic(err);
  };

  defer db.DB.Close();

  router := gin.Default();

  router.GET("/students", controllers.GetAllStudentsController);
  router.GET("/students/:id", controllers.GetStudentByIdController);

  router.POST("/students", controllers.CreateStudentController);

  router.DELETE("/students/:id", controllers.DeleteStudentController);

  router.PATCH("/students/:id", controllers.UpdateStudentController);

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err);
  };
};
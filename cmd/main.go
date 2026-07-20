package main

import (
  "GoCrud/cmd/controllers"
  "GoCrud/cmd/db"
  "GoCrud/cmd/middlewares"

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
  router.GET("/students/:id", middlewares.ValidateId, controllers.GetStudentByIdController);

  router.POST("/students", controllers.CreateStudentController);

  router.DELETE("/students/:id", middlewares.ValidateId, controllers.DeleteStudentController);

  router.PATCH("/students/:id", middlewares.ValidateId, controllers.UpdateStudentController);

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err);
  };
};
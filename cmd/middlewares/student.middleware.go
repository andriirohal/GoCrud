package middlewares

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
);

func ValidateId(c *gin.Context) {
  id := c.Param("id");

  if _, err := uuid.Parse(id); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
      "error": "Invalid student ID",
	  });
	  return;
  };

  c.Next();
};
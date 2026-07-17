package models

type Student struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Age int `json:"age"`
  Grade float64 `json:"grade"`
  Active bool `json:"active"`
};

type UpdateStudent struct {
  Id *string `json:"id"`
  Name *string `json:"name"`
  Age *int `json:"age"`
  Grade *float64 `json:"grade"`
  Active *bool `json:"actice"`
};
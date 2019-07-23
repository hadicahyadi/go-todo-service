package models

type TodoItem struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Status bool `json:"status"`
}

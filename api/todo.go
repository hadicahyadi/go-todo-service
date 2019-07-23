package api

import(
	"net/http"
  "log"
  "encoding/json"

	"github.com/gorilla/mux"
  "github.com/hadicahyadi/todo-service/models"
  "github.com/hadicahyadi/todo-service/db"
)

func Router() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/todo-list", getTodoList).Methods("GET")
  router.HandleFunc("/todo", saveTodo).Methods("POST")
  
  return router
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
  dbConn := db.Connect()
	defer dbConn.Close()

  rows, err := dbConn.Query("SELECT id, name, status from todo_item")
  if err != nil {
		log.Print(err)
	}

  var item models.TodoItem
  var todos []models.TodoItem
	var response models.Response

  for rows.Next() {
    if err := rows.Scan(&item.ID, &item.Name, &item.Status); err != nil {
			log.Print(err)
		}
		todos = append(todos, item)
  }

  response.Status = 1
	response.Message = "Success"
	response.Data = todos

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func saveTodo(w http.ResponseWriter, r *http.Request) {
  var todo models.TodoItem
  var response models.Response
  json.NewDecoder(r.Body).Decode(&todo)

  dbConn := db.Connect()
  defer dbConn.Close()

  _, err := dbConn.Exec("INSERT INTO todo_item (name, status) values (?, ?)",
    todo.Name,
    todo.Status,
  )

  if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Item has been saved."
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

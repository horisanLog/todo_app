package main

import (
	"fmt"
	// "log"
	// "log"
	"todo_app/app/controllers"
	"todo_app/app/models"
	// "todo_app/config"
	// "todo_app_heroku/app/controllers"
	// "todo_app_heroku/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)

	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

  // t, _ := models.GetTodo(1)
	// fmt.Println(t) 

	// user, _ := models.GetUser(2)
  // user.CreateTodo("Second Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user, _ := models.GetUser(3)
  // user.CreateTodo("Third Todo")

	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
  // for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)

	// t.Content = "Update Todo"
	// t.UpdateTodo()

	// t, _ := models.GetTodo(5)
	// t.DeleteTodo()

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test@test.com")

	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}


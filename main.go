package main

func main() {
	firstTodo := NewTodo("First Todo", "This is the first todo")
	todos := AddTodos(Todos{}, firstTodo)
	secondTodo := NewTodo("Second Todo", "This is the second todo")
	todos = AddTodos(todos, secondTodo)
	thirdTodo := NewTodo("Third Todo", "This is the third todo")
	todos = AddTodos(todos, thirdTodo)
	fourthTodo := NewTodo("Fourth Todo", "This is the fourth todo")
	todos = AddTodos(todos, fourthTodo)
	ListTodos(todos)

	todos = DeleteTodo(todos, 2)

	ListTodos(todos)
	fifthTodo := NewTodo("Fifth Todo", "This is the fifth todo")
	todos = AddTodos(todos, fifthTodo)
	ListTodos(todos)

	todos = DeleteTodo(todos, 5)
	ListTodos(todos)
	todos = AddTodos(todos, fifthTodo)
	ListTodos(todos)

}

package todo

type TodoList struct { //структура списка дел
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct { //связь между пользователями и списком дел
	Id     int
	UserId int
	ListId int
}

type TodoItem struct { //структура задач
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Decription string `json:"description"`
	Done       bool   `json:"done"`
}

type ListsItem struct { //связь между списками и задачами
	Id     int
	ListId int
	ItemId int
}

package main

/*
	# Разрешения бота
	# Настройка в функции init()
	# Пример: 1234567: ROLE_MODERATOR,
	# 89101112: ROLE_ADMIN,

	# ROLE_NOROLE - Нет роли, это не нужно трогать
	# getRoleName() возращает имя роли, его можно редактировать
*/

func init(){
	UsersDB = map[int]int {
		104002857: ROLE_ADMIN, //роли
	}
}

const (
	ROLE_NOROLE = iota //нет роли
	ROLE_MODERATOR //1
	ROLE_ADMIN //2
)

var UsersDB map[int]int

func getRoleName(roleid int) string {
	switch roleid {
	case ROLE_NOROLE:
		return "Без роли"
	case ROLE_MODERATOR:
		return "Модератор"
	case ROLE_ADMIN:
		return "Админ"
	}
	return "Неизвестная роль"
}

func getUserRole(userid int) int { //роль, есть или нет
	role, _ := UsersDB[userid]
	return role  //0 если нет
}




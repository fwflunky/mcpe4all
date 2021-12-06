package main

import (
	"github.com/SevereCloud/vksdk/v2/events"
	"strings"
)

func handleMessage(obj events.MessageNewObject, userid, peerid int, message string){
	cmd := CommandFromString(message)
	switch cmd.GetName() {
	case "help", "помощь", "помргите":
		if role := getUserRole(userid); role < ROLE_MODERATOR { //можно выполнять от ROLE_MODERATOR (1)
			sendTo(peerid, "🚫 Недостаточно прав для выполнения этой команды. Требуется роль " + getRoleName(ROLE_MODERATOR) + ", Ваша роль - " + getRoleName(role))
		} else {
			helpmessage := "Доступные команды для Вашей роли (" + getRoleName(role) + ")\n"
			switch role {
			case ROLE_ADMIN:
				helpmessage += "/ban <ник> [причина] - забанить игрока\n"
				fallthrough //+права модера
			case ROLE_MODERATOR:
				helpmessage += "/kick <ник> - кикнуть игрока\n"
			}
			sendTo(peerid, helpmessage)
		}
	case "роль", "role", "мояроль":
		sendTo(peerid, "Ваша роль - " + getRoleName(getUserRole(userid)))
	case "ban", "бан":
		if role := getUserRole(userid); role < ROLE_ADMIN { //можно выполнять от ROLE_ADMIN (3)
			sendTo(peerid, "🚫 Недостаточно прав для выполнения этой команды. Требуется роль " + getRoleName(ROLE_ADMIN) + ", Ваша роль - " + getRoleName(role))
		} else {
			if !cmd.IssetArg(0) {
				sendTo(peerid, "Вы не указали ник первым аргументом")
			} else {
				reason := "не указана причина"
				nick := cmd.GetArgs()[0]
				if cmd.IssetArg(1) {
					reason = strings.Join(cmd.GetArgs()[0:], " ")
				}
				if response, err := SendRCONCommand("ban " + nick + " " + reason); err != nil {
					sendTo(peerid, "Не удалось выполнить команду: " + err.Error())
				} else {
					sendTo(peerid, "Команда успешно выполнена, сервер ответил: " + response)
				}

			}

		}
	case "kick", "кик":
		if role := getUserRole(userid); role < ROLE_MODERATOR { //можно выполнять от ROLE_MODERATOR (2)
			sendTo(peerid, "🚫 Недостаточно прав для выполнения этой команды. Требуется роль " + getRoleName(ROLE_MODERATOR) + ", Ваша роль - " + getRoleName(role))
		} else {
			if !cmd.IssetArg(0) {
				sendTo(peerid, "Вы не указали ник первым аргументом")
			} else {
				nick := cmd.GetArgs()[0]
				if response, err := SendRCONCommand("kick " + nick); err != nil {
					sendTo(peerid, "Не удалось выполнить команду: " + err.Error())
				} else {
					sendTo(peerid, "Команда успешно выполнена, сервер ответил: " + response)
				}

			}

		}

	}
}

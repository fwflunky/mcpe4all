package main

/*
	# RCON бот для ВКонтакте с базовым набором функций
	# для взаимодействия с сервером Minecraft: PE/BE
	# Автор - vk.com/lywulf ; github.com/fwflunky
*/

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
)

const (
	GroupID    = 111111111 // Айди группы ВК
	GroupToken = "токен" // Токен группы ВК
)

var VKApi *api.VK

func main() {
	VKApi = api.NewVK(GroupToken)

	lp, err := longpoll.NewLongPoll(VKApi, GroupID) //подключение к лонгполлу группы ВК
	if err != nil {
		log.Fatalln("Ошибка при подключении к лонгполлу ВК:", err)
	}

	lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		userid, peerid, message := obj.Message.FromID, obj.Message.PeerID, obj.Message.Text //айди юзера, айди диалога, сообщение
		handleMessage(obj, userid, peerid, message)
	})
	log.Println("Бот запущен")
	log.Fatalln(lp.Run())
}

func sendTo(peerid int, message string){
	_, _ = VKApi.MessagesSend(api.Params{ //чтобы небыло варнинга в иде
		"peer_id":    peerid,
		"random_id":  0,
		"message":    message,
		"dont_parse_links": 1,
	})
}
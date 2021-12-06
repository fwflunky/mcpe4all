package main

/*
	# RCON - подключение
	# Настройка - вписать айпи сервера в SERVER_IP = "АЙПИ:ПОРТ"
	# RCON_PASS = "ПАРОЛЬ РКОНА"
*/

import (
	mcpercon "github.com/fwflunky/mcpe4all/mcpe-rcon"
	"log"
)

const (
	SERVER_IP = "192.168.0.103:19132"
	RCON_PASS = "testtest"
)

var RconConnection *mcpercon.MCConn
func SendRCONCommand(cmd string) (string, error) {
	return RconConnection.SendCommand(cmd)
}

func init(){
	RconConnection = new(mcpercon.MCConn)
	if err := RconConnection.Open(SERVER_IP, RCON_PASS); err != nil {
		log.Fatalln("Не удалось подключится к RCON:", err)
	}

	if err := RconConnection.Authenticate(); err != nil {
		log.Fatalln("Не удалось авторизироваться в RCON:", err)
	}
	log.Println("Успешно подключен к RCON")
}

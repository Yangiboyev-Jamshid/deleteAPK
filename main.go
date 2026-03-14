package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("AUTO_DELETE_APK_BOT")
	if token == "" {
		log.Fatalln("Token topilmadi!!!")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Bot yaratishda xatolik!!! %v", err)
	}

	bot.Debug = false
	log.Printf("Bot ishga tushdi: @%s", &bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 40

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := update.Message

		if msg.Chat.Type != "group" && msg.Chat.Type != "supergroup" {
			continue
		}

		if msg.Document == nil {
			continue
		}

		fileName := msg.Document.FileName

		if isAPKFile(fileName) {
			deleteConfig := tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID)

			_, err := bot.Request(deleteConfig)
			if err != nil {
				log.Panicf("Xabarni o'chirishda xatolik (chat: %d, msg: %d): %v",
					msg.Chat.ID, msg.MessageID, err,
				)
				continue
			}

			log.Panicf("APK file ochirildi: '%s' | Guruh %s | Yuboruvchi: @%s",
				fileName, msg.Chat.Title, msg.From.UserName,
			)
		}
	}
}

func isAPKFile(fileName string) bool {
	return strings.HasSuffix(strings.ToLower(fileName), ".apk")
}

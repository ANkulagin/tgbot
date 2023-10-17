package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из файла .env
	godotenv.Load(`C:\Project\GolangProject\tg_bot\.env`)

	// Получаем токен из переменных окружения
	token := os.Getenv("TOKEN")

	// Инициализируем нового бота с использованием токена
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Включаем режим отладки
	bot.Debug = true

	// Выводим информацию о том, что бот успешно авторизован
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем канал обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем канал для обновлений
	updates := bot.GetUpdatesChan(u)

	// Обрабатываем полученные обновления
	for update := range updates {
		// Проверяем, является ли сообщение не пустым
		if update.Message != nil {

			// Проверяем, была ли введена команда "/help"
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			default:
				// Если не была введена команда "/help", выполняем стандартное поведение
				defaultBehavior(bot, update.Message)
			}

			// Если введена команда "/help", переходим к следующему обновлению
			if update.Message.Command() == "help" {
				continue
			}

			// Выводим информацию о полученном сообщении в лог
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Отправляем ответное сообщение с тем же текстом
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

// Функция обработки команды "/help"
func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	// Создаем новое сообщение с информацией о команде "/help"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")

	// Отправляем сообщение
	bot.Send(msg)
}

// Функция стандартного поведения
func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	// Выводим информацию о полученном сообщении в лог
	log.Printf("[%s],%s", inputMessage.From.UserName, inputMessage.Text)

	// Отправляем ответное сообщение с тем же текстом
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	bot.Send(msg)
}

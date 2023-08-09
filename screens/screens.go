package screens

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

var screens = make(map[string]Screen)
var currentScreen string = "null"
var StartScreen string = "main_menu"
var Language string = "eng"

type Screen interface {
	GetPrev() string
	setPrev(prev string)
	GetText() string
	ReadInput(ctx context.Context, b *bot.Bot, update *models.Update)
	GetKeyboard(b *bot.Bot) *inline.Keyboard
}

func onInlineKeyboard(ctx context.Context, b *bot.Bot, mes *models.Message, data []byte) {
	screen := screens[string(data)]
	screen.setPrev(currentScreen)
	currentScreen = string(data)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      mes.Chat.ID,
		Text:        screen.GetText(),
		ReplyMarkup: screen.GetKeyboard(b),
	})
}

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if currentScreen == "null" {
		currentScreen = StartScreen
		startScreen := screens[currentScreen]

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        startScreen.GetText(),
			ReplyMarkup: startScreen.GetKeyboard(b),
		})
	}
	screens[currentScreen].ReadInput(ctx, b, update)
}

func deleteOnInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.ID,
	})
}

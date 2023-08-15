package screens

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type MenuScreen struct {
	idServies
	prev         string
	text         *map[string]string
	keyboardTo   *[][]string
	keyboardText *map[string][][]string
}

func NewMenuScreen(address string, text *map[string]string, keyboardTo *[][]string, keyboardText *map[string][][]string) *MenuScreen {
	ms := MenuScreen{
		prev:         StartScreen,
		text:         text,
		keyboardTo:   keyboardTo,
		keyboardText: keyboardText,
	}
	screens[address] = &ms
	return &ms
}

func (ms *MenuScreen) GetPrev() string {
	return ms.prev
}

func (ms *MenuScreen) setPrev(prev string) {
	ms.prev = prev
}

func (ms *MenuScreen) GetText() string {
	return (*ms.text)[Language]
}

func (ms *MenuScreen) ReadInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	deleteOnInput(ctx, b, update)
}

func (ms *MenuScreen) GetKeyboard(b *bot.Bot) *inline.Keyboard {
	kb := inline.New(b)
	for i := 0; i < len(*ms.keyboardTo); i++ {
		kb.Row()
		for j := 0; j < len((*ms.keyboardTo)[i]); j++ {
			if (*ms.keyboardTo)[i][j] == "prev" {
				kb.Button((*ms.keyboardText)[Language][i][j], []byte(((*ms).GetPrev())), onInlineKeyboard)
			} else {
				kb.Button((*ms.keyboardText)[Language][i][j], []byte((*ms.keyboardTo)[i][j]), onInlineKeyboard)
			}
		}
	}
	return kb
}

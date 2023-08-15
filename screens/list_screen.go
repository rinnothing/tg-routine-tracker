package screens

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type buttonData struct {
	text    string
	goal    string
	context string
}

type getListFunc func() *[]buttonData

type ListScreen struct {
	idServies

	address string
	prev    string
	text    *map[string]string

	topButtonText *map[string]string
	topButtonTo   string

	getList getListFunc
	list    *[]buttonData
	context string

	bottomButtonText *map[string]string
	bottomButtonTo   string
}

func NewListScreen(address string, text *map[string]string, topButtonTo string, topButtonText *map[string]string, bottomButtonTo string, bottomButtonText *map[string]string, getList getListFunc) *ListScreen {
	ls := ListScreen{
		address:          address,
		prev:             StartScreen,
		topButtonText:    topButtonText,
		topButtonTo:      topButtonTo,
		text:             text,
		getList:          getList,
		bottomButtonText: bottomButtonText,
		bottomButtonTo:   bottomButtonTo,
	}
	screens[address] = &ls
	return &ls
}

func (ls *ListScreen) GetPrev() string {
	return ls.prev
}

func (ls *ListScreen) setPrev(prev string) {
	ls.prev = prev
}

func (ls *ListScreen) GetText() string {
	return (*ls.text)[Language]
}

func (ls *ListScreen) ReadInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	deleteOnInput(ctx, b, update)
}

func (ls *ListScreen) GetKeyboard(b *bot.Bot) *inline.Keyboard {
	ls.list = ls.getList()

	kb := inline.New(b)
	if ls.topButtonTo == "prev" {
		kb.Row().Button((*ls.topButtonText)[Language], []byte(ls.prev), onInlineKeyboard)
	} else {
		kb.Row().Button((*ls.topButtonText)[Language], []byte(ls.topButtonTo), onInlineKeyboard)
	}

	for i := 0; i < len(*ls.list); i++ {
		kb.Row().Button((*ls.list)[i].text, []byte(fmt.Sprintln(ls.address, (*ls.list)[i].context, (*ls.list)[i].goal)), onInlineKeyboardList)
	}
	if ls.bottomButtonTo == "prev" {
		kb.Row().Button((*ls.bottomButtonText)[Language], []byte(ls.prev), onInlineKeyboard)
	} else {
		kb.Row().Button((*ls.bottomButtonText)[Language], []byte(ls.bottomButtonTo), onInlineKeyboard)
	}

	return kb
}

func (ls *ListScreen) GetContext() []string {
	return []string{ls.context}
}

func onInlineKeyboardList(ctx context.Context, b *bot.Bot, mes *models.Message, data []byte) {
	words := strings.Fields(string(data))
	screens[words[0]].(*ListScreen).context = words[1]
	onInlineKeyboard(ctx, b, mes, []byte(words[2]))
}

func MakeDataDummy(length int, goal string) func() *[]buttonData {
	return func() *[]buttonData {
		dl := make([]buttonData, length)
		for i := 0; i < length; i++ {
			dl[i] = buttonData{text: fmt.Sprint(i, " String"), goal: goal, context: fmt.Sprint(i)}
		}
		return &dl
	}
}

package screens

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type buttonData struct {
	text string
	to   string
}

type getListFunc func() *[]buttonData

type ListScreen struct {
	address string
	goal    string
	prev    string
	text    *map[string]string

	getList getListFunc
	list    *[]buttonData
	option  int

	lowerButtonText *map[string]string
	lowerButtonTo   string
}

func NewListScreen(address string, goal string, text *map[string]string, lowerButtonTo string, lowerButtonText *map[string]string, getList getListFunc) *ListScreen {
	ls := ListScreen{
		address:         address,
		goal:            goal,
		prev:            StartScreen,
		text:            text,
		getList:         getList,
		lowerButtonText: lowerButtonText,
		lowerButtonTo:   lowerButtonTo,
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
	for i := 0; i < len(*ls.list); i++ {
		kb.Row().Button((*ls.list)[i].text, []byte(fmt.Sprintln(ls.address, i, ls.goal)), onInlineKeyboardList)
	}
	kb.Row().Button((*ls.lowerButtonText)[Language], []byte(ls.lowerButtonTo), onInlineKeyboardList)

	return inline.New(b)
}

func onInlineKeyboardList(ctx context.Context, b *bot.Bot, mes *models.Message, data []byte) {
	words := strings.Fields(string(data))
	screens[words[0]].(*ListScreen).option, _ = strconv.Atoi(words[1])
	onInlineKeyboard(ctx, b, mes, []byte(words[2]))
}

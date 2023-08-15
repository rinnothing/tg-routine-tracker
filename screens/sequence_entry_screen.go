package screens

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type ContextcheckerFunc func(context string) bool

type doAfterFunc func(ctx context.Context, b *bot.Bot, update *models.Update, ses *SequenceEntryScreen)

type SequenceEntryScreen struct {
	idServies
	prev           string
	text           *map[string]string
	option         int
	options        *map[string][]string
	errorText      *map[string][]string
	context        []string
	contextChecker *[]ContextcheckerFunc
	doAfter        doAfterFunc
}

func NewSequenceEntryScreen(address string, text *map[string]string, options *map[string][]string, errorText *map[string][]string, contextChecker *[]ContextcheckerFunc, doAfter doAfterFunc) *SequenceEntryScreen {
	ses := SequenceEntryScreen{
		idServies:      idServies{},
		prev:           StartScreen,
		text:           text,
		option:         0,
		options:        options,
		errorText:      errorText,
		context:        make([]string, len((*options)[Language])),
		contextChecker: contextChecker,
		doAfter:        doAfter,
	}
	screens[address] = &ses
	return &ses
}

func (ses *SequenceEntryScreen) GetPrev() string {
	return ses.prev
}

func (ses *SequenceEntryScreen) setPrev(prev string) {
	ses.prev = prev
}

func (ses *SequenceEntryScreen) GetText() string {
	return fmt.Sprint((*ses.text)[Language], "\n", (*ses.options)[Language][ses.option])
}

func (ses *SequenceEntryScreen) ReadInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := update.Message.Text
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.ID,
	})
	if (*ses.contextChecker)[ses.option](text) {
		(ses.context)[ses.option] = text
		ses.option++
		if ses.option < len((*ses.options)[Language]) {
			b.EditMessageText(ctx, &bot.EditMessageTextParams{
				ChatID:    ses.GetChatID(),
				MessageID: ses.GetMessageID(),
				Text:      ses.GetText(),
			})
		} else {
			b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    ses.GetChatID(),
				MessageID: ses.GetMessageID(),
			})
			ses.doAfter(ctx, b, update, ses)
		}
	} else {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    ses.GetChatID(),
			MessageID: ses.GetMessageID(),
			Text:      fmt.Sprint(ses.GetText(), '\n', (*ses.errorText)[Language][ses.option]),
		})
	}
}

func (ses *SequenceEntryScreen) GetKeyboard(b *bot.Bot) *inline.Keyboard {
	return inline.New(b)
}

func (ses *SequenceEntryScreen) GetContext() []string {
	return ses.context
}

func GoPrevAfter(ctx context.Context, b *bot.Bot, update *models.Update, ses *SequenceEntryScreen) {
	onInlineKeyboard(ctx, b, update.Message, []byte(ses.GetPrev()))
}

func DummyChecker(context string) bool {
	return true
}

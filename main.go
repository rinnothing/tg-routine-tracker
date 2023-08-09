package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"tg-routine-tracker/screens"

	"github.com/go-telegram/bot"
)

func main() {
	screens.StartScreen = "main_menu"
	screens.Language = "rus"

	dat, err := os.ReadFile("token")
	if err != nil {
		fmt.Print("Error reading token, check if you have \"token\" file containing your token.")
	}
	token := string(dat)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(screens.DefaultHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	screens.NewMenuScreen("main_menu", &map[string]string{"rus": "Главное меню"},
		&[][]string{{"my_routines", "coming_reminders"},
			{"common_user_stats", "settings"}},
		&map[string][][]string{"rus": {{"Мои привычки", "Ближайшие напоминания"},
			{"Моя статистика", "Настройки"}}})

	screens.NewMenuScreen("my_routines", &map[string]string{"rus": "Мои привычки"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	// newRoutine := screen{
	// 	text: "Новая привычка",
	// 	makeKeyboard: func(b *bot.Bot, s *screen) *inline.Keyboard {
	// 		return inline.New(b).
	// 			Row().
	// 			Button("Назад", []byte(s.prev), onInlineKeyboard)
	// 	},
	// 	update: func(s *screen) {},
	// 	input:  prevRemove,
	// }
	// routineInList := screen{
	// 	text: "",
	// 	makeKeyboard: func(b *bot.Bot, s *screen) *inline.Keyboard {
	// 		return inline.New(b).
	// 			Row().
	// 			Button("Назад", []byte(s.prev), onInlineKeyboard)
	// 	},
	// 	update: func(s *screen) {},
	// 	input:  prevRemove,
	// }
	screens.NewMenuScreen("coming_reminders", &map[string]string{"rus": "Ближайшие уведомления"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})
	// userRoutineStats := screen{
	// 	text: "",
	// 	makeKeyboard: func(b *bot.Bot, s *screen) *inline.Keyboard {
	// 		return inline.New(b).
	// 			Row().
	// 			Button("Назад", []byte(s.prev), onInlineKeyboard)
	// 	},
	// 	update: func(s *screen) {},
	// 	input:  prevRemove,
	// }
	// commonRoutineStats := screen{
	// 	text: "",
	// 	makeKeyboard: func(b *bot.Bot, s *screen) *inline.Keyboard {
	// 		return inline.New(b).
	// 			Row().
	// 			Button("Назад", []byte(s.prev), onInlineKeyboard)
	// 	},
	// 	update: func(s *screen) {},
	// 	input:  prevRemove,
	// }
	screens.NewMenuScreen("common_user_stats", &map[string]string{"rus": "Статистика пользователя"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	screens.NewMenuScreen("settings", &map[string]string{"rus": "Настройки"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	b.Start(ctx)
}
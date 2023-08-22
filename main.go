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

	screens.NewListScreen("my_routines", &map[string]string{"rus": "Мои привычки"},
		"new_routine", &map[string]string{"rus": "Новая привычка"},
		"prev", &map[string]string{"rus": "Назад"},
		screens.MakeDataDummy(1, "routine"))

	screens.NewCustomTextScreen("routine", screens.DummyContextReader,
		&[][]string{{"routine_stats", "change_time"},
			{"change_name", "change_description"},
			{"prev"}},
		&map[string][][]string{"rus": {{"Статистика", "Изменить время"},
			{"Изменить имя", "Изменить описание"},
			{"Назад"}}})

	screens.NewSequenceEntryScreen("new_routine", &map[string]string{"rus": "Новая привычка"},
		&map[string][]string{"rus": {"Название", "Время", "Частота", "Описание"}},
		&map[string][]string{"rus": {"Проверьте правильность названия", "Проверьте правильность времени", "Проверьте правильность частоты", "Проверьте правильность описания"}},
		&[]screens.ContextcheckerFunc{screens.DummyChecker, screens.DummyChecker, screens.DummyChecker, screens.DummyChecker},
		screens.GoPrevAfter)

	screens.NewMenuScreen("coming_reminders", &map[string]string{"rus": "Ближайшие уведомления"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	screens.NewMenuScreen("routine_stats", &map[string]string{"rus": "Статистика привычки"},
		&[][]string{{"common_routine_stats"},
			{"prev"}},
		&map[string][][]string{"rus": {{"Общая статистика"},
			{"Назад"}}})

	screens.NewMenuScreen("common_routine_stats", &map[string]string{"rus": "Общая статистика привычки"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	screens.NewMenuScreen("common_user_stats", &map[string]string{"rus": "Статистика пользователя"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	screens.NewSequenceEntryScreen("change_time", &map[string]string{"rus": "Введите новое время"},
		&map[string][]string{"rus": {""}}, &map[string][]string{"rus": {"Проверьте правильность времени"}},
		&[]screens.ContextcheckerFunc{screens.DummyChecker}, screens.GoPrevAfter)

	screens.NewSequenceEntryScreen("change_name", &map[string]string{"rus": "Введите новое имя"},
		&map[string][]string{"rus": {""}}, &map[string][]string{"rus": {"Проверьте правильность имени"}},
		&[]screens.ContextcheckerFunc{screens.DummyChecker}, screens.GoPrevAfter)

	screens.NewSequenceEntryScreen("change_description", &map[string]string{"rus": "Введите новое описание"},
		&map[string][]string{"rus": {""}}, &map[string][]string{"rus": {"Проверьте правильность описания"}},
		&[]screens.ContextcheckerFunc{screens.DummyChecker}, screens.GoPrevAfter)

	screens.NewMenuScreen("settings", &map[string]string{"rus": "Настройки"},
		&[][]string{{"prev"}},
		&map[string][][]string{"rus": {{"Назад"}}})

	b.Start(ctx)
}

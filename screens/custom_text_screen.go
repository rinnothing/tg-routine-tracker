package screens

import "fmt"

type getTextFunc func(ct *CustomTextScreen) *map[string]string

type CustomTextScreen struct {
	MenuScreen
	getText getTextFunc
}

func NewCustomTextScreen(address string, getText getTextFunc, keyboardTo *[][]string, keyboardText *map[string][][]string) *CustomTextScreen {
	ct := CustomTextScreen{
		MenuScreen: MenuScreen{prev: StartScreen, keyboardTo: keyboardTo, keyboardText: keyboardText},
		getText:    getText,
	}
	screens[address] = &ct
	return &ct
}

func (ct *CustomTextScreen) GetText() string {
	return (*ct.getText(ct))[Language]
}

func DummyContextReader(ct *CustomTextScreen) *map[string]string {
	ctx := screens[ct.prev].(ScreenContext).GetContext()
	return &map[string]string{"rus": fmt.Sprint("Вы выбрали " + ctx[0] + " вариант.")}
}

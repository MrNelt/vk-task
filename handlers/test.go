package handlers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

type HandlerTest struct {
	keyboardTest  *tele.ReplyMarkup
	keyboardStart *tele.ReplyMarkup
}

func NewHandlerTest(keyboardTest, keyboardStart *tele.ReplyMarkup) *HandlerTest {
	return &HandlerTest{keyboardTest: keyboardTest, keyboardStart: keyboardStart}
}

func (h *HandlerTest) Test(ctx tele.Context) error {
	if ctx.Text() == "Да" {
		message := fmt.Sprintf("В каком году была основана социальная сеть ВКонтакте?\n\n1) 2004\n2) 2005\n3) 2006\n4) 2010\n")
		return ctx.Send(message, h.keyboardTest)
	}
	message := fmt.Sprintf("Хорошо, если передумаешь, /start")
	return ctx.Send(message, h.keyboardStart)
}

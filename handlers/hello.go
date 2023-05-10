package handlers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

type HandlerHello struct {
	keyboard *tele.ReplyMarkup
}

func NewHandlerHello(keyboard *tele.ReplyMarkup) *HandlerHello {
	return &HandlerHello{keyboard: keyboard}
}

func (h *HandlerHello) Hello(ctx tele.Context) error {
	message := fmt.Sprintf("Привет, %s, хочешь ответить на вопрос?", ctx.Sender().FirstName)
	return ctx.Send(message, h.keyboard)
}

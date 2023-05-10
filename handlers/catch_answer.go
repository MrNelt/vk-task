package handlers

import (
	tele "gopkg.in/telebot.v3"
)

type HandlerCatchAnswer struct {
	keyboard *tele.ReplyMarkup
	answer   string
}

func NewHandlerCatchAnswer(keyboard *tele.ReplyMarkup, answer string) *HandlerCatchAnswer {
	return &HandlerCatchAnswer{answer: answer, keyboard: keyboard}
}

func (h *HandlerCatchAnswer) CatchRight(ctx tele.Context) error {
	message := ""
	message += "✅ Верно!"
	message += "\n\n"
	message += "Начать заново /start"
	return ctx.Send(message, h.keyboard)
}

func (h *HandlerCatchAnswer) CatchWrong(ctx tele.Context) error {
	message := ""
	message += "❌ Неверно!"
	message += "\n\n"
	message += "Начать заново /start"
	return ctx.Send(message, h.keyboard)
}

package keyboards

import tele "gopkg.in/telebot.v3"

func NewKeyboardStart() *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := keyboard.Text("/start")
	keyboard.Reply(
		keyboard.Row(btn1),
	)
	return keyboard
}

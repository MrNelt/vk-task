package keyboards

import tele "gopkg.in/telebot.v3"

func NewKeyboardVariants() (*tele.ReplyMarkup, *tele.Btn) {
	keyboard := &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := keyboard.Text("1")
	btn2 := keyboard.Text("2")
	btn3 := keyboard.Text("3")
	btn4 := keyboard.Text("4")
	keyboard.Reply(
		keyboard.Row(btn1, btn2),
		keyboard.Row(btn3, btn4),
	)
	return keyboard, &btn3
}

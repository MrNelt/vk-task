package keyboards

import tele "gopkg.in/telebot.v3"

func NewKeyboardYesNo() (*tele.ReplyMarkup, *tele.Btn, *tele.Btn) {
	keyboard := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnYes := keyboard.Text("Да")
	btnNo := keyboard.Text("Нет")
	keyboard.Reply(
		keyboard.Row(btnYes, btnNo),
	)
	return keyboard, &btnYes, &btnNo
}

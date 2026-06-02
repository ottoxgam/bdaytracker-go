package tgview

import (
	"fmt"

	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/bdaytracker-go/internal/usersession"

	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
)

type Menu struct {
}

func (m Menu) Send(s *usersession.Session, edit bool) {
	text := fmt.Sprintf(`<b>%s</b>

<b>%s</b> — add a new friend manually.

<b>%s</b> — link a Telegram friend by @username.

<b>%s</b> — view your friends list.`, btn.Menu, btn.AddFriend, btn.AddFromTelegram, btn.FriendList)
	keyboard := [][]telegram.InlineKeyboardButton{
		{
			tgcallback.Button(btn.AddFriend, tgcallback.AddFriend{}),
			tgcallback.Button(btn.AddFromTelegram, tgcallback.AddFromTelegram{}),
		},
		{
			tgcallback.Button(btn.FriendList, tgcallback.FriendList{}),
		},
	}

	s.SendEditText(text, keyboard, edit)
}

func (m Menu) Keyboard() [][]telegram.KeyboardButton {
	return [][]telegram.KeyboardButton{
		{
			{
				Text: btn.Menu,
			},
		},
	}
}

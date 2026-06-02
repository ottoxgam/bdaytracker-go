package tgview

import (
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
	"github.com/lodthe/bdaytracker-go/internal/usersession"
)

type RemoveFromTelegram struct {
}

func (f RemoveFromTelegram) Success(s *usersession.Session) {
	_ = s.SendEditText("✅ Telegram friends have been removed!", [][]telegram.InlineKeyboardButton{
		{
			tgcallback.Button(btn.FriendList, tgcallback.FriendList{}),
		},
		{
			tgcallback.Button(btn.Menu, tgcallback.OpenMenu{}),
		},
	}, true)
}

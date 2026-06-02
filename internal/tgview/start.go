package tgview

import (
	"github.com/lodthe/bdaytracker-go/internal/usersession"
)

type Start struct {
}

func (st Start) Send(s *usersession.Session) {
	s.SendInlinePhoto(`Hi! I can remind you about your friends' 🎁 Birthdays.

You can add them manually or link a Telegram friend.

When someone's Birthday arrives, I'll remind you!`, "greetings.png", nil)
}

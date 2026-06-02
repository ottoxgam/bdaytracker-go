package tgview

import (
	"fmt"

	"github.com/petuhovskiy/telegram"

	friendship "github.com/lodthe/bdaytracker-go/internal/friendship"
	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
	"github.com/lodthe/bdaytracker-go/internal/usersession"
)

type AddFromTelegram struct {
}

func (f AddFromTelegram) AskForUsername(s *usersession.Session) {
	f.send(s, "Send your friend's Telegram @username, e.g. <code>@johndoe</code>.")
}

func (f AddFromTelegram) AskForDate(s *usersession.Session) {
	f.send(s, "Now send their birthday as <code>MM.DD</code> or <code>MM/DD</code>\n\nFor example, <code>07.09</code> or <code>07/09</code> means July 9th.")
}

func (f AddFromTelegram) FailedToParseDate(s *usersession.Session) {
	f.send(s, "I couldn't understand that 😔\nPlease use <code>MM.DD</code> or <code>MM/DD</code> format.\n\nTry again! 😉")
}

func (f AddFromTelegram) WrongNumberOfDays(s *usersession.Session) {
	f.send(s, "❌ That month can't have that many days. Try again! 😉")
}

func (f AddFromTelegram) InvalidUsername(s *usersession.Session) {
	f.send(s, "That doesn't look like a valid username. Try again, e.g. <code>@johndoe</code>.")
}

func (f AddFromTelegram) Cancel(s *usersession.Session) {
	_ = s.SendText("Cancelled.", Menu{}.Keyboard())
}

func (f AddFromTelegram) Success(s *usersession.Session, friend friendship.Friend) {
	keyboard := [][]telegram.InlineKeyboardButton{
		{
			tgcallback.Button(btn.AddFromTelegram, tgcallback.AddFromTelegram{}),
		},
		{
			tgcallback.Button(btn.FriendList, tgcallback.FriendList{}),
			tgcallback.Button(btn.Menu, tgcallback.OpenMenu{}),
		},
	}

	_ = s.SendText("👥", menuKeyboard())
	_ = s.SendText(fmt.Sprintf("<a href=\"https://t.me/%s\">%s</a> has been added to your friends list!", friend.Name, friend.Name), keyboard)
}

func (f AddFromTelegram) send(s *usersession.Session, text string) {
	_ = s.SendText(text, cancelKeyboard())
}

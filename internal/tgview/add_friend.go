package tgview

import (
	"github.com/petuhovskiy/telegram"

	friendship "github.com/lodthe/bdaytracker-go/internal/friendship"
	"github.com/lodthe/bdaytracker-go/internal/usersession"

	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
)

type AddFriend struct {
}

func (a AddFriend) AskName(s *usersession.Session) {
	a.send(s, "Send the name of your new friend 🧑‍\U0001F9B0.")
}

func (a AddFriend) AskDate(s *usersession.Session) {
	a.send(s, `Send their birthday in the following format:

<code>DD.MM</code>

For example, 09.07 means July 9th.
`)
}

func (a AddFriend) FailedToParseDate(s *usersession.Session) {
	a.send(s, `I couldn't understand that 😔
The message must match the following format:
<code>DD.MM</code>

Try again! 😉`)
}

func (a AddFriend) WrongNumberOfDays(s *usersession.Session) {
	a.send(s, `❌ That month can't have that many days. Try again! 😉`)
}

func (a AddFriend) Cancel(s *usersession.Session) {
	_ = s.SendText(`Cancelled.

Maybe next time.`, Menu{}.Keyboard())
}

func (a AddFriend) Success(s *usersession.Session, newFriend friendship.Friend) {
	keyboard := [][]telegram.InlineKeyboardButton{
		{
			tgcallback.Button(btn.AddFriend, tgcallback.AddFriend{}),
		},
		{
			tgcallback.Button(btn.FriendList, tgcallback.FriendList{}),
			tgcallback.Button(btn.Menu, tgcallback.OpenMenu{}),
		},
	}

	_ = s.SendText("👥", menuKeyboard())
	_ = s.SendText("<code>"+newFriend.Name+"</code> has been added to your friends list!", keyboard)
}

func (a AddFriend) send(s *usersession.Session, text string) {
	_ = s.SendText(text, [][]telegram.KeyboardButton{{
		{
			Text: btn.Cancel,
		},
	}})
}

package tgview

import (
	"fmt"

	"github.com/petuhovskiy/telegram"

	friendship "github.com/lodthe/bdaytracker-go/internal/friendship"
	"github.com/lodthe/bdaytracker-go/internal/usersession"

	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
)

type RemoveFriend struct {
}

func (f RemoveFriend) AskIndexOrName(s *usersession.Session) {
	s.SendText("Send the friend's full name or their number from the friends list.", cancelKeyboard())
}

func (f RemoveFriend) WrongIndex(s *usersession.Session) {
	s.SendText("Invalid number. Try again!", cancelKeyboard())
}

func (f RemoveFriend) WrongName(s *usersession.Session) {
	s.SendText("Can't find a friend with that name. The name must match the one in your friends list.\n\nTry again!", cancelKeyboard())
}

func (f RemoveFriend) AskForApprove(s *usersession.Session, friend friendship.Friend) {
	text := fmt.Sprintf("The following entry will be removed from your friends list:\n%s", formatFriend(friend))
	s.SendText(text, [][]telegram.InlineKeyboardButton{
		{
			tgcallback.Button(btn.Approve, tgcallback.RemoveFriendApprove{
				UUID: friend.UUID,
			}),
			tgcallback.Button(btn.Cancel, tgcallback.RemoveFriendCancel{
				UUID: friend.UUID,
			}),
		},
	})
}

func (f RemoveFriend) Approved(s *usersession.Session, _ tgcallback.RemoveFriendApprove) {
	_ = s.DeleteLastMessage()
	_ = s.SendText("Friends list updated.", menuKeyboard())
}

func (f RemoveFriend) Canceled(s *usersession.Session, _ tgcallback.RemoveFriendCancel) {
	_ = s.DeleteLastMessage()
	_ = s.SendText("Got it, nobody's being removed!", menuKeyboard())
}

func (f RemoveFriend) Cancel(s *usersession.Session) {
	_ = s.SendText("Cancelled.", Menu{}.Keyboard())
}

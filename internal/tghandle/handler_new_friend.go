package tghandle

import (
	"github.com/google/uuid"

	friendship "github.com/lodthe/bdaytracker-go/internal/friendship"
	"github.com/lodthe/bdaytracker-go/internal/usersession"

	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgstate"
	"github.com/lodthe/bdaytracker-go/internal/tgview"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
)

type AddFriendHandler struct {
}

func (h *AddFriendHandler) State() tgstate.ID {
	return tgstate.AddFriend
}

func (h *AddFriendHandler) Callback() interface{} {
	return tgcallback.AddFriend{}
}

func (h *AddFriendHandler) HandleCallback(s *usersession.Session, clb interface{}) {
	s.State.State = tgstate.AddFriend
	s.State.NewFriend = friendship.Friend{}
	tgview.AddFriend{}.AskName(s)
}

func (h *AddFriendHandler) HandleMessage(s *usersession.Session, msgText string) {
	switch {
	case msgText == btn.Cancel:
		h.cancel(s)

	case s.State.NewFriend.Name == "":
		h.handleName(s, msgText)

	case s.State.NewFriend.BDay == 0:
		h.handleDate(s, msgText)
	}
}

func (h *AddFriendHandler) cancel(s *usersession.Session) {
	s.State.State = tgstate.None
	tgview.AddFriend{}.Cancel(s)
}

func (h *AddFriendHandler) handleName(s *usersession.Session, msgText string) {
	s.State.NewFriend.Name = msgText
	tgview.AddFriend{}.AskDate(s)
}

func (h *AddFriendHandler) handleDate(s *usersession.Session, msgText string) {
	month, day, result := parseBirthday(msgText)
	switch result {
	case dateParseInvalid:
		tgview.AddFriend{}.FailedToParseDate(s)
		return
	case dateParseWrongDays:
		tgview.AddFriend{}.WrongNumberOfDays(s)
		return
	}

	friend := s.State.NewFriend
	friend.BMonth = month
	friend.BDay = day
	friend.UUID = uuid.New().String()
	s.State.Friends = append(s.State.Friends, friend)
	s.State.State = tgstate.None
	s.State.NewFriend = friendship.Friend{}

	tgview.AddFriend{}.Success(s, friend)
}

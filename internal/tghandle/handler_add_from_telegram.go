package tghandle

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	friendship "github.com/lodthe/bdaytracker-go/internal/friendship"
	"github.com/lodthe/bdaytracker-go/internal/tgcallback"
	"github.com/lodthe/bdaytracker-go/internal/tgstate"
	"github.com/lodthe/bdaytracker-go/internal/tgview"
	"github.com/lodthe/bdaytracker-go/internal/tgview/btn"
	"github.com/lodthe/bdaytracker-go/internal/usersession"
)

type AddFromTelegramHandler struct {
}

func (h *AddFromTelegramHandler) State() tgstate.ID {
	return tgstate.ImportFromTelegram
}

func (h *AddFromTelegramHandler) Callback() interface{} {
	return tgcallback.AddFromTelegram{}
}

func (h *AddFromTelegramHandler) HandleCallback(s *usersession.Session, clb interface{}) {
	s.State.State = tgstate.ImportFromTelegram
	s.State.NewFriend = friendship.Friend{}
	tgview.AddFromTelegram{}.AskForUsername(s)
}

func (h *AddFromTelegramHandler) HandleMessage(s *usersession.Session, msgText string) {
	if msgText == btn.Cancel {
		tgview.AddFromTelegram{}.Cancel(s)
		s.State.State = tgstate.None
		return
	}

	if s.State.NewFriend.Name == "" {
		h.handleUsername(s, msgText)
	} else {
		h.handleDate(s, msgText)
	}
}

func (h *AddFromTelegramHandler) handleUsername(s *usersession.Session, raw string) {
	username := strings.TrimPrefix(raw, "@")
	username = strings.TrimSpace(username)

	if username == "" || strings.ContainsAny(username, " \t\n") {
		tgview.AddFromTelegram{}.InvalidUsername(s)
		return
	}

	s.State.NewFriend.Name = username
	tgview.AddFromTelegram{}.AskForDate(s)
}

func (h *AddFromTelegramHandler) handleDate(s *usersession.Session, msgText string) {
	const numberOfMonths = 12
	var daysBefore = [...]int{
		0,
		31,
		31 + 28,
		31 + 28 + 31,
		31 + 28 + 31 + 30,
		31 + 28 + 31 + 30 + 31,
		31 + 28 + 31 + 30 + 31 + 30,
		31 + 28 + 31 + 30 + 31 + 30 + 31,
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
	}

	friend := s.State.NewFriend
	_, err := fmt.Sscanf(msgText, "%d.%d", &friend.BDay, &friend.BMonth)
	if err != nil {
		tgview.AddFromTelegram{}.FailedToParseDate(s)
		return
	}

	if friend.BMonth < 1 || friend.BMonth > numberOfMonths {
		tgview.AddFromTelegram{}.FailedToParseDate(s)
		return
	}

	daysInMonth := daysBefore[friend.BMonth] - daysBefore[friend.BMonth-1]
	if friend.BMonth == int(time.February) {
		daysInMonth++
	}
	if friend.BDay < 0 || friend.BDay > daysInMonth {
		tgview.AddFromTelegram{}.WrongNumberOfDays(s)
		return
	}

	username := friend.Name
	friend.UUID = uuid.New().String()
	friend.TelegramUsername = &username
	s.State.Friends = append(s.State.Friends, friend)
	s.State.NewFriend = friendship.Friend{}
	s.State.State = tgstate.None

	tgview.AddFromTelegram{}.Success(s, friend)
}

type RemoveFromTelegramHandler struct {
}

func (h *RemoveFromTelegramHandler) Callback() interface{} {
	return tgcallback.RemoveFromTelegram{}
}

func (h *RemoveFromTelegramHandler) HandleCallback(s *usersession.Session, clb interface{}) {
	s.State.Friends = friendship.RemoveTelegramFriends(s.State.Friends)
	tgview.RemoveFromTelegram{}.Success(s)
}

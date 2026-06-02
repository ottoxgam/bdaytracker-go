package usersession

import (
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/bdaytracker-go/internal/conf"
	"github.com/lodthe/bdaytracker-go/internal/tglimiter"
	"github.com/lodthe/bdaytracker-go/internal/tgstate"
)

type controllers struct {
	cfg        *conf.Config
	tgBot      *telegram.Bot
	tgExecutor *tglimiter.Executor

	repo tgstate.Repository
}

type Session struct {
	ctrl controllers

	TelegramID int
	LastUpdate *telegram.Update

	State *tgstate.State
}

func (s *Session) SaveState() error {
	return s.ctrl.repo.Save(s.State)
}

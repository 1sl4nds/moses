//go:build wireinject
// +build wireinject

package main

import (
	"github.com/1sl4nds/moses/client"
	"github.com/1sl4nds/moses/internal/command/authcommand"
	"github.com/1sl4nds/moses/internal/command/closecommand"
	"github.com/1sl4nds/moses/internal/command/countcommand"
	"github.com/1sl4nds/moses/internal/command/eventcommand"
	"github.com/1sl4nds/moses/internal/command/noticecommand"
	"github.com/1sl4nds/moses/internal/command/okcommand"
	"github.com/1sl4nds/moses/internal/command/requestcommand"
	"github.com/google/wire"
)

func buildClient() *client.Client {
	return client.New(&client.Options{
		ReadLimit: 2e6,
	})
}

func buildAuthCommand() *authcommand.AuthCommand {
	wire.Build(wire.NewSet(
		authcommand.New,
	))
	return &authcommand.AuthCommand{}
}

func buildCountCommand() *countcommand.CountCommand {
	wire.Build(
		countcommand.New,
	)
	return &countcommand.CountCommand{}
}

func buildCloseCommand() *closecommand.CloseCommand {
	wire.Build(
		closecommand.New,
	)
	return &closecommand.CloseCommand{}

}

func buildEventCommand() *eventcommand.EventCommand {
	wire.Build(
		wire.NewSet(
			buildClient,
			wire.Struct(new(eventcommand.Options), "Client"),
		),
		eventcommand.New,
	)
	return &eventcommand.EventCommand{}
}

func buildNoticeCommand() *noticecommand.NoticeCommand {
	wire.Build(
		noticecommand.New,
	)
	return &noticecommand.NoticeCommand{}
}

func buildOkCommand() *okcommand.OkCommand {
	wire.Build(
		okcommand.New,
	)
	return &okcommand.OkCommand{}

}

func buildRequestCommand() *requestcommand.RequestCommand {
	wire.Build(
		wire.NewSet(
			buildClient,
			wire.Struct(new(requestcommand.Options), "Client"),
		),
		requestcommand.New,
	)
	return &requestcommand.RequestCommand{}

}

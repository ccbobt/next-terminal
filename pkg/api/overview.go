package api

import (
	"github.com/labstack/echo/v4"
	"next-terminal/pkg/model"
)

type Counter struct {
	User          int64 `json:"user"`
	Asset         int64 `json:"asset"`
	Credential    int64 `json:"credential"`
	OnlineSession int64 `json:"onlineSession"`
}

func OverviewCounterEndPoint(c echo.Context) error {
	countUser, _ := model.CountUser()
	countOnlineSession, _ := model.CountOnlineSession()
	credential, _ := model.CountCredential()
	asset, _ := model.CountAsset()

	counter := Counter{
		User:          countUser,
		OnlineSession: countOnlineSession,
		Credential:    credential,
		Asset:         asset,
	}

	return Success(c, counter)
}

func OverviewSessionPoint(c echo.Context) (err error) {
	d := c.QueryParam("d")
	var results []model.D
	if d == "m" {
		results, err = model.CountSessionByDay(30)
	} else {
		results, err = model.CountSessionByDay(7)
	}
	if err != nil {
		return err
	}
	return Success(c, results)
}

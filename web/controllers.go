package web

import (
	"de.qaware.golang-merit-money/business"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	Index     = ""
	About     = "about.gohtml"
	Give      = "give.gohtml"
	Last      = "last.gohtml"
	ErrorPage = "error.gohtml"
)

type UserControllers struct {
	usecases *business.MeritMoney
}

func NewUserControllers(money *business.MeritMoney) UserControllers {
	controllers := UserControllers{
		usecases: money,
	}

	return controllers
}

func (u *UserControllers) Register(engine *gin.Engine) {
	engine.GET(Index, u.GetIndex)
	engine.GET(About, u.GetAbout)
	engine.GET(Give, u.GetGive)
	engine.GET(Last, u.GetLast)
}

func (u *UserControllers) GetIndex(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, About)
}

func (u *UserControllers) GetAbout(context *gin.Context) {
	context.HTML(200, About, nil)
}

func (u *UserControllers) GetGive(context *gin.Context) {
	users, err := u.usecases.AllUsers(); if err != nil {

	}
	context.HTML(200, Give, nil)
}

func (u *UserControllers) GetLast(context *gin.Context) {
	rewards, err := u.usecases.LastTenRewards()
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}

	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}

	dtos, err := u.rewardsToDto(rewards)
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}

	context.HTML(200, Last, gin.H{
		"Rewards": dtos,
	})
}

type rewardsDto struct {
	From   string
	To     string
	Amount string
	Date   string
	Note   string
}

func (u *UserControllers) rewardsToDto(rewards []business.Reward) ([]rewardsDto, error) {
	users, err := u.usecases.AllUsers()
	if err != nil {
		return nil, err
	}
	userMap := make(map[business.Uuid]business.User)
	for _, user := range users {
		userMap[user.Id] = user
	}

	dtos := make([]rewardsDto, 0)
	for _, reward := range rewards {
		dtos = append(dtos, rewardsDto{
			From:   userMap[reward.From].Name,
			To:     userMap[reward.To].Name,
			Amount: strconv.Itoa(int(reward.Amount)),
			Date:   reward.Date.Format("01-02-2006 15:04"),
			Note:   reward.Note,
		})
	}

	return dtos, nil
}

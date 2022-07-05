package web

import (
	"net/http"
	"strconv"

	"de.qaware.golang-merit-money/business"
	"github.com/gin-gonic/gin"
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
	engine.POST(Give, u.PostGive)
	engine.GET(Last, u.GetLast)
}

func (u *UserControllers) GetIndex(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, About)
}

func (u *UserControllers) GetAbout(context *gin.Context) {
	context.HTML(200, About, nil)
}

func (u *UserControllers) GetGive(context *gin.Context) {
	users, err := u.usecases.AllUsers()
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}
	userModels := u.usersToModel(users)
	giveModel := giveModel{
		Users:   userModels,
		PostUrl: Give,
	}

	context.HTML(200, Give, giveModel)
}

func (u *UserControllers) PostGive(context *gin.Context) {
	from, err := business.NewUuidFromString(context.PostForm("userFrom"))
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}
	to, err := business.NewUuidFromString(context.PostForm("userFor"))
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}
	amount, err := business.NewQaCoin(context.PostForm("amount"))
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}
	note := context.PostForm("note")
	err = u.usecases.GiveReward(from, to, amount, note)
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}

	context.Redirect(http.StatusSeeOther, Last)
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

	dtos, err := u.rewardsToModel(rewards)
	if err != nil {
		context.HTML(http.StatusInternalServerError, ErrorPage, nil)
		return
	}

	context.HTML(200, Last, gin.H{
		"Rewards": dtos,
	})
}

type giveModel struct {
	Users   []userModel
	PostUrl string
}

type userModel struct {
	Id   string
	Name string
}

func (u *UserControllers) usersToModel(users []business.User) []userModel {
	models := make([]userModel, 0)
	for _, user := range users {
		models = append(models, userModel{
			Id:   user.Id.String(),
			Name: user.Name,
		})
	}
	return models
}

type rewardsModel struct {
	From   string
	To     string
	Amount string
	Date   string
	Note   string
}

func (u *UserControllers) rewardsToModel(rewards []business.Reward) ([]rewardsModel, error) {
	users, err := u.usecases.AllUsers()
	if err != nil {
		return nil, err
	}
	userMap := make(map[business.Uuid]business.User)
	for _, user := range users {
		userMap[user.Id] = user
	}

	models := make([]rewardsModel, 0)
	for _, reward := range rewards {
		models = append(models, rewardsModel{
			From:   userMap[reward.From].Name,
			To:     userMap[reward.To].Name,
			Amount: strconv.Itoa(int(reward.Amount)),
			Date:   reward.Date.Format("01-02-2006 15:04"),
			Note:   reward.Note,
		})
	}

	return models, nil
}

package main

import (
	"flag"
	"fmt"

	"de.qaware.golang-merit-money/adapter"
	"de.qaware.golang-merit-money/business"
	"de.qaware.golang-merit-money/web"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	usersRepository := adapter.NewInMemoryUsersRepository()
	meritMoney := business.NewMeritMoney(
		usersRepository,
		adapter.NewInMemoryRewardsRepository(),
		&adapter.RealClock{},
	)

	var developmentMode bool
	flag.BoolVar(&developmentMode, "development", false, "true|false initializes the application with debug data. default: false")
	flag.Parse()
	if developmentMode {
		log.SetLevel(log.DebugLevel)
		fmt.Printf("development mode is %s on")
		initDebugData(meritMoney, usersRepository)

	}

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	userController := web.NewUserControllers(meritMoney)
	userController.Register(engine)

	_ = engine.Run(":8080")
}

func initDebugData(meritMoney *business.MeritMoney, usersRepository *adapter.InMemoryUsersRepository) {
	rewards := []business.Reward{
		{
			From:   1,
			To:     2,
			Note:   "Thanks for helping me out the other day.",
			Amount: 1,
		},
		{
			From:   3,
			To:     4,
			Note:   "It is fun working with you.",
			Amount: 2,
		},
		{
			From:   2,
			To:     1,
			Note:   "Thanks for finding my bug.",
			Amount: 3,
		},
	}
	for _, reward := range rewards {
		err := meritMoney.GiveReward(reward.From, reward.To, reward.Amount, reward.Note)
		if err != nil {
			log.Panic("could not store reward debug data")
		}
	}

	users := []business.User{
		{
			Id:   1,
			Name: "Timo",
		},
		{
			Id:   2,
			Name: "Alex",
		},
		{
			Id:   3,
			Name: "Dirk",
		},
		{
			Id:   4,
			Name: "Markus",
		},
	}

	for _, user := range users {
		err := usersRepository.Store(user)
		if err != nil {
			log.Panicf("could initialize debug data %v", user)
		}
	}
}

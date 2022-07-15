package business

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Uuid uint32

func (u Uuid) String() string {
	return strconv.Itoa(int(u))
}

func NewUuidFromString(from string) (Uuid, error) {
	num, err := strconv.ParseUint(from, 10, 32)
	if err != nil {
		return Uuid(0), errors.Wrapf(err, "could not parse %s as Uuid", from)
	}
	return Uuid(num), nil
}

type User struct {
	Id   Uuid
	Name string
}

type QaCoin uint32

func NewQaCoin(from string) (QaCoin, error) {
	num, err := strconv.ParseUint(from, 10, 32)
	if err != nil {
		return QaCoin(0), errors.Wrapf(err, "could not parse %s as QaCoin", from)
	}
	return QaCoin(num), nil
}

type Reward struct {
	From   Uuid
	To     Uuid
	Amount QaCoin
	Date   time.Time
	Note   string
}

type UsersRepository interface {
	All() ([]User, error)
	Store(newUser User) error
}

type RewardsRepository interface {
	All() ([]Reward, error)
	Last(limit uint) ([]Reward, error)
	Store(reward Reward) error
}
type Clock interface {
	Now() time.Time
}

type MeritMoney struct {
	users   UsersRepository
	Rewards RewardsRepository
	Clock   Clock
}

func NewMeritMoney(users UsersRepository, rewards RewardsRepository, clock Clock) *MeritMoney {
	return &MeritMoney{users: users, Rewards: rewards, Clock: clock}
}

func (m *MeritMoney) GiveReward(from Uuid, to Uuid, amount QaCoin, note string) error {
	return m.Rewards.Store(
		Reward{
			From:   from,
			To:     to,
			Amount: amount,
			Date:   m.Clock.Now(),
			Note:   note,
		},
	)
}

func (m *MeritMoney) AllUsers() ([]User, error) {
	return m.users.All()
}

func (m *MeritMoney) FindById(id Uuid) (bool, User, error) {
	users, err := m.users.All()
	if err != nil {
		return false, User{}, err
	}
	foundUser := sliceFind(users, func(user User) bool { return user.Id == id })
	return true, *foundUser, nil
}

func (m *MeritMoney) LastTenRewards() ([]Reward, error) {
	return m.Rewards.Last(10)
}

func sliceFind[T comparable](slice []T, predicate func(T) bool) *T {
	for _, element := range slice {
		if predicate(element) {
			return &element
		}
	}
	return nil
}

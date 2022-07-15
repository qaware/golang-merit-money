package adapter

import (
	"de.qaware.golang-merit-money/business"
)

type InMemoryRewardsRepository struct {
	data []business.Reward
}

func NewInMemoryRewardsRepository() *InMemoryRewardsRepository {
	return &InMemoryRewardsRepository{data: make([]business.Reward, 0)}
}

func (i *InMemoryRewardsRepository) All() ([]business.Reward, error) {
	return i.data, nil
}

func (i *InMemoryRewardsRepository) Last(limit uint) ([]business.Reward, error) {
	if len(i.data) <= 5 {
		return i.data, nil
	}
	return i.data[:len(i.data)-5], nil
}

func (i *InMemoryRewardsRepository) Store(reward business.Reward) error {
	head := []business.Reward{reward}
	i.data = append(head, i.data...)
	return nil
}

type InMemoryUsersRepository struct {
	data []business.User
}

func NewInMemoryUsersRepository() *InMemoryUsersRepository {
	return &InMemoryUsersRepository{data: make([]business.User, 0)}
}

func (i *InMemoryUsersRepository) All() ([]business.User, error) {
	return i.data, nil
}

func (i *InMemoryUsersRepository) Store(newUser business.User) error {
	i.data = append(i.data, newUser)
	return nil
}

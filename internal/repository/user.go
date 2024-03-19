package repository

import (
	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUser(param *model.GetUserParam) (*entity.User, error)
	GetUserColoumn(colName string, value string) (*entity.User, error)
	UpdateUserData(updatedUser *entity.User) error
	DeleteUser(user entity.User) error
	GetUserSubscriptionStatus(userID string) (int, error)
	GetUserJournalingCount(userID string) (int, error)
	UpdateUserColoumn(colname string, updated *entity.User, userID string) error
	GetUserJoinedEvents(userID string) ([]entity.Event, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) InterfaceUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := ur.db.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUser(param *model.GetUserParam) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.Debug().Where(&param).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserColoumn(colName string, value string) (*entity.User, error) {
	var user entity.User
	err := ur.db.Debug().Where(colName+" = ?", value).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUserData(updatedUser *entity.User) error {
	err := ur.db.Debug().Save(updatedUser).Error

	return err
}

func (ur *UserRepository) UpdateUserColoumn(colname string, updated *entity.User, userID string) error {
	err := ur.db.Debug().Select(colname).Where("id = ?", userID).Updates(updated).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(user entity.User) error {
	err := ur.db.Debug().Delete(user).Error

	return err
}

func (ur *UserRepository) GetUserSubscriptionStatus(userID string) (int, error) {
	var user entity.User

	var subscribtionStatus int

	err := ur.db.Debug().Where("id = ?", userID).First(&user).Error

	if user.IsSubscribed {
		subscribtionStatus = 1
	} else {
		subscribtionStatus = 0
	}

	if err != nil {
		return subscribtionStatus, err
	}

	return subscribtionStatus, nil
}

func (ur *UserRepository) GetUserJournalingCount(userID string) (int, error) {
	var user entity.User
	var count int

	err := ur.db.Debug().Where("id = ?", userID).First(&user).Error

	if err != nil {
		return -1, err
	}

	count = user.JournalingEntryCount

	return count, nil
}

func (ur *UserRepository) GetUserJoinedEvents(userID string) ([]entity.Event, error) {
	var user entity.User
	var events []entity.PaymentEvent
	var joinedEvents []entity.Event

	err := ur.db.Debug().Where("id = ?", userID).Preload("PaymentEvent").First(&user).Error

	if err != nil {
		return nil, err
	}

	events = user.PaymentEvent

	for _, event := range events {
		var eventDetails entity.Event

		if event.IsCompleted {
			err := ur.db.Debug().Where("id = ?", event.EventID).Preload("EventImage").First(&eventDetails).Error
			if err != nil {
				return nil, err
			}

			joinedEvents = append(joinedEvents, eventDetails)
		}
	}

	return joinedEvents, nil
}

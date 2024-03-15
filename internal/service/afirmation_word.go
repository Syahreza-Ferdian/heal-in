package service

import (
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
)

type InterfaceAfirmationWordService interface {
	GetAfirmationWordByMoodID(moodID int) (string, error)
}

type AfirmationWordService struct {
	awr repository.InterfaceAfirmationWordRepository
}

func NewAfirmationWordService(awr repository.InterfaceAfirmationWordRepository) InterfaceAfirmationWordService {
	return &AfirmationWordService{
		awr: awr,
	}
}

func (aws *AfirmationWordService) GetAfirmationWordByMoodID(moodID int) (string, error) {
	affWord, err := aws.awr.GetRandomWordByMoodID(moodID)

	if err != nil {
		return "", err
	}

	return affWord.Word, nil
}

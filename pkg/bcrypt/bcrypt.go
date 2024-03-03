package bcrypt

import bcrypt_import "golang.org/x/crypto/bcrypt"

type BcryptInterface interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}

type Bcrypt struct {
	cost int
}

func NewBcrypt(cost int) BcryptInterface {
	return &Bcrypt{
		cost: cost,
	}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt_import.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (b *Bcrypt) ComparePassword(hashedPassword string, password string) error {
	err := bcrypt_import.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

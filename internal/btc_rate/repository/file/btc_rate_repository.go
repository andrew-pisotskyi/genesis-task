package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/andrew-pisotskyi/genesis-task/domain"
	log "github.com/sirupsen/logrus"
)

type btrRateRepository struct {
	fileName string
}

func NewBtcRateRepository(fileName string) domain.BtcRateRepositoryInterface {
	return &btrRateRepository{fileName}
}

func (brr *btrRateRepository) ExistsEmail(email string) bool {
	emails := brr.GetAllEmails()
	for _, e := range emails {
		if e == email {
			return true
		}
	}

	return false
}

func (brr *btrRateRepository) SaveEmail(email string) error {
	file, err := brr.getFile()
	if err != nil {
		file.Close()
		return err
	}
	fmt.Fprintln(file, email)
	file.Close()

	return nil
}

func (brr *btrRateRepository) GetAllEmails() []string {
	b, err := os.ReadFile(brr.fileName)
	if err != nil {
		return []string{}
	}
	var emails []string
	data := strings.Split(string(b), "\n")
	for _, item := range data {
		if len(item) > 0 {
			emails = append(emails, item)
		}
	}

	return emails
}

func (brr *btrRateRepository) getFile() (*os.File, error) {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorf("Something went wrong with creating data file. err: %s", err.Error())
		return nil, err
	}
	return file, nil
}

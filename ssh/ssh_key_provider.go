package ssh

import (
	"bufio"
	"errors"
	"github.com/docker/machine/drivers/virtualbox"
	"os"
)

type SSHKeyProvider interface {
	Generate(privateKeyPath string) error
	GetPublicKey(privateKeyPath string) (string, error)
}

type defaultSSHKeyProvider struct {
}

func (d defaultSSHKeyProvider) Generate(privateKeyPath string) error {
	return virtualbox.NewSSHKeyGenerator().Generate(privateKeyPath)
}

func (d defaultSSHKeyProvider) GetPublicKey(privateKeyPath string) (string, error) {
	file, err := os.Open(privateKeyPath + ".pub")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", errors.New("private key not found")
}

func NewSSHKeyProvider() SSHKeyProvider {
	return &defaultSSHKeyProvider{}
}

package addressbook

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Contact struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Groups  []string `json:"groups,omitempty"`
	Notes   string   `json:"notes,omitempty"`
}

type AddressBook struct {
	Contacts []Contact `json:"contacts"`
}

const (
	defaultAddressBookDir  = ".email-cli"
	defaultAddressBookFile = "contacts.json"
)

func getAddressBookPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	abDir := filepath.Join(homeDir, defaultAddressBookDir)
	if err := os.MkdirAll(abDir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(abDir, defaultAddressBookFile), nil
}

func LoadAddressBook() (*AddressBook, error) {
	abPath, err := getAddressBookPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(abPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &AddressBook{Contacts: []Contact{}}, nil
		}
		return nil, err
	}

	var ab AddressBook
	if err := json.Unmarshal(data, &ab); err != nil {
		return nil, err
	}

	return &ab, nil
}

func SaveAddressBook(ab *AddressBook) error {
	abPath, err := getAddressBookPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(ab, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(abPath, data, 0600)
}

func (ab *AddressBook) AddContact(contact Contact) error {
	ab.Contacts = append(ab.Contacts, contact)
	return SaveAddressBook(ab)
}

func (ab *AddressBook) RemoveContact(email string) error {
	for i, contact := range ab.Contacts {
		if contact.Email == email {
			ab.Contacts = append(ab.Contacts[:i], ab.Contacts[i+1:]...)
			return SaveAddressBook(ab)
		}
	}
	return nil
}

func (ab *AddressBook) UpdateContact(email string, newContact Contact) error {
	for i, contact := range ab.Contacts {
		if contact.Email == email {
			ab.Contacts[i] = newContact
			return SaveAddressBook(ab)
		}
	}
	return nil
}

func (ab *AddressBook) GetContact(email string) *Contact {
	for _, contact := range ab.Contacts {
		if contact.Email == email {
			return &contact
		}
	}
	return nil
}

func (ab *AddressBook) GetContactsByGroup(group string) []Contact {
	var contacts []Contact
	for _, contact := range ab.Contacts {
		for _, g := range contact.Groups {
			if g == group {
				contacts = append(contacts, contact)
				break
			}
		}
	}
	return contacts
}
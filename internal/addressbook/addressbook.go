package addressbook

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AddressBook struct {
	Contacts []Contact
}

func NewAddressBook() *AddressBook {
	return &AddressBook{}
}

func (ab *AddressBook) AddContact(name, email string) {
	ab.Contacts = append(ab.Contacts, Contact{Name: name, Email: email})
}

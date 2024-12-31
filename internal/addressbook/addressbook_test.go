package addressbook

import (
	"testing"
	"os"
	"path/filepath"
)

func TestAddressBookOperations(t *testing.T) {
	// Setup temporary directory
	originalHome := os.Getenv("HOME")
	tmpHome := t.TempDir()
	os.Setenv("HOME", tmpHome)
	defer os.Setenv("HOME", originalHome)

	// Test loading empty address book
	ab, err := LoadAddressBook()
	if err != nil {
		t.Fatalf("Failed to load address book: %v", err)
	}

	// Test adding contact
	contact := Contact{
		Name: "John Doe",
		Email: "john@test.com",
		Groups: []string{"friends"},
		Notes: "Test contact",
	}

	err = ab.AddContact(contact)
	if err != nil {
		t.Errorf("Failed to add contact: %v", err)
	}

	// Test getting contact
	retrieved := ab.GetContact("john@test.com")
	if retrieved == nil {
		t.Error("Expected contact, got nil")
	}

	// Test getting contacts by group
	groupContacts := ab.GetContactsByGroup("friends")
	if len(groupContacts) != 1 {
		t.Error("Expected 1 contact in group")
	}

	// Test removing contact
	err = ab.RemoveContact("john@test.com")
	if err != nil {
		t.Errorf("Failed to remove contact: %v", err)
	}

	// Verify contact was removed
	ab, _ = LoadAddressBook()
	if len(ab.Contacts) != 0 {
		t.Error("Expected empty contacts after removal")
	}
}
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"net/http"

	"./internal/email"
	"./internal/logging"
)

func main() {
	// Initialize flags
		abCmd       = flag.String("ab", "", "Address book command: add-contact, remove-contact, list-contacts, get-contact")
		contactName  = flag.String("contact-name", "", "Contact name")
		contactEmail = flag.String("contact-email", "", "Contact email")
		contactGroup = flag.String("contact-group", "", "Contact group")
		contactNotes = flag.String("contact-notes", "", "Contact notes")
		configCmd   = flag.String("config", "", "Config command: add-profile, remove-profile, list-profiles")
		profileName = flag.String("profile-name", "", "SMTP profile name")
		smtpHost    = flag.String("smtp-host", "", "SMTP host")
		smtpPort    = flag.Int("smtp-port", 0, "SMTP port")
		smtpUser    = flag.String("smtp-user", "", "SMTP username")
		smtpPass    = flag.String("smtp-pass", "", "SMTP password")
		scheduleCmd = flag.String("schedule", "", "Schedule email: now, later")
		sendAt      = flag.String("send-at", "", "Send time (format: 2006-01-02 15:04:05)")
		recurring   = flag.String("recurring", "", "Recurring schedule: daily, weekly, monthly")
		html        = flag.String("html", "", "HTML version of email body")
		cc          = flag.String("cc", "", "CC recipients (comma-separated)")
		bcc         = flag.String("bcc", "", "BCC recipients (comma-separated)")
		replyTo     = flag.String("reply-to", "", "Reply-To email address")
		setDefault  = flag.Bool("set-default", false, "Set as default profile")
	var (
		to          = flag.String("to", "", "Recipient email address")
		subject     = flag.String("subject", "", "Email subject")
		body        = flag.String("body", "", "Email body")
		from        = flag.String("from", "", "Sender email address")
		template    = flag.String("template", "", "Email template name")
		name        = flag.String("name", "", "Recipient name for template")
	// Validate required flags
	var errors []string
	if *to == "" {
		errors = append(errors, "recipient email address (-to) is required")
	}
	if *from == "" {
		errors = append(errors, "sender email address (-from) is required")
	}
	if *subject == "" {
		errors = append(errors, "email subject (-subject) is required")
	}
	// Handle address book commands
	if *abCmd != "" {
		ab, err := addressbook.LoadAddressBook()
		if err != nil {
			logging.Error("Failed to load address book: %v", err)
			os.Exit(1)
		}

		switch *abCmd {
		case "add-contact":
			if *contactName == "" || *contactEmail == "" {
				logging.Error("Contact name and email are required")
				os.Exit(1)
			}
			contact := addressbook.Contact{
				Name:  *contactName,
				Email: *contactEmail,
				Notes: *contactNotes,
			}
			if *contactGroup != "" {
				contact.Groups = []string{*contactGroup}
			}
			if err := ab.AddContact(contact); err != nil {
				logging.Error("Failed to add contact: %v", err)
				os.Exit(1)
			}
			fmt.Printf("Contact '%s' added successfully\n", *contactName)
			os.Exit(0)

		case "remove-contact":
			if *contactEmail == "" {
				logging.Error("Contact email is required")
				os.Exit(1)
			}
			if err := ab.RemoveContact(*contactEmail); err != nil {
				logging.Error("Failed to remove contact: %v", err)
				os.Exit(1)
			}
			fmt.Printf("Contact '%s' removed successfully\n", *contactEmail)
			os.Exit(0)

		case "list-contacts":
			fmt.Println("Contacts:")
			for _, contact := range ab.Contacts {
				fmt.Printf("- %s <%s>\n", contact.Name, contact.Email)
				if len(contact.Groups) > 0 {
					fmt.Printf("  Groups: %s\n", strings.Join(contact.Groups, ", "))
				}
				if contact.Notes != "" {
					fmt.Printf("  Notes: %s\n", contact.Notes)
				}
			}
			os.Exit(0)

		case "get-contact":
			if *contactEmail == "" {
				logging.Error("Contact email is required")
				os.Exit(1)
			}
			contact := ab.GetContact(*contactEmail)
			if contact == nil {
				logging.Error("Contact not found")
				os.Exit(1)
			}
			fmt.Printf("Name: %s\nEmail: %s\n", contact.Name, contact.Email)
			if len(contact.Groups) > 0 {
				fmt.Printf("Groups: %s\n", strings.Join(contact.Groups, ", "))
			}
			if contact.Notes != "" {
				fmt.Printf("Notes: %s\n", contact.Notes)
			}
			os.Exit(0)

		default:
			logging.Error("Unknown address book command: %s", *abCmd)
			os.Exit(1)
		}
	}
	if *body == "" && *template == "" {
		errors = append(errors, "either email body (-body) or template (-template) is required")
	}

	// Handle config commands
	if *configCmd != "" {
		conf, err := config.LoadConfig()
		if err != nil {
			logging.Error("Failed to load config: %v", err)
			os.Exit(1)
		}

	// Load configuration
	conf, err := config.LoadConfig()
	if err != nil {
		logging.Error("Failed to load config: %v", err)
		os.Exit(1)
	}

	profile := config.GetDefaultProfile(conf)
	if profile == nil {
		logging.Error("No default SMTP profile configured")
		os.Exit(1)
	}

	// Create email configuration
	config := email.EmailConfig{
		SMTPHost:     profile.Host,
		SMTPPort:     profile.Port,
		SMTPUsername: profile.Username,
		SMTPPassword: profile.Password,
	}
				Host:      *smtpHost,
				Port:      *smtpPort,
				Username:  *smtpUser,
				Password:  *smtpPass,
				IsDefault: *setDefault,
			}
			if err := config.AddProfile(conf, profile); err != nil {
				logging.Error("Failed to add profile: %v", err)
				os.Exit(1)
			}
			fmt.Printf("Profile '%s' added successfully\n", *profileName)
			os.Exit(0)

		case "remove-profile":
			if *profileName == "" {
				logging.Error("Profile name is required")
				os.Exit(1)
			}
			if err := config.RemoveProfile(conf, *profileName); err != nil {
				logging.Error("Failed to remove profile: %v", err)
				os.Exit(1)
			}
			fmt.Printf("Profile '%s' removed successfully\n", *profileName)
			os.Exit(0)

		case "list-profiles":
			fmt.Println("SMTP Profiles:")
			for _, profile := range conf.SMTPProfiles {
				fmt.Printf("- %s (Host: %s, Port: %d)%s\n",
					profile.Name, profile.Host, profile.Port,
					if profile.IsDefault ? " [DEFAULT]" : "")
			}
			os.Exit(0)

		default:
			logging.Error("Unknown config command: %s", *configCmd)
			os.Exit(1)
		}
	}
	if len(errors) > 0 {
		logging.Error("Command line validation failed")
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		flag.Usage()
		os.Exit(1)
			if err != nil {
				logging.Error("Failed to read attachment %s: %v", file, err)
				continue
			}
			logging.Info("Successfully processed attachment: %s", file)
		os.Exit(1)
	}

	// Create email configuration
	config := email.EmailConfig{
		SMTPHost:     "smtp.gmail.com",
		SMTPPort:     587,
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}

	emailMsg := email.Email{
		To:          strings.Split(*to, ","),
		From:        *from,
		Subject:     *subject,
		Body:        *body,
		Attachments: attachmentList,
	if err := email.SendEmail(config, emailMsg); err != nil {
		logging.Error("Failed to send email: %v", err)
		os.Exit(1)
	}

	logging.Info("Email sent successfully to %s", *to)
	fmt.Println("Email sent successfully")
			attachmentList = append(attachmentList, email.Attachment{
				Filename:    filepath.Base(file),
				ContentType: contentType,
				Data:        data,
			})
		}
	}
	// Create email message
	emailMsg := email.Email{
		To:      strings.Split(*to, ","),
		From:    *from,
		Subject: *subject,
		Body:    *body,
	}

	// Send email
	if err := email.SendEmail(config, emailMsg); err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully")
}
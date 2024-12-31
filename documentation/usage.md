# Email CLI Usage Guide

## Basic Usage

```bash
email-cli -to recipient@example.com -from sender@example.com -subject "Test" -body "Hello"
```

## Configuration

### Managing SMTP Profiles

```bash
# Add a new SMTP profile
email-cli -config add-profile -profile-name "gmail" -smtp-host smtp.gmail.com -smtp-port 587 -smtp-user "user@gmail.com" -smtp-pass "password" -set-default

# List profiles
email-cli -config list-profiles

# Remove profile
email-cli -config remove-profile -profile-name "gmail"
```

## Address Book Management

```bash
# Add contact
email-cli -ab add-contact -contact-name "John Doe" -contact-email "john@example.com" -contact-group "work" -contact-notes "Project lead"

# List contacts
email-cli -ab list-contacts

# Remove contact
email-cli -ab remove-contact -contact-email "john@example.com"
```

## Using Templates

```bash
email-cli -template welcome -name "John" -to "john@example.com" -from "sender@example.com"
```

## Attachments

```bash
email-cli -to "recipient@example.com" -from "sender@example.com" -subject "Files" -body "See attached" -attach "file1.pdf,file2.txt"
```

## HTML Emails

```bash
email-cli -to "recipient@example.com" -from "sender@example.com" -subject "HTML Test" -html "<h1>Hello</h1>"
```

## Email Scheduling

```bash
email-cli -to "recipient@example.com" -schedule later -send-at "2024-01-01 10:00:00" -recurring daily
```
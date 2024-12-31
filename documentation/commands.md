# Command Line Reference

## Email Sending Options

| Flag | Description | Required | Example |
|------|-------------|-----------|----------|
| `-to` | Recipient email address | Yes | `-to user@example.com` |
| `-from` | Sender email address | Yes | `-from sender@example.com` |
| `-subject` | Email subject | Yes | `-subject "Meeting Notes"` |
| `-body` | Email body text | No* | `-body "Hello World"` |
| `-html` | HTML version of email | No | `-html "<h1>Hello</h1>"` |
| `-cc` | CC recipients | No | `-cc "user1@example.com,user2@example.com"` |
| `-bcc` | BCC recipients | No | `-bcc "manager@example.com"` |
| `-reply-to` | Reply-To address | No | `-reply-to support@example.com` |
| `-attach` | File attachments | No | `-attach "report.pdf,image.jpg"` |

## Template Options

| Flag | Description | Required | Example |
|------|-------------|-----------|----------|
| `-template` | Template name | No* | `-template welcome` |
| `-name` | Recipient name for template | No | `-name "John Doe"` |

## Configuration Options

| Flag | Description | Example |
|------|-------------|----------|
| `-config` | Configuration command | `-config add-profile` |
| `-profile-name` | SMTP profile name | `-profile-name "gmail"` |
| `-smtp-host` | SMTP server host | `-smtp-host smtp.gmail.com` |
| `-smtp-port` | SMTP server port | `-smtp-port 587` |
| `-smtp-user` | SMTP username | `-smtp-user "user@gmail.com"` |
| `-smtp-pass` | SMTP password | `-smtp-pass "password"` |
| `-set-default` | Set as default profile | `-set-default` |

## Address Book Options

| Flag | Description | Example |
|------|-------------|----------|
| `-ab` | Address book command | `-ab add-contact` |
| `-contact-name` | Contact name | `-contact-name "John Doe"` |
| `-contact-email` | Contact email | `-contact-email "john@example.com"` |
| `-contact-group` | Contact group | `-contact-group "work"` |
| `-contact-notes` | Contact notes | `-contact-notes "Project lead"` |

## Scheduling Options

| Flag | Description | Example |
|------|-------------|----------|
| `-schedule` | Schedule type | `-schedule later` |
| `-send-at` | Send time | `-send-at "2024-01-01 10:00:00"` |
| `-recurring` | Recurring schedule | `-recurring daily` |
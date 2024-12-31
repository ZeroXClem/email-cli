# Email CLI in Go

A powerful command-line email client with advanced features including template support, address book management, email scheduling, and secure configuration storage.

## Features

- 📧 Send emails via SMTP
- 📁 Multiple SMTP profile support
- 📝 Email template system
- 👥 Address book management
- 📎 File attachment support
- 📅 Email scheduling
- 🔒 Secure configuration storage
- 🌐 HTML email support
- 👥 CC/BCC support
- ↩️ Reply-To header support

## Installation

```bash
go install github.com/zeroxclem/email-cli@latest
```

## Quick Start

```bash
# Send a simple email
email-cli -to recipient@example.com -from sender@example.com -subject "Hello" -body "World"

# Send with attachment
email-cli -to recipient@example.com -from sender@example.com -subject "Report" -body "Please find attached" -attach "report.pdf"
```

## Documentation

- [Usage Guide](documentation/usage.md)
- [Command Reference](documentation/commands.md)
- [Configuration Examples](documentation/config-examples.md)
- [Contributing Guidelines](documentation/contributing.md)

## Configuration

The CLI supports multiple SMTP profiles and secure storage of credentials. See the [configuration examples](documentation/config-examples.md) for detailed setup instructions.

## Address Book

Manage contacts and groups for easy email sending:

```bash
# Add a contact
email-cli -ab add-contact -contact-name "John" -contact-email "john@example.com"

# List contacts
email-cli -ab list-contacts
```

## Templates

Use email templates for consistent communication:

```bash
email-cli -template welcome -name "John" -to "john@example.com"
```

## Security

- Encrypted storage of sensitive data
- TLS support for SMTP connections
- Input validation and sanitization

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](documentation/contributing.md) before submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

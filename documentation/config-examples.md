# Configuration Examples

## SMTP Profiles

### Gmail Configuration
```json
{
  "smtp_profiles": [
    {
      "name": "gmail",
      "host": "smtp.gmail.com",
      "port": 587,
      "username": "your.email@gmail.com",
      "password": "your-app-specific-password",
      "is_default": true
    }
  ]
}
```

### Office 365 Configuration
```json
{
  "smtp_profiles": [
    {
      "name": "office365",
      "host": "smtp.office365.com",
      "port": 587,
      "username": "your.email@company.com",
      "password": "your-password",
      "is_default": false
    }
  ]
}
```

## Address Book Example
```json
{
  "contacts": [
    {
      "name": "John Doe",
      "email": "john@example.com",
      "groups": ["work", "project-alpha"],
      "notes": "Project lead"
    },
    {
      "name": "Jane Smith",
      "email": "jane@example.com",
      "groups": ["work"],
      "notes": "Developer"
    }
  ]
}
```

## Email Templates

### Welcome Template
```
Subject: Welcome to {{.Company}}

Dear {{.Name}},

Welcome to {{.Company}}! We're excited to have you on board.

Best regards,
The {{.Company}} Team
```

### Meeting Invitation Template
```
Subject: Meeting: {{.Subject}}

Hi {{.Name}},

You are invited to a meeting:

Topic: {{.Subject}}
Date: {{.Date}}
Time: {{.Time}}
Location: {{.Location}}

Please confirm your attendance.

Best regards,
{{.Organizer}}
```
package email

type Template string

const (
	TemplateWelcome             Template = "Welcome"
	TemplateDueDateReminder     Template = "due-date-reminder"
	TemplateOverdueNotification Template = "overdue-notification"
	TemplateWeeklyReport        Template = "weekly-report"
)

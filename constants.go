package main

import (
	"Shelex/tg-watermark-bot/locale"
	"fmt"
)

const helpTextTemplate = `
%s.
%s
%s
%s:
	- /clear_watermark - %s
	- /status - %s
	- /help - %s
`

func HelpText() string {
	return fmt.Sprintf(helpTextTemplate,
		locale.Translate("help_header"),
		locale.Translate("help_if_no_watermark"),
		locale.Translate("help_notify_transparency_issue"),
		locale.Translate("help_commands_list_header"),
		locale.Translate("help_command_list_clear"),
		locale.Translate("help_command_list_status"),
		locale.Translate("help_command_list_help"))
}

func Greet(name string) string {
	return fmt.Sprintf("%s, %s :)", locale.Translate("hello"), name)
}

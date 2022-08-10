package emailSender

import (
	"html/template"
)

type EmailTemplate struct {
	templates *template.Template
}

func LoadTemplate(filename string) (*EmailTemplate, error) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	return &EmailTemplate{templates: t}, nil
}
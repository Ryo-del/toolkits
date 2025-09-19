package docker

import (
	"fmt"
	"strings"
)

// глобальные переменные
var ProgramLang string
var WorkDir string
var File string
var Launch string

// отдельные для образов
var Python_Image string
var Go_Image string
var Java_Image string
var Node_Image string

// GenerateDockerfile формирует Dockerfile в зависимости от языка
func GenerateDockerfile() string {
	var sb strings.Builder

	switch ProgramLang {
	case "Python":
		if Python_Image == "" {
			Python_Image = "python:3.14" // дефолт
		}
		sb.WriteString(fmt.Sprintf("FROM %s\n", Python_Image))
		if WorkDir != "" {
			sb.WriteString(fmt.Sprintf("WORKDIR %s\n", WorkDir))
		}
		if File != "" {
			sb.WriteString(fmt.Sprintf("COPY %s .\n", "."))
			sb.WriteString(fmt.Sprintf("RUN pip install -r %s\n", File))
		}
		if Launch != "" {
			sb.WriteString(fmt.Sprintf("CMD %s\n", Launch))
		}

	case "Go":
		if Go_Image == "" {
			Go_Image = "golang:1.22"
		}
		sb.WriteString(fmt.Sprintf("FROM %s\n", Go_Image))
		if WorkDir != "" {
			sb.WriteString(fmt.Sprintf("WORKDIR %s\n", WorkDir))
		}
		if File != "" {
			sb.WriteString(fmt.Sprintf("COPY %s .\n", "."))
		}
		sb.WriteString("RUN go mod tidy\n")
		if Launch != "" {
			sb.WriteString(fmt.Sprintf("CMD %s\n", Launch))
		}

	case "Java":
		if Java_Image == "" {
			Java_Image = "openjdk:17"
		}
		sb.WriteString(fmt.Sprintf("FROM %s\n", Java_Image))
		if WorkDir != "" {
			sb.WriteString(fmt.Sprintf("WORKDIR %s\n", WorkDir))
		}
		if File != "" {
			sb.WriteString(fmt.Sprintf("COPY %s .\n", "."))
		}
		if Launch != "" {
			sb.WriteString(fmt.Sprintf("CMD %s\n", Launch))
		}

	case "Node.js":
		if Node_Image == "" {
			Node_Image = "node:20"
		}
		sb.WriteString(fmt.Sprintf("FROM %s\n", Node_Image))
		if WorkDir != "" {
			sb.WriteString(fmt.Sprintf("WORKDIR %s\n", WorkDir))
		}
		if File != "" {
			sb.WriteString(fmt.Sprintf("COPY %s .\n", ".")) // package.json
		}
		sb.WriteString("RUN npm install\n")
		if Launch != "" {
			sb.WriteString(fmt.Sprintf("CMD %s\n", Launch))
		}

	default:
		return "# Ошибка: язык не выбран"
	}

	return sb.String()
}

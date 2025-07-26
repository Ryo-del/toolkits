package parser

import (
	//"encoding/json"
	"bufio"
	"errors"
	"os"
	"regexp"
)

// Структура для хранения информации из строки лога
type LogEntry struct {
	IP         string
	Method     string
	Path       string
	StatusCode string
	ReqTime    string
}

// Ошибка, если строка не соответствует формату
var ErrNoMatch = errors.New("log line does not match expected format")

// Регулярка для парсинга строки NGINX-лога
var logRegex = regexp.MustCompile(`(?P<ip>\d+\.\d+\.\d+\.\d+).*?"(?P<method>[A-Z]+) (?P<path>[^ ]+) HTTP/[^"]+" (?P<status>\d+) .*? (?P<reqtime>\d+\.\d+)`)

// Функция парсинга одной строки
func ParseLine(line string) (*LogEntry, error) {
	matches := logRegex.FindStringSubmatch(line)
	if len(matches) < 6 {
		return nil, ErrNoMatch // строка не соответствует формату
	}

	return &LogEntry{
		IP:         matches[1],
		Method:     matches[2],
		Path:       matches[3],
		StatusCode: matches[4],
		ReqTime:    matches[5],
	}, nil
}

// Функция парсинга всего файла
func ParseFile(path string) ([]*LogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []*LogEntry

	for scanner.Scan() {
		entry, err := ParseLine(scanner.Text())
		if err != nil {
			if err == ErrNoMatch {
				continue // пропускаем некорректные строки
			}
			return nil, err // другие ошибки — прерываем
		}
		if entry != nil {
			entries = append(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

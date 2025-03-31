package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.infoLog.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.errorLog.Println(v...)
}


func (l *Logger) FileLocationSet(filename string) {
	if filename == "" {
			log.Fatal("Filename cannot be empty")
	}

	_, err := os.Stat(filename)
	if err == nil { 
			log.Fatal("File already exists: ", filename) 
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
			l.errorLog.Fatal("Error opening file: ", err) 
	}

	l.infoLog.SetOutput(file)
	l.errorLog.SetOutput(file)
}
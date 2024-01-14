/* Copyright 2024 Take Control - Software & Infrastructure

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package go_multi_log

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
	"github.com/takecontrolsoft/go_multi_log/logger"
	"github.com/takecontrolsoft/go_multi_log/logger/levels"
	"github.com/takecontrolsoft/go_multi_log/logger/loggers"
)

type person struct{ Name string }
type car struct{ Year string }

func TestAllLogs(t *testing.T) {
	content := readConsole(func() {
		logger.Debug("Test debug log message")
		logger.Trace("Test trace log message")
		logger.Info("Test info log message")

		person := person{Name: "Michael"}
		logger.Info(person)

		logger.Warning("Test warning log message")
		logger.Error(errors.Errorf("Test error object").Err)
		logger.Error("Test error log message")
		// The function Fatal logs the messages and then calls Panic
		// logger.Fatal("Test log [fatal] message")
	})

	// Default log level is levels.InfoLevel
	assert.NotContains(t, content, "DEBUG: [Test debug log message]")
	assert.NotContains(t, content, "TRACE: [Test trace log message]")

	assert.Contains(t, content, "INFO: [Test info log message]")
	assert.Contains(t, content, "INFO: [{Michael}]")

	assert.Contains(t, content, "WARNING: [Test warning log message]")
	assert.Contains(t, content, "Test error object")
	assert.Contains(t, content, "ERROR: [Test error log message]")
}

func TestLogFormattedObject(t *testing.T) {
	logger.DefaultLogger().SetLevel(levels.All)
	content := readConsole(func() {
		person := person{Name: "Michael"}
		logger.DebugF("Person: %v, Car: %v", person, car{Year: "2020"})
		logger.TraceF("Person: %v, Car: %v", person, car{Year: "2021"})
		logger.InfoF("Person: %v, Car: %v", person, car{Year: "2022"})
		logger.WarningF("Person: %v, Car: %v", person, car{Year: "2023"})
		logger.ErrorF("Person: %v, Car: %v", person, car{Year: "2024"})
	})

	assert.Contains(t, content, "Person: {Michael}, Car: {2020}")
	assert.Contains(t, content, "Person: {Michael}, Car: {2021}")
	assert.Contains(t, content, "Person: {Michael}, Car: {2022}")
	assert.Contains(t, content, "Person: {Michael}, Car: {2023}")
	assert.Contains(t, content, "Person: {Michael}, Car: {2024}")
}

func TestLogLevels(t *testing.T) {
	logger.DefaultLogger().SetLevel(levels.Trace)
	content := readConsole(func() {
		logger.Debug("Test debug log message")
		logger.Trace("Test trace log message")
		logger.Info("Test info log message")
		logger.Warning("Test warning log message")
		logger.Error(errors.Errorf("Test error object").Err)
		logger.Error("Test error log message")
	})

	assert.NotContains(t, content, "DEBUG: [Test debug log message]")
	assert.Contains(t, content, "TRACE: [Test trace log message]")
	assert.Contains(t, content, "INFO: [Test info log message]")
	assert.Contains(t, content, "Test error object")
	assert.Contains(t, content, "ERROR: [Test error log message]")
}

func TestStopLog(t *testing.T) {
	content := readConsole(func() {

		logger.DefaultLogger().Stop()
		logger.Info("Test info log message 1")

		logger.DefaultLogger().Start()
		logger.Info("Test info log message 2")

	})

	assert.NotContains(t, content, "INFO: [Test info log message 1]")
	assert.Contains(t, content, "INFO: [Test info log message 2]")
}

func TestAddFileLog(t *testing.T) {
	fileLogger := loggers.NewFileLoggerDefault()
	_, err := logger.RegisterLogger("file", fileLogger)
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Test log info message")

	pattern := fmt.Sprintf("%s*%s", fileLogger.FileOptions.FilePrefix, fileLogger.FileOptions.FileExtension)
	logFiles, err := walkMatch(t, "./", pattern)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range logFiles {
		content := readFileContent(t, f)
		os.Remove(f)
		assert.Contains(t, content, "Test log info message")
	}
}

func TestCustomizedFileLog(t *testing.T) {
	level := levels.Error
	format := "***error:'%s'"
	fileOptions := loggers.FileOptions{
		Directory:     "./",
		FilePrefix:    generateRandomString(5),
		FileExtension: ".txt",
	}

	fileLogger := loggers.NewFileLogger(level, format, fileOptions)
	key := "txt_file"
	_, err := logger.RegisterLogger(key, fileLogger)
	if err != nil {
		t.Fatal(err)
	}
	logger.DefaultLogger().Stop()
	assert.Equal(t, fileLogger.Level, level)
	assert.Equal(t, logger.GetLogger(key).GetLevel(), fileLogger.Level)
	logger.Error("Test log error message")
	logger.Info("Test log info message")

	pattern := fmt.Sprintf("%s*%s", fileOptions.FilePrefix, fileOptions.FileExtension)
	logFiles, err := walkMatch(t, fileOptions.Directory, pattern)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range logFiles {
		content := readFileContent(t, f)
		os.Remove(f)
		assert.Contains(t, content, "***error:'Test log error message'")
		assert.NotContains(t, content, "Test log info message")
	}
}

type logConsole func()

func readConsole(fn logConsole) string {
	currentStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	consoleBytes, _ := io.ReadAll(r)
	content := string(consoleBytes)
	os.Stdout = currentStdout
	return content
}

func walkMatch(t *testing.T, root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func readFileContent(t *testing.T, fName string) string {
	bytesData, err := os.ReadFile(fName)
	if err != nil {
		t.Fatal(err)
	}
	return string(bytesData)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

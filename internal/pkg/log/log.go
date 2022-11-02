package log

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type Logger interface {
	Log(ctx context.Context, level Level, fmtMsg string, args ...interface{})
}

type Level string

const (
	Info    Level = "info"
	Warning Level = "warn"
	Error   Level = "error"
	Fatal   Level = "fatal"
)

func (l Level) String() string {
	switch l {
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		panic(fmt.Sprintf("Unexpected log level: %s", string(l)))
	}
}

type loggerContext struct {
	name   string
	vars   Vars
	logger Logger
}

func New() Logger {
	return &loggerContext{}
}

func (l *loggerContext) Log(ctx context.Context, level Level, fmtMsg string, args ...interface{}) {
	fmtMsg = fmt.Sprintf("%s%s%s", l.name, l.vars.String(), fmtMsg)
	if l.logger == nil {
		log.Println(level.String(), fmt.Sprintf(fmtMsg, args...))
	} else {
		l.logger.Log(ctx, level, fmtMsg, args...)
	}
	if level == Fatal {
		os.Exit(1)
	}
}

type Vars map[string]interface{}

func (vars Vars) String() string {
	if len(vars) == 0 {
		return ""
	}
	varsStrs := make([]string, 0, len(vars))
	for k, v := range vars {
		varsStrs = append(varsStrs, fmt.Sprintf("%s: %v", k, v))
	}
	return fmt.Sprintf("{%s}", strings.Join(varsStrs, ","))
}

func Context(name string, vars Vars, logger Logger) Logger {
	return &loggerContext{
		name:   name,
		vars:   vars,
		logger: logger,
	}
}

func StructVars(s interface{}) Vars {
	var (
		reflectValue = reflect.ValueOf(s)
		reflectType  = reflect.TypeOf(s)
		fieldsNum    = reflectValue.NumField()
	)

	if reflectType.Kind() != reflect.Struct {
		panic("unexpected reflect type kind")
	}

	if reflectType.NumField() != fieldsNum {
		panic("reflect type and value struct fields number mismatch")
	}

	var vars Vars
	if fieldsNum > 0 {
		vars = make(Vars, fieldsNum)
		for i := 0; i < fieldsNum; i++ {
			vars[reflectType.Field(i).Name] = reflectValue.Field(i).Interface()
		}
	}
	return vars
}

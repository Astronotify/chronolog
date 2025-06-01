package internal

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"sort"
	"strings"
	"time"
)

type PrettyConsoleHandler struct {
	writer io.Writer
}

func NewPrettyConsoleHandler(w io.Writer) *PrettyConsoleHandler {
	return &PrettyConsoleHandler{writer: w}
}

func (h *PrettyConsoleHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *PrettyConsoleHandler) Handle(_ context.Context, record slog.Record) error {
	var event any

	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "event" {
			event = attr.Value.Any()
			return false
		}
		return true
	})

	if event == nil {
		return nil
	}

	timestamp := time.Now().UTC().Format(time.RFC3339)
	level := strings.ToUpper(record.Level.String())
	typeName := reflect.TypeOf(event).Name()
	if typeName == "" {
		typeName = "LogEntry"
	}

	line := summarizeAllFields(event)

	fmt.Fprintf(h.writer, "%s\t%s\t%s\t%s\n", timestamp, level, typeName, line)
	return nil
}

func (h *PrettyConsoleHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *PrettyConsoleHandler) WithGroup(_ string) slog.Handler {
	return h
}

func summarizeAllFields(event any) string {
	v := reflect.ValueOf(event)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	// Mapa onde armazenamos todos os pares chave=valor
	fields := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !fieldValue.CanInterface() {
			continue
		}

		// Ignora o campo Context
		if field.Type == reflect.TypeOf((*context.Context)(nil)).Elem() {
			continue
		}

		// Trata struct embutida (LogEntry, por exemplo)
		if field.Anonymous && fieldValue.Kind() == reflect.Struct {
			embedded := summarizeAllFields(fieldValue.Interface())
			for _, kv := range strings.Split(embedded, "  ") {
				parts := strings.SplitN(kv, "=", 2)
				if len(parts) == 2 {
					fields[parts[0]] = parts[1]
				}
			}
			continue
		}

		// Pega o nome do campo ou do json tag se existir
		key := field.Name
		if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			key = strings.Split(jsonTag, ",")[0]
		}

		val := fieldValue.Interface()
		switch v := val.(type) {
		case string:
			if v != "" {
				fields[key] = fmt.Sprintf(`"%s"`, v)
			}
		case time.Time:
			if !v.IsZero() {
				fields[key] = v.Format(time.RFC3339)
			}
		case int, int64, float64, bool:
			fields[key] = fmt.Sprintf("%v", v)
		case map[string]any:
			if len(v) > 0 {
				fields[key] = fmt.Sprintf("%v", v)
			}
		default:
			// Tenta fallback razoável
			fields[key] = fmt.Sprintf("%v", v)
		}
	}

	// Ordena as chaves
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Constrói os pares chave=valor formatados
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, fields[k]))
	}

	return strings.Join(parts, "  ")
}

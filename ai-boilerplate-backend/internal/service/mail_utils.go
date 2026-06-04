package service

import (
	"bytes"
	"encoding/json"
	"html/template"
	"strings"

	"gorm.io/datatypes"
)

func normalizeHelpFeedbackImages(raw string) datatypes.JSON {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return datatypes.JSON([]byte("[]"))
	}
	return datatypes.JSON([]byte(trimmed))
}

func parseMailTemplateParams(raw datatypes.JSON) ([]string, error) {
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" {
		return nil, nil
	}
	var params []string
	if err := json.Unmarshal([]byte(trimmed), &params); err != nil {
		return nil, err
	}
	cleaned := make([]string, 0, len(params))
	seen := make(map[string]struct{}, len(params))
	for _, param := range params {
		param = strings.TrimSpace(param)
		if param == "" {
			continue
		}
		if _, ok := seen[param]; ok {
			continue
		}
		seen[param] = struct{}{}
		cleaned = append(cleaned, param)
	}
	return cleaned, nil
}

func renderMailTemplateString(content string, params map[string]string) (string, error) {
	trimmed := strings.TrimSpace(content)
	if trimmed == "" {
		return "", nil
	}
	tpl, err := template.New("mail").Option("missingkey=zero").Parse(content)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, params); err != nil {
		return "", err
	}
	return buf.String(), nil
}

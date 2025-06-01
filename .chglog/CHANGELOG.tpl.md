{{- /* Title */ -}}
# 📦 Changelog

All notable changes to this project will be documented in this file.

{{- range .Versions }}

---

## 🚀 {{ if .Tag.Previous }}[{{ .Tag.Name }}]({{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}){{ else }}{{ .Tag.Name }}{{ end }} - {{ datetime "2006-01-02" .Tag.Date }}

{{- if .Tag.Subject }}
> {{ .Tag.Subject }}
{{- end }}

{{- range .Sections }}
### {{ .Title }}

{{- range .Entries }}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Description }}
{{- if .References }} ({{ range $index, $ref := .References }}{{ if $index }}, {{ end }}[{{ $ref.Ref }}]({{ $ref.Link }}){{ end }}){{ end }}
{{- end }}

{{ end -}}
{{ end }}

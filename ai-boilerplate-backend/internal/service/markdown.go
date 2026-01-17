package service

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	markdownPolicy = bluemonday.UGCPolicy()
	markdownEngine = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithXHTML()),
	)
)

func renderMarkdownToSafeHTML(markdownContent string) (string, error) {
	var buf bytes.Buffer
	if err := markdownEngine.Convert([]byte(markdownContent), &buf); err != nil {
		return "", err
	}
	return markdownPolicy.Sanitize(buf.String()), nil
}

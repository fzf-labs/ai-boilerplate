package service

import "testing"

func TestNormalizeHelpFeedbackImages(t *testing.T) {
	t.Parallel()

	got := normalizeHelpFeedbackImages(`["a.png","b.png"]`)
	if string(got) != `["a.png","b.png"]` {
		t.Fatalf("unexpected images json: %s", string(got))
	}

	empty := normalizeHelpFeedbackImages("")
	if string(empty) != "[]" {
		t.Fatalf("unexpected empty images json: %s", string(empty))
	}
}

func TestRenderMailTemplateString(t *testing.T) {
	t.Parallel()

	got, err := renderMailTemplateString("Hello {{.name}}, code {{.code}}", map[string]string{
		"name": "Alice",
		"code": "1234",
	})
	if err != nil {
		t.Fatalf("renderMailTemplateString returned error: %v", err)
	}
	if got != "Hello Alice, code 1234" {
		t.Fatalf("unexpected rendered content: %q", got)
	}
}


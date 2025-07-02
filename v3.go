package mytemplate

import (
	"bytes"
	"context"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/a-h/templ"
)

type Component struct {
	ComponentFunc templ.ComponentFunc
}

func (Component) RemoveScriptTags(html string) string {
	re := regexp.MustCompile(`(?is)<script.*?>.*?</script>`)
	return re.ReplaceAllString(html, "")
}

func (Component) RenderHtmlToString(removeWords []string, components ...templ.Component) (string, error) {
	var buf bytes.Buffer
	ctx := context.Background()

	// Render all components to the buffer
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			return "", err
		}
	}

	// Get the rendered HTML as a string
	html := buf.String()

	// Remove specified words
	for _, word := range removeWords {
		html = strings.ReplaceAll(html, word, "")
	}

	// Remove all remaining whitespace (spaces, tabs, newlines)
	//whitespace := regexp.MustCompile(`\s+`)
	// html = whitespace.ReplaceAllString(html, "")

	return html, nil
}

func RemoveScriptTags(html string) string {
	re := regexp.MustCompile(`(?is)<script.*?>.*?</script>`)
	return re.ReplaceAllString(html, "")
}

func RenderHtmlToString(removeWords []string, components ...templ.Component) (string, error) {
	var buf bytes.Buffer
	ctx := context.Background()

	// Render all components to the buffer
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			return "", err
		}
	}

	// Get the rendered HTML as a string
	html := buf.String()

	// Remove specified words
	for _, word := range removeWords {
		html = strings.ReplaceAll(html, word, "")
	}

	// Remove all remaining whitespace (spaces, tabs, newlines)
	//whitespace := regexp.MustCompile(`\s+`)
	// html = whitespace.ReplaceAllString(html, "")

	return html, nil
}

func GenerateComponent(component templ.Component) templ.ComponentFunc {
	var mycomponent Component

	str, err := mycomponent.RenderHtmlToString(nil, component)
	if err != nil {
		panic(err)
	}

	str = mycomponent.RemoveScriptTags(str)

	mycomponent.ComponentFunc = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, str)
		return err
	})

	return mycomponent.ComponentFunc
}

func GenerateJS(components ...templ.ComponentScript) error {
	// Generate JS with newlines between functions
	var jsBuf bytes.Buffer
	for i, script := range components {
		// Render the components to a string with word removal
		jsStr := script.Function

		if i > 0 {
			jsBuf.WriteString("\n") // Add newline before second and subsequent functions
		}
		jsBuf.WriteString(jsStr)
	}

	// Write the string to output.html
	if err := os.WriteFile("index.js", jsBuf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

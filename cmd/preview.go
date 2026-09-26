/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	alert "github.com/yuin/goldmark-alert"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

var htmlTemplate = `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>mdprv</title>
<style>
/*
  Josh's Custom CSS Reset
  https://www.joshwcomeau.com/css/custom-css-reset/
*/
/* 1. Use a more-intuitive box-sizing model */
*, *::before, *::after {
  box-sizing: border-box;
}

/* 2. Remove default margin */
*:not(dialog) {
  margin: 0;
}

/* 3. Enable keyword animations */
@media (prefers-reduced-motion: no-preference) {
  html {
    interpolate-size: allow-keywords;
  }
}

body {
  /* 4. Increase line-height */
  line-height: 1.5;
  /* 5. Improve text rendering */
  -webkit-font-smoothing: antialiased;
}

/* 6. Improve media defaults */
img, picture, video, canvas, svg {
  display: block;
  max-width: 100%%;
}

/* 7. Inherit fonts for form controls */
input, button, textarea, select {
  font: inherit;
}

/* 8. Avoid text overflows */
p, h1, h2, h3, h4, h5, h6 {
  overflow-wrap: break-word;
}

/* 9. Improve line wrapping */
p {
  text-wrap: pretty;
}
h1, h2, h3, h4, h5, h6 {
  text-wrap: balance;
}

/*
  10. Create a root stacking context
*/
#root, #__next {
  isolation: isolate;
}

/*
  11. Custom CSS
*/
main {
  color: #333;
  background-color: #f1f1f1;
  font-family: "Noto Sans JP", sans-serif;
  font-optical-sizing: auto;
  font-weight: 400;
  font-style: normal;
  font-size: 15px;
}

article {
  background-color: #fff;
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 100px;
  box-shadow: 0px 0px 4px rgba(51, 51, 51, 0.2);
  line-height: 2;
}

h1, h2, h3, h4, h5, h6 {
  font-weight: 500;
  line-height: 1.1;
}

h1 {
  font-size: 40px;
  margin-top: 4px;
  margin-bottom: 10px;
  padding-bottom: 50px;
  letter-spacing: normal;
  text-align: center;
  border-bottom: 2px solid #eee;
}

h2 {
  margin-top: 30px;
  margin-bottom: 20px;
  padding: 10px 0 10px 25px;
  font-size: 24px;
  font-weight: bold;
  border-bottom: none;
  border-left: 6px solid #4979f5;
}

h3 {
  margin: 30px 0 20px 6px;
  font-size: 20px;
  font-weight: bold;
}

h4 {
  margin-left: 6px;
  margin-bottom: 20px;
  padding-top: 20px;
  font-size: 16px;
  font-weight: bold;
  text-decoration: underline;
}

h5 {
  margin-bottom: 20px;
  padding-top: 20px;
  font-size: 15px;
  font-weight: bold;
}

h6 {
  font-size: 12px;
  padding-bottom: 10px;
}

p {
  margin: 0 15px 15px 15px;
  font-size: 16px;
  line-height: 180%%;
  letter-spacing: .05em;
}

strong {
  font-weight: bold;
}

ul, ol {
  margin: .5em 0;
  padding: 0 15px 0 2em;
}

ul li, ol li {
  margin: 0 15px .5em 15px;
  letter-spacing: .05em;
  line-height: 180%%;
  font-size: 16px;
}

a {
  color: #0055ad;
}

img {
  max-width: 100%%;
  max-height: 650px;
  margin-bottom: 5px;
  border: none;
  font-size: 14px;
}

code, pre, kbd {
  font-family: "M PLUS 1 Code", monospace;
  font-optical-sizing: auto;
  font-weight: 400;
  font-style: normal;
}

code {
  padding: 2px 4px;
  font-size: 90%%;
  color: #c7254e;
  background-color: #f9f2f4;
  border-radius: 4px;
}

pre {
  display: block;
  margin: 0 15px 15px 15px;
  overflow: auto;
  font-size: 16px;
  line-height: 180%%;
  letter-spacing: 0;
  word-break: break-all;
  overflow-wrap: break-word;
  border: solid 1px #ccc;
  border-radius: 4px;
}

pre code {
  font-size: initial;
}

kbd {
  padding: 2px 4px;
  font-size: 90%%;
  color: #fff;
  background-color: #333;
  border-radius: 4px;
}

blockquote {
  margin: 1em 0 1em 15px;
  padding: 7px;
  color: #666;
  border-left: 7px solid #eee;
}

blockquote p {
  margin: 7px 15px;
  font-size: 16px;
  line-height: 180%%;
  letter-spacing: .05em;
}

hr {
  margin: 0;
  padding: 0;
  height: 2px;
  background-color: #eee;
  border: none;
}

table {
  margin-bottom: 1rem;
  vertical-align: top;
  border-spacing: 0;
}

th, td {
  border-bottom: 1px solid #eee;
  padding: 1rem;
}

details p {
  font-size: 15px;
}

div.markdown-alert {
  margin: 15px 0;
  padding: 15px 15px 2px 2px;
  border-radius: 4px;
}

div.markdown-alert p.markdown-alert-title {
  font-weight: bold;
}

div.markdown-alert p.markdown-alert-title svg {
  display: inline-block;
  margin-right: 4px;
  vertical-align: middle;
}

div.markdown-alert-note {
  color: #055160;
  background-color: #cff4fc;
}

div.markdown-alert-note svg {
  fill: #055160;
}

div.markdown-alert-tip {
  color: #0a3622;
  background-color: #d1e7dd;
}

div.markdown-alert-tip svg {
  fill: #0a3622;
}

div.markdown-alert-important {
  color: #052c65;
  background-color: #cfe2ff;
}

div.markdown-alert-important svg {
  fill: #052c65;
}

div.markdown-alert-warning {
  color: #664d03;
  background-color: #fff3cd;
}

div.markdown-alert-warning svg {
  fill: #664d03;
}

div.markdown-alert-caution {
  color: #58151c;
  background-color: #f8d7da;
}

div.markdown-alert-caution svg {
  fill: #58151c;
}

@media print {
  article {
    padding: 0;
    box-shadow: none;
  }
}
</style>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.12.0/styles/foundation.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.12.0/highlight.min.js"></script>
</head>
<body>
<main>
<article>
%s
</article>
</main>
<script>hljs.highlightAll();</script>
</body>
</html>
`

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview [file]",
	Short: "Preview a Markdown file in your browser",
	Long:  `Render a Markdown file as HTML and start a local HTTP server to preview it.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runPreview,
}

func runPreview(cmd *cobra.Command, args []string) error {
	filePath := args[0]
	log.Printf("Create a preview for %s\n", filePath)

	source, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var buf bytes.Buffer
	mdParser := parser.New(parser.WithExtensions(extension.GFMParser, extension.FootnoteParser, alert.Parser))
	mdRenderer := html.New(html.WithUnsafe(), html.WithExtensions(extension.GFMHTMLRenderer, extension.FootnoteHTMLRenderer, alert.HTMLRenderer))

	doc := mdParser.Parse(source)
	if err := mdRenderer.Render(&buf, source, doc); err != nil {
		return fmt.Errorf("failed to render markdown: %w", err)
	}

	renderedHTML := fmt.Sprintf(htmlTemplate, buf.String())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "%s\n", renderedHTML)
	})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("failed to start listener: %w", err)
	}
	defer listener.Close()

	log.Printf("Starting the server: http://%s", listener.Addr().String())
	return http.Serve(listener, mux)
}

func init() {
	rootCmd.AddCommand(previewCmd)
}

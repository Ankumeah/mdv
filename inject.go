package main

import (
	"strings"
)

const tag = "<head>"
const injectedCSS = `
<style>
  * {
    color: white;
    background-color: #171717;
  }

  pre, code {
    background-color: #1f1f1f
  }
</style>
`
const injectedJS = `
<script>
  function main() {
    for (const link of document.getElementsByTagName("a")) {
      if (link.href.startsWith("data:")) {
        link.addEventListener(
          "click",
          (e) => {
            e.preventDefault();

            const base64 = link.href.slice(link.href.indexOf(",") + 1);
            const bytes = Uint8Array.from(atob(base64), char => char.charCodeAt(0));
            const text = new TextDecoder("utf-8").decode(bytes);

            window.setHTML(text);
          }
        );
      }
    }
  }

  document.addEventListener(
    "DOMContentLoaded",
    main,
  );
</script>
`

func inject(html string) string {
	idx := strings.Index(html, tag) + len(tag)

	return html[:idx] + injectedCSS + injectedJS + html[idx+1:]
}

func styleText(text string) string {
	return injectedCSS + "<pre>\n" + text + "\n</pre>"
}

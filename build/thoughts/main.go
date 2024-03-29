/*******************************************************************************
*
* Copyright 2016-2018 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under the
* terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	//list posts
	posts := allPosts()
	sort.Sort(posts) //by creation timestamp

	//deduplicate slugs (note that slugs for each specific post are constant over
	//time because of how we just sorted)
	slugSeen := make(map[string]bool)
	for _, post := range posts {
		if slugSeen[post.Slug] {
			//deduplicate "$slug" to "$slug-1", "$slug-2" etc.
			i := 0
			for {
				i++
				altSlug := fmt.Sprintf("%s-%d", post.Slug, i)
				if !slugSeen[altSlug] {
					post.Slug = altSlug
					break
				}
			}
		}
		slugSeen[post.Slug] = true
		continue
	}

	//render posts
	for _, p := range posts {
		p.Render()
	}

	//index.html and sitemap.html show posts in reverse order
	reverse(posts)
	RenderIndex(posts)
	RenderAll(posts)
	RenderRSS(posts)
}

////////////////////////////////////////////////////////////////////////////////
// output formatting

//RenderIndex generates the index.html page.
func RenderIndex(posts []*Post) {
	//not more than 10 posts
	if len(posts) > 10 {
		posts = posts[:10]
	}

	//accumulate posts
	articlesStr := ""
	if len(posts) > 0 {
		articles := make([]string, 0, len(posts))
		for _, post := range posts {
			//shorten post.HTML if it contains multiple headings
			htmlStr := post.ShortenedHTML()
			//include permalink in initial heading
			htmlStr = initialHeadingRx.ReplaceAllStringFunc(htmlStr, func(h1str string) string {
				match := initialHeadingRx.FindStringSubmatch(h1str)
				return fmt.Sprintf("<h1><a href=\"%s\" title=\"Permalink\">[l]</a> %s</h1>",
					post.OutputURL(), match[1],
				)
			})
			articles = append(articles, htmlStr)
		}
		articlesStr = "<article>" + strings.Join(articles, "</article><article>") + "</article>"
	}

	metadata := map[string]string{
		"og:title": PageName,
		"og:type":  "website",
		"og:url":   "/thoughts/",
	}
	writeFile("thoughts/index.html", "", articlesStr, metadata)
}

//RenderAll generates the sitemap.html page.
func RenderAll(posts []*Post) {
	items := ""
	currentMonth := ""

	for _, post := range posts {
		//add a month header when this post is from a different month than the previous one
		month := post.CreationTime().Format("Jan 2006")
		if month != currentMonth {
			items += fmt.Sprintf("</ul><h2>%s</h2><ul>", month)
			currentMonth = month
		}
		//show either the initial <h1> or fall back to the slug
		items += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", post.OutputURL(), post.Title())
	}

	items = strings.TrimPrefix(items, "</ul>")
	writeFile("thoughts/sitemap.html", "Article list",
		"<main class=\"sitemap\">"+items+"</ul></main>",
		map[string]string{},
	)
}

//RenderRSS generates the rss.xml document.
func RenderRSS(posts []*Post) {
	//not more than 10 posts
	if len(posts) > 10 {
		posts = posts[:10]
	}

	var lines []string
	addLine := func(line string, args ...interface{}) {
		if len(args) > 0 {
			line = fmt.Sprintf(line, args...)
		}
		lines = append(lines, line)
	}

	addLine(`<?xml version="1.0"?>`)
	addLine(`<rss version="2.0"><channel>`)
	addLine(`  <title>%s</title>`, PageName)
	addLine(`  <link>%s/thoughts/</link>`, TargetURL)
	addLine(`  <description>%s</description>`, PageDescription)
	addLine(`  <language>en</language>`)
	addLine(`  <lastBuildDate>%s</lastBuildDate>`, time.Now().UTC().Format(time.RFC1123Z))
	for _, post := range posts {
		addLine(`  <item>`)
		addLine(`    <title>%s</title>`, post.Title())
		addLine(`    <description>%s</description>`, escapeHTML(post.ShortenedHTML()))
		addLine(`    <link>%s/thoughts/posts/%s.html</link>`, TargetURL, post.Slug)
		addLine(`    <guid>%s/id/thoughts/%s</guid>`, TargetURL, post.Slug)
		addLine(`    <pubDate>%s</pubDate>`, post.CreationTime().Format(time.RFC1123Z))
		addLine(`  </item>`)
	}
	addLine("</channel></rss>\n")

	FailOnErr(ioutil.WriteFile("thoughts/rss.xml", []byte(strings.Join(lines, "\n")), 0644))
}

func escapeHTML(s string) string {
	s = strings.Replace(s, "&", "&amp;", -1)
	s = strings.Replace(s, "'", "&#39;", -1)
	s = strings.Replace(s, `"`, "&quot;", -1)
	s = strings.Replace(s, "<", "&lt;", -1)
	return strings.Replace(s, ">", "&gt;", -1)
}

////////////////////////////////////////////////////////////////////////////////
// utilities

//FailOnErr complains and aborts if `err != nil`.
func FailOnErr(err error) {
	if err != nil {
		os.Stderr.Write([]byte(err.Error() + "\n"))
		os.Exit(1)
	}
}

func reverse(list []*Post) {
	max := len(list) - 1
	cnt := len(list) / 2
	for idx := 0; idx < cnt; idx++ {
		list[idx], list[max-idx] = list[max-idx], list[idx]
	}
}

func writeFile(path, title, contents string, metadata map[string]string) {
	str := Config.TemplateHTML

	slashCount := strings.Count(path, "/")
	dotdots := make([]string, 0, slashCount)
	for idx := 0; idx < slashCount; idx++ {
		dotdots = append(dotdots, "..")
	}
	if len(dotdots) == 0 {
		dotdots = []string{"."}
	}
	str = strings.Replace(str, "%PATH_TO_ROOT%", strings.Join(dotdots, "/"), -1)

	if title == "" {
		str = strings.Replace(str, "%TITLE%", PageName, -1)
	} else {
		str = strings.Replace(str, "%TITLE%", title+" &ndash; "+PageName, -1)
	}
	str = strings.Replace(str, "%CONTENT%", contents, -1)

	//sort metadata keys deterministically to avoid diff noise
	metadataKeys := make([]string, 0, len(metadata))
	for key := range metadata {
		metadataKeys = append(metadataKeys, key)
	}
	sort.Strings(metadataKeys)

	//format metadata according to http://ogp.me/
	var metadataStr string
	for _, key := range metadataKeys {
		metadataStr += fmt.Sprintf(
			`<meta property="%s" content="%s" />`,
			template.HTMLEscapeString(key), template.HTMLEscapeString(metadata[key]),
		)
	}
	str = strings.Replace(str, "%META%", metadataStr, -1)

	FailOnErr(ioutil.WriteFile(path, []byte(str), 0644))
}

/*******************************************************************************
*
* Copyright 2018 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
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
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-commonmark/markdown"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	inputBytes, err := ioutil.ReadFile("build/noises/data.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	var data struct {
		Shows            map[string]*show `yaml:"shows"`
		Files            []*file          `yaml:"files"`
		CurrentShowID    string           `yaml:"-"`
		CurrentFileIndex int              `yaml:"-"`
		XMLIntro         template.HTML    `yaml:"-"`
	}
	//this needs to be interpolated by the template, otherwise html/template
	//breaks the <? opener
	data.XMLIntro = `<?xml version="1.0" encoding="UTF-8"?>`

	//load input data
	err = yaml.Unmarshal(inputBytes, &data)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, show := range data.Shows {
		show.CompileMarkdown()
	}
	for _, file := range data.Files {
		file.CompileMarkdown()
		file.FindDownloads()
	}

	templateFuncs := template.FuncMap{
		"reverseFiles":           reverseFiles,
		"readableFileSize":       readableFileSize,
		"readableLengthSeconds":  readableLengthSeconds,
		"unixTimeToRFC1123":      unixTimeToRFC1123,
		"unixTimeToReadableDate": unixTimeToReadableDate,
	}

	//render noises/index.html
	tmplBytes, err := ioutil.ReadFile("build/noises/index.html.tpl")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmpl, err := template.New("noises/index.html").Funcs(templateFuncs).Parse(string(tmplBytes))
	if err != nil {
		log.Fatal(err.Error())
	}
	out, err := os.Create("noises/index.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer out.Close()
	err = tmpl.Execute(out, data)
	if err != nil {
		log.Fatal(err.Error())
	}

	//render pages for individual shows
	tmplBytes, err = ioutil.ReadFile("build/noises/show.html.tpl")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmpl, err = template.New("noises/show.html").Funcs(templateFuncs).Parse(string(tmplBytes))
	if err != nil {
		log.Fatal(err.Error())
	}
	for showID, show := range data.Shows {
		if show.IsExternal {
			continue
		}
		dirPath := filepath.Join("noises", showID)
		err = os.MkdirAll(dirPath, 0777)
		if err != nil {
			log.Fatal(err.Error())
		}
		out, err := os.Create(dirPath + "/index.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer out.Close()
		data.CurrentShowID = showID
		err = tmpl.Execute(out, data)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	//render pages for individual episodes
	tmplBytes, err = ioutil.ReadFile("build/noises/episode.html.tpl")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmpl, err = template.New("noises/episode.html").Funcs(templateFuncs).Parse(string(tmplBytes))
	if err != nil {
		log.Fatal(err.Error())
	}
	for idx, file := range data.Files {
		show := data.Shows[file.ShowID]
		if show.IsExternal || file.Episode == 0 {
			continue
		}
		dirPath := filepath.Join("noises", file.ShowID, file.Slug)
		err = os.MkdirAll(dirPath, 0777)
		if err != nil {
			log.Fatal(err.Error())
		}
		out, err := os.Create(dirPath + "/index.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer out.Close()
		data.CurrentFileIndex = idx
		err = tmpl.Execute(out, data)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	//render RSS feeds
	tmplBytes, err = ioutil.ReadFile("build/noises/rss.xml.tpl")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmpl, err = template.New("noises/rss.xml").Funcs(templateFuncs).Parse(string(tmplBytes))
	if err != nil {
		log.Fatal(err.Error())
	}
	for showID, show := range data.Shows {
		if show.FeedPath != "" {
			data.CurrentShowID = showID
			out, err := os.Create(show.FeedPath)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer out.Close()
			err = tmpl.Execute(out, data)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// data structures

type show struct {
	Title           string        `yaml:"title"`
	Subtitle        string        `yaml:"subtitle"`
	Description     string        `yaml:"description"`
	DescriptionHTML template.HTML `yaml:"-"`
	Category        string        `yaml:"category"` //iTunes podcast category, e.g. "Technology"
	FeedPath        string        `yaml:"feed"`
	URL             string        `yaml:"href"`
	IsExternal      bool          `yaml:"external"`

	Covers struct {
		ForFeed string `yaml:"feed"`
		ForHTML string `yaml:"html"`
	} `yaml:"covers"`

	Subscriptions []struct {
		Name string `yaml:"name"`
		URL  string `yaml:"href"`
	}
}

type file struct {
	ShowID   string `yaml:"show"`
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
	Episode  uint   `yaml:"episode"`
	Slug     string `yaml:"slug"`

	PublicationTimeUnix uint64 `yaml:"pubtime"`

	Markdown struct {
		Description string   `yaml:"description"`
		Notes       string   `yaml:"notes"`
		Sources     []string `yaml:"sources"`
	} `yaml:"markdown"`

	HTML struct {
		Description template.HTML
		Notes       template.HTML
		Sources     []template.HTML
	} `yaml:"-"`

	Music []struct {
		Artist  string `yaml:"artist"`
		Title   string `yaml:"title"`
		URL     string `yaml:"href"`
		Variant string `yaml:"variant"` //optional
	} `yaml:"music"`

	Downloads []struct {
		Format    string
		MIMEType  string
		FileName  string
		SizeBytes uint64
	} `yaml:"-"`
	LengthSeconds uint `yaml:"-"`
}

func compileMarkdown(input string) template.HTML {
	md := markdown.New(
		markdown.HTML(true),
		markdown.Typographer(false),
	)
	if input == "" {
		return ""
	}
	out := strings.TrimSpace(md.RenderToString([]byte(input)))
	out = strings.TrimPrefix(out, "<p>")
	out = strings.TrimSuffix(out, "</p>")
	return template.HTML(out)
}

func (s *show) CompileMarkdown() {
	s.DescriptionHTML = compileMarkdown(s.Description)
}

func (f *file) CompileMarkdown() {
	f.HTML.Description = compileMarkdown(f.Markdown.Description)
	f.HTML.Notes = compileMarkdown(f.Markdown.Notes)
	f.HTML.Sources = make([]template.HTML, len(f.Markdown.Sources))
	for idx, str := range f.Markdown.Sources {
		f.HTML.Sources[idx] = compileMarkdown(str)
	}
}

func (f *file) FindDownloads() {
	if f.Slug == "" {
		f.Downloads = nil
		f.LengthSeconds = 0
		return
	}

	formats := []struct {
		Extension string
		Name      string
		MIMEType  string
	}{
		{"mp3", "MP3", "audio/mpeg"},
		{"ogg", "Ogg Vorbis", "audio/ogg"},
		{"opus", "Opus", "audio/opus"},
		{"flac", "FLAC", "audio/flac"},
	}

	hasFLAC := false
	for _, fmt := range formats {
		fileName := f.ShowID + "-" + f.Slug + "." + fmt.Extension
		fi, err := os.Stat("dl/" + fileName)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(err.Error())
		}

		if fmt.Extension == "flac" {
			hasFLAC = true
		}
		f.Downloads = append(f.Downloads,
			struct {
				Format    string
				MIMEType  string
				FileName  string
				SizeBytes uint64
			}{
				Format:    fmt.Name,
				MIMEType:  fmt.MIMEType,
				FileName:  fileName,
				SizeBytes: uint64(fi.Size()),
			},
		)
	}

	if hasFLAC {
		//file(1) reports enough metadata from FLAC files to calculate the duration
		//of the audio file from
		path := "dl/" + f.ShowID + "-" + f.Slug + ".flac"
		cmd := exec.Command("file", path)
		cmd.Stderr = os.Stderr
		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err.Error())
		}
		match := regexp.MustCompile(`([0-9.]+) kHz, ([0-9]+) samples`).FindStringSubmatch(string(output))
		if match == nil {
			log.Fatalf("unexpected output from `file %s`: %q\n", path, string(output))
		}
		sampleRateKHz, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		sampleCount, err := strconv.ParseUint(match[2], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		f.LengthSeconds = uint(float64(sampleCount) / (sampleRateKHz * 1000))
	}
}

////////////////////////////////////////////////////////////////////////////////
// template helper functions

func reverseFiles(in []*file) (out []*file) {
	out = make([]*file, len(in))
	for idx, f := range in {
		out[len(in)-idx-1] = f
	}
	return
}

func readableFileSize(sizeBytes uint64) string {
	out := fmt.Sprintf("%.1f MiB", float64(sizeBytes)/(1<<20))
	return strings.Replace(out, ".", ",", -1)
}

func readableLengthSeconds(lengthSeconds uint) string {
	return fmt.Sprintf("%02d:%02d", lengthSeconds/60, lengthSeconds%60)
}

func unixTimeToRFC1123(in uint64) string {
	gmt := time.FixedZone("GMT", 0)
	return time.Unix(int64(in), 0).In(gmt).Format(time.RFC1123)
}

var germanMonths = map[time.Month]string{
	time.January:   "Januar",
	time.February:  "Februar",
	time.March:     "März",
	time.April:     "April",
	time.May:       "Mai",
	time.June:      "Juni",
	time.July:      "Juli",
	time.August:    "August",
	time.September: "September",
	time.October:   "Oktober",
	time.November:  "November",
	time.December:  "Dezember",
}

func unixTimeToReadableDate(in uint64) string {
	t := time.Unix(int64(in), 0).In(time.FixedZone("GMT", 0))
	str := t.Format("2. XXX 2006")
	return strings.Replace(str, "XXX", germanMonths[t.Month()], -1)
}

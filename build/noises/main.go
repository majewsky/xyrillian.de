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
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-commonmark/markdown"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	inputBytes, err := ioutil.ReadFile("build/noises/data.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	//load input data
	var data struct {
		Shows map[string]show `yaml:"shows"`
		Files []*file         `yaml:"files"`
	}
	err = yaml.Unmarshal(inputBytes, &data)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, file := range data.Files {
		file.CompileMarkdown()
		file.FindDownloads()
	}

	templateFuncs := template.FuncMap{
		"reverseFiles":          reverseFiles,
		"readableFileSize":      readableFileSize,
		"readableLengthSeconds": readableLengthSeconds,
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
}

////////////////////////////////////////////////////////////////////////////////
// data structures

type show struct {
	Title       string `yaml:"title"`
	Subtitle    string `yaml:"subtitle"`
	Description string `yaml:"description"`
	Category    string `yaml:"category"` //iTunes podcast category, e.g. "Technology"
	FeedPath    string `yaml:"feed"`
	URL         string `yaml:"href"`
	IsExternal  bool   `yaml:"external"`

	Covers struct {
		ForFeed string `yaml:"feed"`
		ForHTML string `yaml:"html"`
	} `yaml:"covers"`
}

type file struct {
	ShowID   string `yaml:"show"`
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
	Episode  uint   `yaml:"episode"`
	Slug     string `yaml:"slug"`

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
		FileName  string
		SizeBytes uint64
	} `yaml:"-"`
	LengthSeconds uint `yaml:"-"`
}

func (f *file) CompileMarkdown() {
	md := markdown.New(
		markdown.HTML(true),
		markdown.Typographer(false),
	)
	compile := func(input string) template.HTML {
		if input == "" {
			return ""
		}
		out := strings.TrimSpace(md.RenderToString([]byte(input)))
		out = strings.TrimPrefix(out, "<p>")
		out = strings.TrimSuffix(out, "</p>")
		return template.HTML(out)
	}

	f.HTML.Description = compile(f.Markdown.Description)
	f.HTML.Notes = compile(f.Markdown.Notes)
	f.HTML.Sources = make([]template.HTML, len(f.Markdown.Sources))
	for idx, str := range f.Markdown.Sources {
		f.HTML.Sources[idx] = compile(str)
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
	}{
		{"mp3", "MP3"},
		{"ogg", "Ogg Vorbis"},
		{"opus", "Opus"},
		{"flac", "FLAC"},
	}

	hasFLAC := false
	for _, fmt := range formats {
		fileName := f.Slug + "." + fmt.Extension
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
				FileName  string
				SizeBytes uint64
			}{
				Format:    fmt.Name,
				FileName:  fileName,
				SizeBytes: uint64(fi.Size()),
			},
		)
	}

	if hasFLAC {
		//file(1) reports enough metadata from FLAC files to calculate the duration
		//of the audio file from
		path := "dl/" + f.Slug + ".flac"
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

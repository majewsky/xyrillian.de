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
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "github.com/majewsky/xyrillian.de/build/util" // nolint:staticcheck
)

func main() {
	_ = MustReturn(os.Stat("./dl/"))

	//load metadata cache
	var cache audioMetadataCache
	buf := MustReturn(os.ReadFile("build/noises/cache.json"))
	MustSucceed(json.Unmarshal(buf, &cache))

	var data struct {
		FeaturedShows    []string         `json:"featuredShows"`
		Shows            map[string]*show `json:"shows"`
		Files            []*file          `json:"files"`
		CurrentShowID    string           `json:"-"`
		CurrentFileIndex int              `json:"-"`
		XMLIntro         template.HTML    `json:"-"`
	}
	//this needs to be interpolated by the template, otherwise html/template
	//breaks the <? opener
	data.XMLIntro = `<?xml version="1.0" encoding="UTF-8"?>`

	//load input data
	buf = MustReturn(os.ReadFile("build/noises/data.yaml"))
	MustSucceed(UnmarshalYAMLviaJSON(buf, &data))
	for showID, show := range data.Shows {
		show.ComputeCopyright(showID, data.Files)
		show.CompileMarkdown()
	}
	for _, file := range data.Files {
		file.CompileMarkdown()
		file.AudioMetadata = file.GetAudioMetadata(cache)
		if file.Episode != nil {
			file.EpisodeAsInt = *file.Episode
		}
	}

	//update metadata cache
	cache = audioMetadataCache{make(map[string]audioMetadata)}
	for _, f := range data.Files {
		if f.Slug != "" {
			cache.AudioMetadata[f.CacheKey()] = f.AudioMetadata
		}
	}
	buf = MustReturn(json.MarshalIndent(cache, "", "  "))
	MustSucceed(os.WriteFile("build/noises/cache.json", append(buf, '\n'), 0666))

	//render noises/index.html
	render := buildTemplateRenderer("build/noises/index.html.tpl")
	render("noises/index.html", data)

	//render pages for individual shows
	render = buildTemplateRenderer("build/noises/show.html.tpl")
	for showID, show := range data.Shows {
		if show.ExternalURL != "" {
			continue
		}
		dirPath := filepath.Join("noises", showID)
		MustSucceed(os.MkdirAll(dirPath, 0777))
		data.CurrentShowID = showID
		render(dirPath+"/index.html", data)
	}

	//render pages for individual episodes
	render = buildTemplateRenderer("build/noises/episode.html.tpl")
	for idx, file := range data.Files {
		show := data.Shows[file.ShowID]
		if show.ExternalURL != "" || file.Episode == nil {
			continue
		}
		dirPath := filepath.Join("noises", file.ShowID, file.Slug)
		MustSucceed(os.MkdirAll(dirPath, 0777))
		data.CurrentFileIndex = idx
		render(dirPath+"/index.html", data)
	}

	//render RSS feeds
	render = buildTemplateRenderer("build/noises/rss.xml.tpl")
	for showID, show := range data.Shows {
		if show.FeedPath != "" {
			data.CurrentShowID = showID
			render(show.FeedPath, data)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// data structures

type show struct {
	Title           string        `json:"title"`
	Subtitle        string        `json:"subtitle"`
	Description     string        `json:"description"`
	DescriptionHTML template.HTML `json:"-"`
	ExtraAuthor     string        `json:"extraAuthor"`
	Category        string        `json:"category"` //iTunes podcast category, e.g. "Technology"
	FeedPath        string        `json:"feed"`
	URL             string        `json:"href"`
	ExternalURL     string        `json:"external"`

	FeedConfig struct {
		EpisodeNumberInTitle bool   `json:"episodeNumberInTitle"`
		HumanReadableShowID  string `json:"humanReadableShowID"`
	} `json:"feedConfig"`

	Covers struct {
		ForFeed string `json:"feed"`
		ForHTML string `json:"html"`
	} `json:"covers"`

	Subscriptions []struct {
		Name string `json:"name"`
		URL  string `json:"href"`
	}

	Copyright struct {
		MinYear int
		MaxYear int
	} `json:"-"`
}

type file struct {
	ShowID       string `json:"show"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Episode      *uint  `json:"episode"`
	EpisodeAsInt uint   `json:"-"`
	Slug         string `json:"slug"`

	PublicationTimeUnix uint64 `json:"pubtime"`
	LegacyGUID          bool   `json:"legacyGUID"`

	Markdown struct {
		Description string   `json:"description"`
		Notes       string   `json:"notes"`
		Sources     []string `json:"sources"`
		ShowNotes   bool     `json:"shownotes"`
	} `json:"markdown"`

	HTML struct {
		Description template.HTML
		Notes       template.HTML
		Sources     []template.HTML
		ShowNotes   template.HTML
	} `json:"-"`

	RSS struct {
		Description string
		ShowNotes   string
	} `json:"-"`

	Music []struct {
		Artist  string `json:"artist"`
		Title   string `json:"title"`
		URL     string `json:"href"`
		Variant string `json:"variant"` //optional
	} `json:"music"`

	AudioMetadata audioMetadata `json:"-"`
}

type download struct {
	Format    string `json:"format"`
	MIMEType  string `json:"mime_type"`
	FileName  string `json:"filename"`
	SizeBytes uint64 `json:"size_bytes"`
}

type chapter struct {
	StartSeconds float64 `json:"start_secs"`
	EndSeconds   float64 `json:"end_secs"`
	Title        string  `json:"title"`
	URL          string  `json:"url,omitempty"`
}

type audioMetadata struct {
	LengthSeconds uint       `json:"length_secs"`
	Downloads     []download `json:"downloads"`
	Chapters      []chapter  `json:"chapters"`
}

type audioMetadataCache struct {
	AudioMetadata map[string]audioMetadata `json:"audio_metadata"`
}

func compileMarkdown(input string) template.HTML {
	if input == "" {
		return ""
	}
	out := strings.TrimSpace(MustReturn(RenderMarkdown([]byte(input))))
	out = strings.TrimPrefix(out, "<p>")
	out = strings.TrimSuffix(out, "</p>")
	return template.HTML(out)
}

func (s *show) ComputeCopyright(showID string, allFiles []*file) {
	min := math.MaxInt32
	max := math.MinInt32
	for _, f := range allFiles {
		if showID != f.ShowID {
			continue
		}
		year := time.Unix(int64(f.PublicationTimeUnix), 0).Year()
		if min > year {
			min = year
		}
		if max < year {
			max = year
		}
	}
	if min > max { //no files for this show
		currentYear := time.Now().Year()
		s.Copyright.MinYear = currentYear
		s.Copyright.MaxYear = currentYear
	} else {
		s.Copyright.MinYear = min
		s.Copyright.MaxYear = max
	}
}

func (s *show) CompileMarkdown() {
	s.DescriptionHTML = compileMarkdown(s.Description)
}

func (f file) CacheKey() string {
	return fmt.Sprintf("%s:%s", f.ShowID, f.Slug)
}

func (f *file) CompileMarkdown() {
	f.HTML.Description = compileMarkdown(f.Markdown.Description)
	f.HTML.Notes = compileMarkdown(f.Markdown.Notes)

	f.HTML.Sources = make([]template.HTML, len(f.Markdown.Sources))
	rssSources := make([]string, len(f.Markdown.Sources))
	for idx, str := range f.Markdown.Sources {
		f.HTML.Sources[idx] = compileMarkdown(str)
		rssSources[idx] = string(f.HTML.Sources[idx])
	}

	f.RSS.Description = string(f.HTML.Description)
	f.RSS.ShowNotes = fmt.Sprintf(`<p>%s</p>`, f.HTML.Description)
	if len(rssSources) > 0 {
		f.RSS.ShowNotes += fmt.Sprintf(`<h2>Quellen und Links</h2><ul><li>%s</li></ul>`,
			strings.Join(rssSources, `</li><li>`),
		)
	}

	if f.Markdown.ShowNotes {
		showNotesPath := fmt.Sprintf("build/noises/shownotes/%s/%s.md", f.ShowID, f.Slug)
		f.HTML.ShowNotes = compileMarkdown(string(MustReturn(os.ReadFile(showNotesPath))))
		f.RSS.ShowNotes += `<h2>Shownotes</h2>` + string(f.HTML.ShowNotes)
	}
}

func (f *file) GetAudioMetadata(cache audioMetadataCache) (result audioMetadata) {
	if f.Slug == "" {
		return audioMetadata{}
	}
	if m, exists := cache.AudioMetadata[f.CacheKey()]; exists {
		return m
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

	firstOggPath := ""
	for _, fmt := range formats {
		fileName := f.ShowID + "-" + f.Slug + "." + fmt.Extension
		path := "dl/" + fileName
		fi, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(err.Error())
		}

		if (fmt.Extension == "ogg" || fmt.Extension == "opus") && firstOggPath == "" {
			firstOggPath = path
		}

		result.Downloads = append(result.Downloads, download{
			Format:    fmt.Name,
			MIMEType:  fmt.MIMEType,
			FileName:  fileName,
			SizeBytes: uint64(fi.Size()),
		})
	}

	if firstOggPath == "" {
		log.Fatal("cannot find .ogg file for " + f.Slug)
	}

	//use ffprobe(1) to read audio metadata (we use Ogg files for this since
	//ffprobe only reads chapter URLs reliably from these)
	cmd := exec.Command("ffprobe", "-v", "quiet", "-of", "json", "-show_chapters", "-show_format", "-show_streams", firstOggPath)
	cmd.Stderr = os.Stderr
	output := MustReturn(cmd.Output())

	type ffprobeStream struct {
		CodecType       string            `json:"codec_type"`
		DurationSecsStr string            `json:"duration"`
		Tags            map[string]string `json:"tags"`
	}
	var data struct {
		Streams  []ffprobeStream `json:"streams"`
		Chapters []struct {
			StartSecsStr string `json:"start_time"`
			EndSecsStr   string `json:"end_time"`
			Tags         struct {
				Title string `json:"title"`
			} `json:"tags"`
		} `json:"chapters"`
	}
	MustSucceed(json.Unmarshal(output, &data))

	var audioStream ffprobeStream
	for _, stream := range data.Streams {
		if stream.CodecType == "audio" {
			audioStream = stream
			break
		}
	}
	result.LengthSeconds = uint(mustParseFloat(audioStream.DurationSecsStr))

	if len(data.Chapters) > 0 {
		result.Chapters = make([]chapter, len(data.Chapters))
		for idx, ch := range data.Chapters {
			result.Chapters[idx] = chapter{
				StartSeconds: mustParseFloat(ch.StartSecsStr),
				EndSeconds:   mustParseFloat(ch.EndSecsStr),
				Title:        ch.Tags.Title,
				URL:          audioStream.Tags[fmt.Sprintf("CHAPTER%03dURL", idx+1)],
			}
		}
	}

	return result
}

func mustParseFloat(in string) float64 {
	return MustReturn(strconv.ParseFloat(in, 64))
}

////////////////////////////////////////////////////////////////////////////////
// template helper functions

func buildTemplateRenderer(path string) func(path string, data any) {
	buf := MustReturn(os.ReadFile(path))
	name := strings.TrimPrefix(path, ".tpl")
	tmpl := MustReturn(template.New(name).Funcs(templateFuncs).Parse(string(buf)))

	// for separation of concerns, this only returns a closure that renders the
	// template with the given data and writes it to the given path
	return func(path string, data any) {
		var buf bytes.Buffer
		MustSucceed(tmpl.Execute(&buf, data))
		MustSucceed(os.WriteFile(path, buf.Bytes(), 0666))
	}
}

var templateFuncs = template.FuncMap{
	"floatToUint":            floatToUint,
	"reverseFiles":           reverseFiles,
	"readableFileSize":       readableFileSize,
	"readableLengthSeconds":  readableLengthSeconds,
	"startSecondsToHHMMSS":   startSecondsToHHMMSS,
	"unixTimeToRFC1123":      unixTimeToRFC1123,
	"unixTimeToReadableDate": unixTimeToReadableDate,
}

func floatToUint(x float64) uint {
	return uint(x)
}

func reverseFiles(in []*file) (out []*file) {
	out = make([]*file, len(in))
	for idx, f := range in {
		out[len(in)-idx-1] = f
	}
	return
}

func readableFileSize(sizeBytes uint64) string {
	out := fmt.Sprintf("%.1f MiB", float64(sizeBytes)/(1<<20))
	return strings.ReplaceAll(out, ".", ",")
}

func readableLengthSeconds(lengthSeconds uint) string {
	return fmt.Sprintf("%02d:%02d", lengthSeconds/60, lengthSeconds%60)
}

func startSecondsToHHMMSS(startSecondsFloat float64) string {
	startSeconds := uint(startSecondsFloat)
	startMinutes := startSeconds / 60
	startHours := startMinutes / 60
	return fmt.Sprintf("%02d:%02d:%02d", startHours, startMinutes%60, startSeconds%60)
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
	return strings.ReplaceAll(str, "XXX", germanMonths[t.Month()])
}

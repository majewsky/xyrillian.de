{{- $file := index .Files .CurrentFileIndex -}}
{{- $show := index .Shows $file.ShowID -}}
<!DOCTYPE html>
<html lang="de">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{$file.Title}} - Xyrillian Noises</title>
  <link rel="stylesheet" href="/res/xyrillian.css" />
</head>
<body class="noises episode-{{$file.ShowID}} has-nav-top-card">
  <nav class="top-card">
    <div>
      <a href="/noises/">
        <img class="sitelogo" src="/res/logo-noises.svg" alt="Xyrillian Noises Hörfunkproduktion">
      </a>
    </div>
  </nav>
  <header>
    <img class="coverart" src="/res/{{$show.Covers.ForHTML}}" alt="Cover-Art für: {{$show.Title}}">
    <h2>
      {{- if $file.Subtitle}}
        {{- $file.Subtitle}}
      {{- else}}
        {{- $show.Title}} #{{$file.Episode}} vom {{$file.PublicationTimeUnix | unixTimeToReadableDate }}
      {{- end -}}
    </h2>
    <h1>{{$file.Title}}</h1>
    {{- if $file.HTML.Description }}
      <p>{{$file.HTML.Description}}</p>
    {{- end }}
    {{- if $file.HTML.Notes }}
      <p>{{$file.HTML.Notes}}</p>
    {{- end }}
    {{- if $file.LengthSeconds }}
      <p><strong>Länge:</strong> {{$file.LengthSeconds | readableLengthSeconds}} Minuten</p>
    {{- end }}
  </header>
  {{- if $file.Downloads }}
    <audio controls preload="none">
      {{- range $file.Downloads }}
        <source src="/dl/{{.FileName}}" type="{{.MIMEType}}">
      {{- end }}
      <div class="audio-not-supported">
        <span>Webplayer nicht vom Browser unterstützt</span>
      </div>
    </audio>
    <section>
      <h3>Download</h3>
      <ul>
        {{- range $file.Downloads }}
          <li><a href="/dl/{{.FileName}}">{{.Format}}</a> ({{.SizeBytes | readableFileSize }})</li>
        {{- end }}
      </ul>
    </section>
  {{- end }}
  {{- if $file.HTML.ShowNotes }}
    <section>
      <h3>Shownotes</h3>
      <p>{{- $file.HTML.ShowNotes }}</p>
    </section>
  {{- end }}
  {{- if $file.HTML.Sources }}
    <section>
      <h3>Quellen</h3>
      <ul>
        {{- range $file.HTML.Sources }}
          <li>{{.}}</li>
        {{- end }}
      </ul>
    </section>
  {{- end }}
  {{- if $file.Music }}
    <section>
      <h3>Musiktitel<span>&nbsp;in Abspielreihenfolge</span></h3>
      <ul>
        {{- range $file.Music }}
          <li><a href="{{.URL}}">{{.Artist}}: &quot;{{.Title}}&quot;</a>{{if .Variant}} ({{.Variant}}){{end}}</li>
        {{- end }}
      </ul>
    </section>
  {{- end }}

  <footer class="legal">
    <a href="/legal/de/">Impressum/Datenschutz</a>
  </footer>
</body>
</html>

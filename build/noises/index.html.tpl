<html lang="de">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Xyrillian Noises Hörfunkproduktion</title>
  <link rel="stylesheet" type="text/css" href="/res/xyrillian.css" />
</head>
<body class="noises">
  <header>
    <img src="/res/logo-noises.svg" alt="Xyrillian Noises Hörfunkproduktion">
    <p>Beschallungen von und mit Stefan Majewsky. Alle Inhalte frei verfügbar unter <a href="https://creativecommons.org/licenses/by-sa/3.0/de/">CC BY-SA</a>.</p>
    <p>Keine Folge verpassen! Jetzt den <a href="/noises/rss.xml">RSS-Feed für &quot;Gar nicht so einfach&quot;</a> abonnieren. Gibt's auch bei <a href="https://bitlove.org/xyrillian-noises/gar-nicht-so-einfach">Bitlove</a>, <a href="https://gpodder.net/podcast/gar-nicht-so-einfach">GPodder</a>, <a href="https://itunes.apple.com/de/podcast/gar-nicht-so-einfach/id1435500811?l=de">iTunes</a> und <a href="https://www.stitcher.com/s?fid=230739&refid=stpr">Stitcher</a>.</p>
  </header>

  {{- range (.Files | reverseFiles) }}
  {{- $show := index $.Shows .ShowID }}
    <article class="episode-{{.ShowID}}">
      <img class="coverart" src="/res/{{$show.Covers.ForHTML}}" alt="Cover-Art für: {{$show.Title}}">
      <h2>{{if .Subtitle}}{{.Subtitle}}{{else}}{{$show.Title}} #{{.Episode}}{{end}}</h2>
      <h1>{{.Title}}</h1>
      {{- if .HTML.Description }}
        <p>{{.HTML.Description}}</p>
      {{- end }}
      {{- if .HTML.Notes }}
        <p>{{.HTML.Notes}}</p>
      {{- end }}
      {{- if .LengthSeconds }}
        <p><strong>Länge:</strong> {{.LengthSeconds | readableLengthSeconds}} Minuten</p>
      {{- end }}
      {{- if not $show.IsExternal }}
        <nav>
          {{- if .Downloads }}
            <button class="selected" data-tab="download">Download</button>
          {{- end }}
          {{- if .HTML.Sources }}
            <button data-tab="sources">Quellen</button>
          {{- end }}
          {{- if .Music }}
            <button data-tab="tracks">Musiktitel</button>
          {{- end }}
        </nav>
        {{- if .Downloads }}
          <ul class="nav-show" data-tab="download">
            {{- range .Downloads }}
              <li><a href="/dl/{{.FileName}}">{{.Format}}</a> ({{.SizeBytes | readableFileSize }})</li>
            {{- end }}
          </ul>
        {{- end }}
        {{- if .HTML.Sources }}
          <ul class="nav-hide" data-tab="sources">
            {{- range .HTML.Sources }}
              <li>{{.}}</li>
            {{- end }}
          </ul>
        {{- end }}
        {{- if .Music }}
          <p class="nav-hide" data-tab="tracks">In Abspielreihenfolge:</p>
          <ul class="nav-hide" data-tab="tracks">
            {{- range .Music }}
              <li><a href="{{.URL}}">{{.Artist}}: &quot;{{.Title}}&quot;</a>{{if .Variant}} ({{.Variant}}){{end}}</li>
            {{- end }}
          </ul>
        {{- end }}
      {{- end }}
    </article>
  {{- end }}

  <footer>
    <a href="/legal/">Impressum/Datenschutz</a>
  </footer>
  <script type="text/javascript" src="/res/xyrillian.js"></script>
</body>
</html>

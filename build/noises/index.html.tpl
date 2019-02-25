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
    <img class="sitelogo" src="/res/logo-noises.svg" alt="Xyrillian Noises Hörfunkproduktion">
    <p>Beschallungen von und mit Stefan Majewsky. Alle Inhalte frei verfügbar unter <a href="https://creativecommons.org/licenses/by-sa/3.0/de/">CC BY-SA</a>.</p>
    <p>Keine Folge verpassen! Jetzt die RSS-Feeds für <a href="/noises/rss.xml">&quot;Gar nicht so einfach&quot;</a> und <a href="/noises/st/rss.xml">&quot;Space Tourist&quot;</a> abonnieren. Gibt's auch bei <a href="https://bitlove.org/xyrillian-noises/gar-nicht-so-einfach">Bitlove</a>, <a href="https://gpodder.net/podcast/gar-nicht-so-einfach">GPodder</a>, <a href="https://itunes.apple.com/de/podcast/gar-nicht-so-einfach/id1435500811?l=de">iTunes</a> und <a href="https://www.stitcher.com/s?fid=230739&refid=stpr">Stitcher</a>.</p>
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
      {{- if $show.IsExternal }}
        {{- if .HTML.Notes }}
          <p>{{.HTML.Notes}}</p>
        {{- end }}
      {{- else }}
        <div class="player">
          <a href="/noises/{{.ShowID}}/{{.Slug}}">Details</a>
          <audio controls preload="none">
            {{- range .Downloads }}
              <source src="/dl/{{.FileName}}" type="{{.MIMEType}}">
            {{- end }}
            <div class="audio-not-supported">
              <span>Webplayer nicht vom Browser unterstützt</span>
            </div>
          </audio>
        </div>
      {{- end }}
    </article>
  {{- end }}

  <footer class="legal">
    <a href="/legal/de/">Impressum/Datenschutz</a>
  </footer>
</body>
</html>

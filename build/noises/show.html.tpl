{{- $show := index .Shows .CurrentShowID -}}
<html lang="de">
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{$show.Title}} - Xyrillian Noises</title>
  <link rel="stylesheet" type="text/css" href="/res/xyrillian.css" />
</head>
<body class="noises episode-{{.CurrentShowID}} has-nav-top-card">
  <nav class="top-card">
    <div>
      <a href="/noises/">
        <img class="sitelogo" src="/res/logo-noises.svg" alt="Xyrillian Noises Hörfunkproduktion">
      </a>
    </div>
  </nav>
  <header>
    <img class="coverart" src="/res/{{$show.Covers.ForHTML}}" alt="Cover-Art für: {{$show.Title}}">
    <h1>{{$show.Title}}</h1>
    <h2>{{$show.Subtitle}}</h2>
    <p>{{$show.DescriptionHTML}}</p>
    <p class="subscribe-options">
      <a class="rss-feed">RSS-Feed abonnieren</a>
      oder über
      {{ range $idx, $sub := $show.Subscriptions -}}
        {{ if gt $idx 0 }}, {{ end -}}
        <a href="{{$sub.URL}}">{{$sub.Name}}</a>
      {{- end }}
    </p>
  </header>

  {{- range (.Files | reverseFiles) }}
  {{- if eq .ShowID $.CurrentShowID }}
    <article class="episode-{{.ShowID}}">
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
  {{- end }}

  <footer class="legal">
    <a href="/legal/de/">Impressum/Datenschutz</a>
  </footer>
</body>
</html>

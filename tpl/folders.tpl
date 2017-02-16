{{define "folders"}}
<ul>
  {{range . }}
    <li id="{{.Name}}">{{.Name}}</li>
  {{end}}
</ul>
{{end}}

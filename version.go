package main

import (
	"bytes"
	"runtime"
	"text/template"
)

var (
	Version   = "v0.0.2"
	BuildTime = "2020/07/02"
)

type VersionOptions struct {
	GitCommit string
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = ` Version:      {{.Version}}
 Go version:   {{.GoVersion}}
 Built:        {{.BuildTime}}
 OS/Arch:      {{.Os}}/{{.Arch}}
 `

func getVersion() string {
	var doc bytes.Buffer
	vo := VersionOptions{
		Version:   Version,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
	tmpl, _ := template.New("version").Parse(versionTemplate)
	tmpl.Execute(&doc, vo)
	return doc.String()
}

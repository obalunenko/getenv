package gofish

type templateData struct {
	Name            string
	Desc            string
	Homepage        string
	Version         string
	License         string
	ReleasePackages []releasePackage
}

type releasePackage struct {
	DownloadURL string
	SHA256      string
	OS          string
	Arch        string
	Binaries    []binary
}

type binary struct {
	Name, Target string
}

const foodTemplate = `local name = "{{ .Name }}"
local version = "{{ .Version }}"

food = {
    name = name,
    description = "{{ .Desc }}",
    license = "{{ .License }}",
    homepage = "{{ .Homepage }}",
    version = version,
    packages = {
    {{- range $element := .ReleasePackages}}
    {{- if ne $element.OS ""}}
        {
            os = "{{ $element.OS }}",
            arch = "{{ $element.Arch }}",
            url = "{{ $element.DownloadURL }}",
            sha256 = "{{ $element.SHA256 }}",
            resources = {
                {{- range $binary := $element.Binaries }}
                {
                    path = "{{ $binary.Name }}",
                    installpath = {{if ne $element.OS "windows"}}"bin/{{ $binary.Target }}"{{else}}"bin\\{{ $binary.Target }}"{{end}},
                    {{- if ne $element.OS "windows"}}
                    executable = true
                    {{- end }}
                },
                {{- end }}
            }
        },
    {{- end }}
    {{- end}}
    }
}`

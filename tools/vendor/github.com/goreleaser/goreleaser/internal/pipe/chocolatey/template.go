package chocolatey

type templateData struct {
	Packages []releasePackage
}

type releasePackage struct {
	DownloadURL string
	Checksum    string
	Arch        string
}

const scriptTemplate = `# This file was generated by GoReleaser. DO NOT EDIT.
$ErrorActionPreference = 'Stop';

$version = $env:chocolateyPackageVersion
$packageName = $env:chocolateyPackageName
$toolsDir = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"

$packageArgs = @{
    packageName    = $packageName
    unzipLocation  = $toolsDir
    fileType       = 'exe'
    {{- range $release := .Packages }}
    {{- if eq $release.Arch "amd64" }}
    url64bit       = '{{ $release.DownloadURL }}'
    checksum64     = '{{ $release.Checksum }}'
    checksumType64 = 'sha256'
    {{- else }}
    url            = '{{ $release.DownloadURL }}'
    checksum       = '{{ $release.Checksum }}'
    checksumType   = 'sha256'
    {{- end }}
    {{- end }}
}

Install-ChocolateyZipPackage @packageArgs
`
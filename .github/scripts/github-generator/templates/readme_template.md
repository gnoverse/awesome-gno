<div align="center">
	<img src="./banner.png" />
</div>

# Awesome Gno [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

A curated list of awesome frameworks, libraries, software and resources related to the <a href='https://gno.land'>Gnoland</a> blockchain.

Gnoland is a robust blockchain that provides concurrency and scalability with smart contracts programmed in Gno, a Go interpreter.

> Do you gno?

## Contents

{{ range . }}
{{ .Order }}. [{{ .Title }}](#{{ .Slug }})
{{ end }}

{{ range . }}

## {{ .Title }}

{{ if .Note }}{{ .Note }}{{ end }}

{{ range .Entries }}

- [{{ .Title }}]({{ .Link }}){{ if .IsStaffPick }} - ![staffpick](./staffpick.png){{ end }} {{ if .Description }}- {{ .Description }}{{ end }}
{{ end }}
{{ end }}

## Contributing

Your contributions are always welcome! Please take a look at the [contribution guidelines](https://github.com/gnolang/awesome-gno/blob/master/CONTRIBUTING.md) first.

We will keep some pull requests open if we're not sure whether those libraries are awesome. You could [vote for them](https://github.com/gnolang/awesome-gno/pulls) by adding :+1: to them.

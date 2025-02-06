<div align="center">
    <a href="https://github.com/hurfy/hashGo"><img src="" alt="hashGo" /></a>
</div>

<div align="center">
    <img src="https://img.shields.io/github/issues/hurfy/hashGo?style=for-the-badge" alt="open issues" />
    <img src="https://img.shields.io/badge/version-1.1.0-blue?style=for-the-badge" alt="version" /></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/hurfy/hashGo?style=for-the-badge" alt="license" /></a>
</div>

<br />

<div align="center">
  Simple hashing for your files
</div>

<div align="center">
  <sub>
    Built with love 
    &bull; Brought to you by <a href="https://github.com/hurfy">@hurfy</a>
    and other <a href="https://github.com/hurfy/hashGo/graphs/contributors">contributors</a>
  </sub>
</div>

## Introduction
This lightweight and fast project allows you to generate hash sums for files in a specified directory. Hashing is useful for quickly identifying changes in files, comparing versions, or ensuring that files haven't been altered. Its efficiency makes it ideal for processing large numbers of files without a significant delay.

## Features
- Supports the following hashing algorithms: `md5`, `sha1`, `sha256`, `sha512`.
- Works with files in any specified directory.
- Ability to save hash results in a JSON file for easy reference.
- Option to exclude directories from the hashing process.
- Simple command-line interface for easy usage.


## Quick Start
Use make to build:
```bash
make build
```
Or do it yourself:
```bash
go build -ldflags "-s -w" ./cmd/hashGo
```
**\*You can also see other configurations of Make**

Startup example:
```bash
hashGo.exe -p C:/Games/SuperDuperGame -o data -f sha256 -e Img;Video -s
```

### Flags
| Flag | Name         | Default       | Description                              |
|------|--------------|---------------|------------------------------------------|
| p    | Input Path   | Current       | Root directory                           |
| o    | Output File  | None          | Output file name(console output if None) |
| f    | Format       | md5           | Hash format(md5, sha1, sha256, sha512)   |
| s    | Subdirs      | false         | Include subdirectories                   |
| e    | Exclude dirs | None          | Exclude directories(semicolon-separated) |

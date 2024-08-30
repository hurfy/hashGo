## About
This is a small project that will help you hash files in a specified directory.<br>
Just run it with the type, directory, and other parameters.<br>
It supports the following types: md5, sha1, sha256, sha512.

## Build
Use Make:
```bash
make
```

Or build it yourself:
```bash
go build -ldflags "-s -w" ./cmd/hashGo
```

## Run & Flags

```bash
./hashGo.exe -p C:/Games/SuperDuperGame -f sha256 -sd -o -s -op ./output/output.json
```

| Flag | Name         | Default       | Description                               |
|------|--------------|---------------|-------------------------------------------|
| p    | Input path   | ./            | Path to directory                         |
| f    | Format       | md5           | Output format (md5, sha1, sha256, sha512) |
| o    | Show Output  | false         | Show console output                       |
| sd   | Subdirs      | false         | Include subdirectories                    |
| s    | Save to file | false         | Save output to file                       |
| op   | Output path  | ./output.json | Path to output file                       |

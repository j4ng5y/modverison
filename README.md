# modversion

A simple little app that helps me in SemVer stuff, including, but not limited to, go modules.

## Usage

I feel like the usage is fairly self explanitory, but if anyone wants me to make more documentation, let me know in an [Issue](https://github.com/j4ng5y/modversion/issues).

```bash
bin/modversion -h                       
A simple method of automated Go module versioning

Usage:
  version [flags]
  version [command]

Available Commands:
  help        Help about any command
  major++     Increment Major Version
  major--     Decrement Major Version
  minor++     Increment Minor Version
  minor--     Decrement Minor Version
  patch++     Increment Patch Version
  patch--     Decrement Patch Version

Flags:
  -h, --help                  help for version
  -v, --version               version for version
  -f, --version-file string   The file in which to store the versioning information (default "./VERSION")

Use "version [command] --help" for more information about a command.
```
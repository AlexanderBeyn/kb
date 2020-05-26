# kb

`kb` is a command line client for [Kanboard](https://kanboard.org/).

## Installation

`kb` executable has no dependencies. Download the 
[latest release](releases/latest), uncompress, and place somewhere in your
executable path.

Alternatively, `kb` can be installed with `go get` and `go install`:
```
$ go get -u github.com/AlexanderBeyn/kb
$ go install github.com/AlexanderBeyn/kb
```

## Configuration

`kb` reads its configuration from `.kb.yaml` in your home directory. 
Additionally, `kb` caches some data from the server in this file for 
quicker responses.

### Server access
The quickest way to create a new configuration file is with 
`kb config prompt`. This will prompt you for a Kanboard API URL, your
username, and your password or API key.

### Defaults
The default project and column can be set with `kb config defaults`.

## Usage

```
Usage:
  kb [command]

Available Commands:
  add         Add a new task
  completion  Generate shell completion code for the specified shell
  config      Manage kb configuration
  help        Help about any command
  move        Move a task to a new column
  show        Show tasks
```

`kb help` provides more detailed usage information. This information is
also available in the [docs/](docs/kb.md) directory.

### Common arguments

Most commands accept the project, columns, and search terms on the command
line by prepending them with a sigil:

| Sigil | Meaning |
| ----- | ------- |
| `%%` | Project | 
| `%` | Column | 
| `^%` | Source Column | 
| `/` | Search query |

For columns and projects, the text following the sigil a prefix. If there
are multiple items matching this prefix, `kb` will prompt for a more
specific item.

For example, given a Kanboard instance with columns `Today`, `This week`,
and `Backlog`:

| Argument | Meaning |
| -------- | ------- |
| %today | Uses `Today` column |
| %b | Uses `Backlog` column |
| %t | Prompts the user to select between `Today` and `This week` |

Some commands (such as [kb add](docs/kb_add.md)) also take `+` as a final
argument, to signify extended input. There needs to be a space before the
`+` on the command line.


## License
[MIT](https://choosealicense.com/licenses/mit/)
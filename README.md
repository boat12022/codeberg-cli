# codeberg-cli

A command-line interface for interacting with [Codeberg](https://codeberg.org), written in Go.

## Installation

Requires Go 1.26 or newer.

```sh
go install codeberg.org/13thab/codeberg-cli@latest
```

Or build from source:

```sh
git clone https://codeberg.org/13thab/codeberg-cli
cd codeberg-cli
go build -o cb
```

## Usage

```
cb <command> [arguments]
```

### Commands

#### `user <username>`

Get information about a Codeberg user.

```sh
cb user 13thab
```

Output:

```
Username: 13thab
Starred Repos: 4
Followers: 2
Following: 3
Created: 2024-08-12T18:21:00Z
```

#### `repos <username>`

List the public repositories of a user.

```sh
cb repos 13thab
```

Output:

```
codeberg-cli (Go) ⭐ 1
dotfiles (Shell) ⭐ 0
```

#### `repo <username> <repo>`

Get detailed information about a single repository.

```sh
cb repo 13thab codeberg-cli
```

Output:

```
Repository: codeberg-cli
Description: A command-line interface for Codeberg
Owner: 13thab
Language: Go
Stars: 1
Forks: 0
Watchers: 1
Created: 2024-08-12T18:21:00Z
Updated: 2025-06-21T00:00:00Z
URL: https://codeberg.org/13thab/codeberg-cli
```

#### `issues <username> <repo>`

List issues for a repository.

```sh
cb issues 13thab codeberg-cli
```

Output:

```
#3 Add clone command
#1 Document the repo command
```

Flags:

- `-s`, `--state` — filter by state (default `open`)
- `-l`, `--limit` — number of issues to display (default `10`)

#### `prs <username> <repo>`

List pull requests for a repository.

```sh
cb prs 13thab codeberg-cli
```

Output:

```
#5 [open] Add prs command
#2 [open] Add issues command
```

Flags:

- `-s`, `--state` — filter by state (default `open`)
- `-l`, `--limit` — number of results (default `10`)

#### `clone <username> <repo>`

Clone a Codeberg repository.

```sh
cb clone 13thab codeberg-cli
```

#### `stats <username>`

Show statistics for a Codeberg user.

```sh
cb stats 13thab
```

## Project structure

```
.
├── cmd/                 # Cobra command definitions (root, user, repos, repo, issues, prs, clone, stats)
├── internal/
│   ├── codeberg/        # HTTP client for the Codeberg API
│   └── models/          # API response types
└── main.go              # Entry point
```

The client talks to the public Codeberg API at `https://codeberg.org/api/v1`.

## Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) — command framework

## License

No Licence 😝

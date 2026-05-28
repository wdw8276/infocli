# infocli

A simple CLI tool to store and query text records, backed by SQLite.

## Install

```bash
make install
```

Auto-detects the current OS and architecture (Linux, macOS amd64/arm64),
builds the binary, and installs it to `/usr/local/bin/infocli`.

Default database file: `~/.<username>.db` (e.g. `~/.fish.db`)

## Usage

```
infocli [flags] <command>
```

Global flags can be placed before or after the subcommand:

```bash
infocli -i 1 d      # flag before subcommand
infocli d -i 1      # flag after subcommand
```

### Global Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--file` | `-f` | `~/.<user>.db` | Database file path |
| `--id` | `-i` | `0` | Record ID (required for update/delete) |
| `--detail` | `-d` | false | Show full fields (ID, Created, Updated) |
| `--debug` | `-D` | false | Enable debug mode |

### Commands

| Command | Description |
|---------|-------------|
| `a <name> [data]` | Add a record; data can be piped from stdin |
| `q <name>` | Query by name (simple: shows ID + Name only) |
| `q name <name>` | Query by name (full output) |
| `q data <data>` | Query by data field |
| `q id <id>` | Query by ID |
| `u name -i <id> <name>` | Update name of a record |
| `u data -i <id> [data]` | Update data of a record; data can be piped |
| `d -i <id>` | Delete a record |
| `c` | Count records and show last update time |
| `init` | Initialize / migrate the database |
| `v` | Print version |

### Examples

```bash
# Add a record
infocli a mykey "some value"

# Add a record with piped data
cat file.txt | infocli a mykey

# Query by name (fuzzy match)
infocli q my

# Query with full detail
infocli q name mykey -d

# Update data by ID
infocli u data -i 1 "new value"

# Update data from stdin
cat new.txt | infocli u data -i 1

# Delete by ID (-i can be before or after subcommand)
infocli d -i 1
infocli -i 1 d

# Count all records
infocli c

# Use a custom database file
infocli -f /path/to/custom.db init
infocli -f /path/to/custom.db a mykey "value"
infocli -f /path/to/custom.db q mykey
```

## Development

```bash
make dev      # quick local build → build/infocli-dev
make test     # build and run smoke tests
make all      # cross-compile for linux / darwin-amd64 / darwin-arm64
make release  # build all platforms, tag and push version
make clean    # remove build/
```

## Shell Alias (Recommended)

Add a short alias to your shell config for faster usage:

**bash** (`~/.bashrc`):
```bash
alias ic='infocli'
```

**zsh** (`~/.zshrc`):
```zsh
alias ic='infocli'
```

**fish** (`~/.config/fish/config.fish`):
```fish
alias ic='infocli'
```

Then reload your shell (`source ~/.bashrc` / `source ~/.zshrc`) and use:

```bash
ic c                        # count records
ic a mykey "some value"     # add
ic q mykey                  # query
ic d -i 1                   # delete
```

If you use a non-default database file, bake it into the alias:

```bash
alias ic='infocli -f ~/mydata.db'
```

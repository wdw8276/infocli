# infocli

A simple CLI tool to store and query text records, backed by SQLite.

## Install

```bash
go build -o infocli .
```

Default database file: `~/.<username>.db` (e.g. `~/.fish.db`)

## Usage

```
infocli [flags] <command>
```

### Global Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--file` | `-f` | `~/.<user>.db` | Database file path |
| `--id` | `-i` | `0` | Record ID |
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

# Delete by ID
infocli d -i 1

# Count all records
infocli c
```

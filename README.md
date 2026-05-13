# Pokedex CLI

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/API-PokeAPI-blue?style=for-the-badge" alt="API">
</p>

A command-line Pokedex built with Go, powered by the [PokeAPI](https://pokeapi.co/).

## Features

- **Explore Locations** - Discover Pokemon in different location areas
- **Catch Pokemon** - Catch and collect Pokemon
- **Inspect Pokemon** - View stats of caught Pokemon
- **Pagination** - Navigate through location areas with `map` and `mapb`

---

## Installation

```bash
# Clone the repository
git clone https://github.com/VoluteTech/pokedexcli.git
cd pokedexcli

# Build the project
go build -o pokedexcli .

# Run the CLI
./pokedexcli
```

---

## Usage

Once running, you'll see the prompt:

```
Pokedex >
```

### Available Commands

| Command | Description |
|---------|-------------|
| `map` | List the next 20 location areas |
| `mapb` | List the previous 20 location areas |
| `explore <location>` | List all Pokemon found in a location |
| `catch <pokemon>` | Attempt to catch a Pokemon |
| `inspect <pokemon>` | View stats of a caught Pokemon |
| `help` | Show available commands |
| `exit` | Exit the Pokedex |

---

## Examples

### Browse Location Areas

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > mapb
(goes back to previous page)
```

### Explore a Location

```
Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - absol
 - aerodactyl
 - aggron
 - aipom
 ...
```

### Catch a Pokemon

```
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu has been captured!
```

### Inspect Caught Pokemon

```
Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
- hp: 35
- attack: 55
- defense: 40
- special-attack: 50
- special-defense: 50
- speed: 90
```

---

## Project Structure

```
pokedexcli/
├── cmd.go                  # CLI command implementations
├── main.go                 # Entry point
├── repl.go                 # REPL loop and command registration
├── internal/
│   ├── api/                # PokeAPI client and types
│   │   ├── client.go       # HTTP client with caching
│   │   ├── locations_list.go
│   │   ├── type_locations.go
│   │   └── type_pokemon.go
│   └── cache/              # In-memory cache
│       └── cache.go
└── README.md
```

---

## How It Works

1. **REPL Loop** - The CLI runs an interactive read-eval-print loop in `repl.go`
2. **API Client** - `internal/api/client.go` handles HTTP requests to PokeAPI
3. **Caching** - Responses are cached in-memory to reduce API calls
4. **State** - The `config` struct in `repl.go` tracks pagination URLs and caught Pokemon

---

## License

MIT License - See [LICENSE](LICENSE) for details.

---

<p align="center">Made with  by VoluteTech</p>

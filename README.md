# Pokedex CLI

A command-line Pokédex built in **Go** that interacts with the **PokéAPI** to explore the Pokémon world directly from your terminal. Navigate through map locations, discover Pokémon in different areas, catch them, inspect their stats, and build your own Pokédex.

---

# What this project does

This project is an interactive CLI application that simulates a simplified Pokédex experience.

Using terminal commands, you can:

* Explore Pokémon world locations
* Navigate forward and backward through map areas
* Discover which Pokémon appear in each location
* Attempt to catch Pokémon
* Store successfully caught Pokémon in your own Pokédex
* Inspect detailed information about captured Pokémon

The application communicates with the PokéAPI in real time and demonstrates how to build interactive command-line software using Go.

---

# Topics I mastered doing it

While building this project, I practiced several important Go concepts, including:

* Building interactive CLI applications
* Parsing user input
* Command routing using maps and callbacks
* HTTP client requests
* Consuming REST APIs
* JSON decoding with `encoding/json`
* Struct modeling
* Error handling
* State management
* Caching API responses
* Working with pointers
* Variadic function parameters
* Organizing code into multiple files
* Randomized game mechanics (catch probability)
* Writing idiomatic Go code

---

# Available Commands

| Command              | Description                                        |
| -------------------- | -------------------------------------------------- |
| `help`               | Display all available commands                     |
| `exit`               | Exit the application                               |
| `map`                | Display the next Pokémon location areas            |
| `mapb`               | Display the previous location areas                |
| `explore <location>` | Show every Pokémon that can be found in a location |
| `catch <pokemon>`    | Attempt to catch a Pokémon                         |
| `inspect <pokemon>`  | Inspect a caught Pokémon and display its stats     |
| `pokedex`            | List every Pokémon you've successfully caught      |

Example:

```bash
map

explore pastoria-city-area

catch pikachu

inspect pikachu

pokedex
```

---

# Why someone should care

Although this project looks simple, it introduces many real-world backend concepts that Go developers use every day.

It demonstrates how to:

* Consume third-party APIs
* Build maintainable command-line tools
* Organize Go projects cleanly
* Work with JSON responses
* Handle application state
* Implement reusable command patterns
* Build software without external frameworks

If you're learning Go, this project is an excellent example of how multiple language features come together in a practical application.

---

# Technologies

* Go
* PokéAPI
* HTTP
* JSON
* Standard Library

---

# Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/pokedex-cli.git
```

Enter the project directory:

```bash
cd pokedex-cli
```

Run the application:

```bash
go run .
```

Or build it:

```bash
go build
./pokedex-cli
```

---

# Credits

This project was built as part of the **Boot.dev** Go Backend curriculum and uses data provided by the **PokéAPI**.

https://pokeapi.co/

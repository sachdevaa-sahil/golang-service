# Go basics

Small runnable examples of Go data types, arrays, slices, and operators.

## Project structure

```text
main.go             # Entry point; choose which examples to run here
cmd/users/          # User examples
cmd/basics/
  basic_types.go    # Variables and basic data types
  arrays_slices.go  # Arrays and slices
  operators.go      # Arithmetic, comparisons, and other operators
.vscode/launch.json # VS Code debugger configuration
docker-compose.yml # Local MongoDB service
go.mod              # Go module and toolchain version
```

All files in `cmd/basics` belong to the same `basics` package and compile together.

## Run

From the project root:

```sh
go run .
```

The program prints "Main Program" by default. In `main.go`, uncomment an
example call and its matching package import to run it.

## Debug

With the VS Code Go extension and Delve installed, select **Debug Go Project**
in Run and Debug and press **F5**. Set breakpoints in any example file.

## MongoDB

With Docker running, start the optional local database from the project root:

```sh
docker compose up -d mongodb
```

Stop it with `docker compose down`. Data persists in the named Docker volume.
The Go examples do not currently connect to MongoDB.

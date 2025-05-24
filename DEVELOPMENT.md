# Developing phrasegen

Development for phrasegen should occur in the included [devconatiner](./.devcontainer).

```bash
code .
# When prompted, re-open in the development container
# by selecting the bottom right popup
```

## Development

Common entrypoints are driven by [`just`](https://github.com/casey/just).

```bash
# List all possible entrypoints
just
```

### Build

```bash
just build
```

By default, the above will build for a common set of platforms/architectures
(which can be seen via `just list-default-platforms`).

The resulting binary will be stored in the `./binaries` directory.

If you want to build for a single platform, use `just build-for <platform>`. To
get a list of possible platform values, run `just list-all-platforms`.

### Testing

```bash
just test
```

### Running

```bash
just run
```

Or, manually invoke the binary from `./binaries` after building.

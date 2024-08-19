# Daggerverse Vault Module

[Dagger](https://dagger.io/) module for [daggerverse](https://daggerverse.dev/) providing Vault functionality.

The Dagger module is located in the [vault-kv](./vault-kv/) directory.

## Usage

Basic usage guide.

The [vault-kv](./vault-kv/) directory contains a [daggerverse](https://daggerverse.dev/) [Dagger](https://dagger.io/) module.

Check the official Dagger Module documentation: https://docs.dagger.io/

The [Dagger CLI](https://docs.dagger.io/cli) is needed.

### Functions

List all functions of the module. This command is provided by the [Dagger CLI](https://docs.dagger.io/reference/cli/). 

```bash
dagger functions -m ./vault-kv/
```

The vault-kv module is referenced locally.

## Development

Basic development guide.

### Set up Dagger module

```bash
# enter into the module's directory
cd vault-kv/
# initialize the module
dagger develop --sdk go
```

### Testing

This module contains a testing module that aims to test Dagger vault-kv module.

```bash
# enter into the test module's directory
cd tests/
# initialize the module
dagger develop --sdk go
# execute the tests
dagger call test
```

## To Do

- [ ] Add cache mounts
- [ ] Add environment variables
- [ ] Add more examples

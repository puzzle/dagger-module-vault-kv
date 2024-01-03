# daggerverse Vault Module

[Dagger](https://dagger.io/) module for [daggerverse](https://daggerverse.dev/) providing Vault functionality.

The Dagger module is located in the [vault-kv](./vault-kv/) directory.

## usage

Basic usage guide.

The [vault-kv](./vault-kv/) directory contains a [daggerverse](https://daggerverse.dev/) [Dagger](https://dagger.io/) module.

Check the official Dagger Module documentation: https://docs.dagger.io/zenith/

The [Dagger CLI](https://docs.dagger.io/cli) is needed.

### functions

List all functions of the module. This command is provided by the [Dagger CLI](https://docs.dagger.io/cli). 

```bash
dagger functions -m ./vault-kv/
```

The vault-kv module is referenced locally.

## development

Basic development guide.

### setup Dagger module

Setup the Dagger module.

Create the directory for the module and initialize it.

```bash
mkdir vault-kv/
cd vault-kv/

# initialize Dagger module
dagger mod init --sdk go --name vault-kv
```

### setup development module

Setup the outer module to be able to develop the Dagger vault-kv module.

```bash
dagger mod init --sdk go --name modest
dagger mod use ./vault-kv
```

Generate or re-generate the Go definitions file (dagger.gen.go) for use in code completion.

```bash
dagger mod install
```

The functions of the module are available by the `dag` variable. Type `dag.` in your Go file for code completion.


Update the module:

```bash
dagger mod update
```

## To Do

- [ ] document functions
- [ ] Add cache mounts
- [ ] Add environment variables
- [ ] Add more examples
- [ ] Add tests

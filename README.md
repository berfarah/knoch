# Knoch

### `config`

Add configuration to `.knoch` file

```sh
kh config <KEY> <VALUE>
```

### `help`

Output help for everything

```sh
kh help
```

### `init`

Create a directory with the `.knoch` file

```sh
kh init [<DIR>]
```

### `add`

Clone the repository, add it to `.knoch`

```sh
kh add <REPO> [<DIR>]
```

### `remove`

Remove the repository, remove it from `.knoch`

```sh
kh remove <REPO/DIR>
```

### `bundle`

Clone missing repositories, sync existing ones

```sh
kh bundle
```

### `sync`

Fetch all repositories. Fast forward if possible. If there is unpushed work,
warn but do nothing.

```sh
kh sync
```

### `open`

`cd`s into the project

V2: If there are multiple, give a choice?

```sh
kh open <REPO>
```

### `edit`

`cd` and `$EDITOR` the project

```sh
kh edit <REPO>
```

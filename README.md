# Knoch

### `bundle`

Clone missing repositories, sync existing ones

```sh
knoch bundle
```

### `add`

Clone the repository, add it to `.knoch`

```sh
knoch add <REPO/DIR> [<DIR>]
```

### `remove`

Remove the repository, remove it from `.knoch`

```sh
knoch remove <DIR>
```

### `list`

List tracked repositories

```sh
knoch remove <DIR>
```

### `open`

Open the selected project in $EDITOR

```sh
knoch open <DIR>
```

### `show`

Show full path of selected project

```sh
knoch show <DIR>
```

### `help`

Output help for everything

```sh
knoch help
```

## If you want `knoch cd [DIR]`

Unfortunately, due to limitations of how - well, I guess how system calls work,
this program cannot change directories for you. However, you can create a bash
command that can do it for you pretty easily:

```sh
# Add this into your .*_profile or .*rc
knoch-cd() {
  cd $(knoch show $1)
}

# If you're on Zsh, this is how you add command completion:
_knoch-cd() {
  _arguments '*:projects:_values $(knoch ls --name-only)'
}
compdef _knoch-cd knoch-cd
```

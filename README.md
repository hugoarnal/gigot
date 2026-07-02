# gigot

An easy gitconfig switcher

## Explanation

Sometimes, you must use different git configurations (example: one personal and one for work).

However, it can be annoying to switch between these configurations.
The most common way to this is using the [includeIf](https://git-scm.com/docs/git-config#_includes).

## Usage

<!-- TODO: perhaps this should be moved to a docs folder -->

> All the configuration files can be found in `$XDG_CONFIG_HOME/gigot`

### Shell Configuration

Inside your shell's configuration file, you must add the `init` script at the end of your file.

For example, for zsh:

```sh
eval "$(gigot init zsh)"
```

### Adding a new gitconfig

You need to specify the path and a name you would like to give to your associated gitconfig file:
```sh
gigot add --name "Work" --path "/home/hello-world/work.gitconfig"
```

### Switching to another gitconfig

If you want to enable your "Work" config:

```sh
gigot switch "Work"
```

If you want to disable "Work", run the previous command again.

If you want to disable any enabled configuration:

```sh
gigot switch --disable
```

### Removing a gitconfig

If you want to delete the "Work" configuration:

```sh
gigot remove --name "Work"
```

### Getting the currently enabled gitconfig

```sh
gigot get-enabled
gigot get-enabled --name
gigot get-enabled --path
```

If no gitconfigs are enabled, gigot returns just an empty string.

### List all gitconfigs

```sh
gigot list
```

## Build

Build the binary using Go:

```sh
make
```

## Tips

### Extending previous configuration with some changes

If you wish to extend your previous "main" configuration with just some changes (changing the ssh key, the email etc...), here's how you can do that:

```ini
[include]
    path = "~/.gitconfig"
[user]
    name = John Doe
    email = john.doe@work.com
[core]
    sshCommand = ssh -i ~/.ssh/id_work
```

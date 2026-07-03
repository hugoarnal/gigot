# shellcheck shell=bash

GIGOT_SHELL="$(basename "$0")"
export GIGOT_SHELL

GIGOT_BINARY="gigot"

# GIGOT_DEV must be set when the script is evaluated
if [ -n "$GIGOT_DEV" ]; then
    GIGOT_BINARY="./gigot"
fi

__gigot_disable_git_config_global() {
    unset GIT_CONFIG_GLOBAL
}

__gigot_set_git_config_global() {
    GIT_CONFIG_GLOBAL="$1"

    # GIT_CONFIG_GLOBAL if empty, doesn't fallback to the "global" gitconfig
    # Therefore, we need to unset it
    if [ -z "$GIT_CONFIG_GLOBAL" ]; then
        __gigot_disable_git_config_global
    else
        export GIT_CONFIG_GLOBAL
    fi
}

__gigot_set_global_config() {
    __gigot_set_git_config_global "$($GIGOT_BINARY get-enabled --path)"
}

__gigot_set_global_config

gigot() {
    command "$GIGOT_BINARY" "$@"

    if [ "$1" = "switch" ]; then
        __gigot_set_global_config
    fi
    if [ "$1" = "temp" ]; then
        if [ "$2" = "-disable" ] || [ "$2" = "--disable" ]; then
            __gigot_disable_git_config_global
        else
            # Silence all errors from the get command
            __gigot_set_git_config_global "$($GIGOT_BINARY get --path "$2" 2>/dev/null)"
        fi
    fi
}

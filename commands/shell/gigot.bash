# shellcheck shell=bash

export GIGOT_SHELL="$(basename "$(echo $0)")"

GIGOT_BINARY="gigot"

# GIGOT_DEV must be set when the script is evaluated
if ! [ -z "$GIGOT_DEV" ]; then
    GIGOT_BINARY="./gigot"
fi

__gigot_set_global_config() {
    GIT_CONFIG_GLOBAL="$($GIGOT_BINARY get-enabled --path)"

    # GIT_CONFIG_GLOBAL if empty, doesn't fallback to the "global" gitconfig
    # Therefore, we need to unset it
    if [ -z "$GIT_CONFIG_GLOBAL" ]; then
        unset GIT_CONFIG_GLOBAL
    else
        export GIT_CONFIG_GLOBAL
    fi
}

__gigot_set_global_config

gigot() {
    command "$GIGOT_BINARY" "$@"

    if [ "$1" = "switch" ]; then
        __gigot_set_global_config
    fi
}

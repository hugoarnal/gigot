# shellcheck shell=zsh

local binary="gigot"

# GIGOT_DEV must be set when the script is evaluated
if ! [ -z "$GIGOT_DEV" ]; then
    binary="./gigot"
fi

gigot() {
    command "$binary" "$@"
}

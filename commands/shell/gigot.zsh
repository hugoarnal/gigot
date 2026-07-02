# shellcheck shell=bash

GIGOT_BINARY="gigot"

# GIGOT_DEV must be set when the script is evaluated
if ! [ -z "$GIGOT_DEV" ]; then
    GIGOT_BINARY="./gigot"
fi

gigot() {
    command "$GIGOT_BINARY" "$@"
}

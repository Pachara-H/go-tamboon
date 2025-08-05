#!/bin/bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
MOCK_DIRS=(
    $(realpath "$SCRIPT_DIR/../pkg/external/omise")
    $(realpath "$SCRIPT_DIR/../pkg/utilities")
)
for DIR in ${MOCK_DIRS[*]}; do
    echo "# MOCKING >> $DIR/"
    MOCKED_FILES=()
    SKIPPED_FILES=()
    # clear mocks dir
    if [ -d "$DIR/mocks" ]; then
        echo "# Cleaning >> $DIR/mocks"
        rm -r "$DIR/mocks"
    fi
    for MOCKING_FILE in $DIR/*; do
        MOCKING_FILE_NAME=$(basename -- "$MOCKING_FILE")
        if [[ $MOCKING_FILE_NAME =~ \.go$ ]] &&
            [[ $MOCKING_FILE_NAME != *_test.go ]] &&
            [[ $MOCKING_FILE_NAME != _* ]]; then
            (cd $DIR && mockgen -source "$MOCKING_FILE_NAME" -destination "./mocks/$MOCKING_FILE_NAME")
            MOCKED_FILES+=($MOCKING_FILE_NAME)
        else
            SKIPPED_FILES+=($MOCKING_FILE_NAME)
        fi
    done

    echo "# mocked files"
    printf '# - %s\n' "${MOCKED_FILES[@]}"
    echo "\n# skipped files"
    printf '# - %s\n' "${SKIPPED_FILES[@]}"
    echo "\n"
done

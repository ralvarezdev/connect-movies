#!/bin/bash

# Function to check if a file requires dos2unix conversion 
# and convert it only if the check passes.
# Usage: check_and_convert_to_unix <filename>
check_and_convert_to_unix() {
    local FILE="$1"

    # Check if the file exists
    if [ ! -f "$FILE" ]; then
        echo "Error: File '$FILE' not found." >&2
        return 1
    fi

    # Use dos2unix info mode (-i) to check for conversion need 
    # (-c prints the filename only if it would be converted)
    if dos2unix -ic "$FILE" | grep -q "$FILE"; then
        echo "Converting '$FILE' from DOS (CRLF) to Unix (LF)..."
        # Execute the conversion
        dos2unix "$FILE"
        if [ $? -eq 0 ]; then
            echo "Conversion successful for '$FILE'."
        else
            echo "Error: dos2unix conversion failed for '$FILE'." >&2
            return 1
        fi
    else
        echo "'$FILE' is already in Unix format or doesn't need conversion."
    fi
    return 0
}

# Check and convert .env file if it exists
if [ -f ./.env ]; then
    check_and_convert_to_unix .env
fi

# Check and convert compile.sh
check_and_convert_to_unix compile.sh
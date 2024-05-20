#!/bin/bash

# Check if the number of arguments is exactly 1
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <N>"
    exit 1
fi

# Check if the argument is a positive integer
if ! [[ "$1" =~ ^[0-9]+$ ]]; then
    echo "Error: N must be a positive integer"
    exit 1
fi

# Number of random entries to generate
N=$1

# Output file
output_file="measurements.txt"

# Clear the file if it exists
> "$output_file"

# Function to generate a random 32-character string
generate_random_string() {
    local chars='abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
    local random_string=""
    for i in {1..32}; do
        random_char=${chars:RANDOM%${#chars}:1}
        random_string+=$random_char
    done
    echo "$random_string"
}

# Generate N random entries and save them to the file
for ((i=0; i<N; i++)); do
    # Generate a random 32-character string
    random_string=$(generate_random_string)

    # Generate a random float with one decimal place
    random_float=$(printf "%.1f\n" "$(echo "scale=1; $((RANDOM % 1201 - 600))/10" | bc)")

    # Combine and write to the file
    echo "$random_string: $random_float" >> "$output_file"
done

echo "Generated $N random entries in $output_file"
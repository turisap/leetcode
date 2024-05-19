#!/bin/bash

# Check if an argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <number_of_integers>"
  exit 1
fi

# Number of random integers to generate
N=$1

# Check if N is a positive integer
if ! [[ "$N" =~ ^[0-9]+$ ]]; then
  echo "Error: Argument must be a positive integer."
  exit 1
fi

# Output file
OUTPUT_FILE="measurements.txt"

# Function to generate a random integer between -60 and 60
generate_random_integer() {
  echo $(( RANDOM % 121 - 60 ))
}

# Generate N random integers and save to the file
> "$OUTPUT_FILE" # Truncate the file if it exists, or create it if it doesn't
for (( i=0; i<N; i++ ))
do
  generate_random_integer >> "$OUTPUT_FILE"
done

echo "Generated $N random integers between -60 and 60 and saved them to $OUTPUT_FILE."
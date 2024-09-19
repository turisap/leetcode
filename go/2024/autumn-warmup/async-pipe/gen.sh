#!/bin/bash

# Output file
output_file="generated_data_1000.csv"

# Write header to the file
echo "integer,float1,float2" > $output_file

# Generate 1000 rows of data
for i in $(seq 1 1000); do
  float1=$(echo "$i + 0.1" | bc)
  float2=$(echo "$i + 0.2" | bc)
  echo "$i,$float1,$float2" >> $output_file
done

echo "CSV file generated: $output_file"

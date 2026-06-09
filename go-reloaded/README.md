# go-reloaded# my project

main.go     (Entry point + main logic) always the entry point

transform.go    (All transform functions)  Feature-specific files

helpers.go      (Helper functions) Utility functions

input.txt       (Test input)

output.txt      (Generated output)

# The main.go

This file handles reading/writing files and orchestrating transformations.

## What this file does:

It Reads arguments from user
It Reads input file
It Calls Transform function (from another file)
It Writes output file

# transform.go

This file has all the transformation functions.

## What this file does:

It Contains all transformation functions
Each function is named with capital first letter (public in Go)
It Functions call helper functions from helpers.go


# helpers.go

This file has helper functions.

## What this file does:

It Contains small transformation functions
It Reusable functions that other files call
It Keeps code organized and clean

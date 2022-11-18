#!/bin/sh

# This script is used to run all tests and generate coverage reports for the
# current go project. It is intended to be run from the root of the project
# directory.

echo "Running tests and generating coverage report..."
go test -coverprofile=coverage.out

echo "Opening coverage report..."
go tool cover -html=coverage.out

echo "Done."
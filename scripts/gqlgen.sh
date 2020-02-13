#!/bin/bash
printf "\nRegenerating gqlgen files\n"
# Optional, delete the resolver to regenerate, only if there are new queries
# or mutations, if you are just changing the input or type definition and
# doesn't impact the resolvers definitions, no need to do it
rm -f internal/gql/resolvers/generated/generated.go
time go run -v github.com/99designs/gqlgen
printf "\nDone.\n\n"
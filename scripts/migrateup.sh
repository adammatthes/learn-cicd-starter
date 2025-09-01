#!/usr/bin/env bash

#if [ -f .env ]; then
#    	set -a
#	source .env
#	set +a
#else
#	echo ".env not found at repo root" >&2
#	exit 1
#fi

#export GOOSE_DRIVER=sqlite3
#export GOOSE_DBSTRING="$DATABASE_URL"

#if [ -z "${DATABASE_URL:-}" ]; then
#	echo "DATABASE_URL is empty" >&2
#	exit 1
#fi

#echo "$GOOSE_DRIVER"
#echo "$GOOSE_DBSTRING"

#goose -v -dir . status
#goose -v -dir sql/schema up

#! /usr/bin/env sh

set -e

umask 002

# Release phase
python migration_manager.py db upgrade

if [ "$FLASK_ENV" = "development" ]; then
    while true; do
        python3 main.py
        echo "'main.py' crashed with exit code $?. Restarting..." >&2
        sleep 1
    done
else
    exec gunicorn -k eventlet -c "$GUNICORN_CONF" "$APP_MODULE"
fi

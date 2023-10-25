#!/bin/bash
print_usage() {
    echo "Usage: start.sh [-r rebuild]"
}

rebuild=false

while getopts "hr" FLAG; do
    case "${FLAG}" in
        h) 
            print_usage
            exit 0
            ;;
        r) 
            rebuild=true
            ;;
        \?)
            exit 1
    esac
done

if [ ! -e ./bin/peer ] || $rebuild
    then
        echo "Building project..."
        ./build.sh
elif [ ! -x ./bin/peer ]
    then
        chmod +x ./bin/peer
fi

cd ./bin && ./peer

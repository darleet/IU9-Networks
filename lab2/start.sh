#!/bin/bash
print_usage() {
    echo "Usage: start.sh [-r rebuild]";
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

if [ ! -e ./bin/server ] || $rebuild;
    then
        echo "Building project...";
        ./build.sh;
elif [ ! -x ./bin/server ];
    then
        chmod +x ./bin/server;
fi

cd ./bin && LOGXI=* LOGXI_FORMAT=pretty,happy ./server;
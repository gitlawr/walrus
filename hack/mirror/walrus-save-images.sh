#!/bin/bash
list="walrus-images.txt"
images="walrus-images.tar.gz"
source_registry=""

usage () {
    echo "USAGE: $0 [--image-list walrus-images.txt] [--images walrus-images.tar.gz]"
    echo "  [-s|--source-registry] source registry to pull images from in registry:port format."
    echo "  [-l|--image-list path] text file with list of images; one image per line."
    echo "  [-i|--images path] tar.gz generated by docker save."
    echo "  [-h|--help] Usage message"
}

POSITIONAL=()
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -i|--images)
        images="$2"
        shift # past argument
        shift # past value
        ;;
        -l|--image-list)
        list="$2"
        shift # past argument
        shift # past value
        ;;
        -s|--source-registry)
        source_registry="$2"
        shift # past argument
        shift # past value
        ;;
        -h|--help)
        help="true"
        shift
        ;;
        *)
        usage
        exit 1
        ;;
    esac
done

if [[ $help ]]; then
    usage
    exit 0
fi

source_registry="${source_registry%!/(MISSING)}"
if [ ! -z "${source_registry}" ]; then
    source_registry="${source_registry}/"
fi

pulled=""
while IFS= read -r i; do
    [ -z "${i}" ] && continue
    i="${source_registry}${i}"
    if docker pull "${i}" > /dev/null 2>&1; then
        echo "Image pull success: ${i}"
        pulled="${pulled} ${i}"
    else
        if docker inspect "${i}" > /dev/null 2>&1; then
            pulled="${pulled} ${i}"
        else
            echo "Image pull failed: ${i}"
        fi
    fi
done < "${list}"

echo "Creating ${images} with $(echo ${pulled} | wc -w | tr -d '[:space:]') images"
docker save $(echo ${pulled}) | gzip --stdout > ${images}

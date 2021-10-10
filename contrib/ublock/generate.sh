#!/bin/bash
set -uex -o pipefail
ROOT=$(cd ../..; pwd)
ruby generate.rb < ${ROOT}/domain-list.yml > uBlockOrigin.txt

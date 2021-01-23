#!/usr/bin/env python3

import tempfile
import os
import re

def convert():
    with open('./uBlacklist.txt') as f:
        current_dir = os.getcwd()
        with tempfile.NamedTemporaryFile("w+", dir=current_dir, delete=False, suffix='.css') as tmpf:
            for url in f.read().splitlines():
                # ext: *://code-examples.net/* -> code-examples.net
                matched_url = re.search('((?<=\*:\/\/\*\.)\w.+|(?<=\*:\/\/)\w.+)', url).group()[:-2]
                style = f'a[href*="{matched_url}"]{{display:none;}}'
                tmpf.write(style)
            os.rename(tmpf.name, os.path.join(current_dir, 'uBlacklist.css'))

if __name__ == '__main__':
    convert()

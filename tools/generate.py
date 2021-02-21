import argparse
import sys

import yaml
from yaml import BaseLoader as Loader


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', nargs='?',
                        type=argparse.FileType('r'), default=sys.stdin)
    args = parser.parse_args()

    f = args.input
    data = yaml.load(f, Loader=Loader)


if __name__ == '__main__':
    main()

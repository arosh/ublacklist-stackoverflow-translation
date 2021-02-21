import yaml
import sys
from collections import OrderedDict

def main():
    data = []
    with open('uBlacklist.txt') as f1, open('evidence.md') as f2:
        for line1, line2 in zip(f1, f2):
            sp = line2.split('|')
            ev = sp[2].strip(' `')
            origin = sp[3].strip(' `')

            o = dict()
            o['pattern'] = line1.rstrip()
            o['evidence'] = ev
            o['original'] = origin
            if len(sp) >= 5:
                note = sp[4].strip()
                if note:
                    o['note'] = note
            data.append(o)
    yaml.dump(data, sys.stdout, sort_keys=False)

if __name__ == '__main__':
    main()

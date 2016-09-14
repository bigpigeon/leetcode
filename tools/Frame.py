import sys
import os


def make_frame(s):
    if isinstance(s, str):
        s = s.decode('utf-8')
    strs = s.replace('\r\n', '\n').split('\n')
    max_length = 0
    height = len(strs)

    for __s in strs:
        if len(__s) > max_length:
            max_length = len(__s)
    new_s = '+' + '-'*max_length + '+\n'
    for __s in strs:
        new_s += '|' + __s + ' ' * (max_length - len(__s)) + '|\n'
    new_s += '+' + '-'*max_length + '+\n'
    return new_s

if __name__ == '__main__':
    if len(sys.argv) > 1:
        f = file(sys.argv[1])
        print make_frame(f.read())
    else:
        f = sys.stdin
        print make_frame(f.read())

#! /bin/env python3

import requests
import sys
import json

def pn_verify(addr, identifier):
    # expand URL
    if addr[0] == ':':
        addr = 'http://localhost' + addr
    # add verification path and id according to Peernotify API
    url_path = addr.strip(' /') + '/verify/' + identifier
    # send request
    try:
        r = requests.get(url_path)
        if not r.status_code == 200:
            print('[ERROR]: ' + r.reason)
    except Exception as e:
        print('[ERROR]: ' + str(e))

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print('[USAGE]:\n\tverify.py <addr> <id>')
    else:
        pn_verify(sys.argv[1], sys.argv[2])
    

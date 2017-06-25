#! /bin/env python3

import requests
import sys
import json


def pn_register(addr, email):
    # construct contact JSON description
    contact = {"methods": [{"protocol": "SMTP", "address": email}]}
    # expand URL
    if addr[0] == ':':
        addr = 'http://localhost' + addr
    # add registration path according to Peernotify API
    url_path = addr.strip(' /') + '/register'
    # send request
    try:
        r = requests.post(url_path, data=json.dumps(contact))
        if not r.status_code == 200:
            print('[ERROR]: ' + r.reason)
    except Exception as e:
        print('[ERROR]: ' + str(e))


if __name__ == '__main__':
    if len(sys.argv) != 3:
        print('[USAGE]:\n\tregister.py <addr> <key> <email>')
    else:
        pn_register(sys.argv[1], sys.argv[2])

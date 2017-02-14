#! /bin/env python3

import requests
import sys
import json

def pn_forward(addr, token, payload):
    # construct message JSON description
    message = {"token" : token, "payload" : payload}
    # expand URL
    if addr[0] == ':':
        addr = 'http://localhost' + addr
    # add forwarding path according to Peernotify API
    url_path = addr.strip(' /') + '/forward'
    # send request
    try:
        r = requests.post(url_path, data=json.dumps(message))
        if not r.status_code == 200:
            print('[ERROR]: ' + r.reason)
    except Exception as e:
        print('[ERROR]: ' + str(e))

if __name__ == '__main__':
    if len(sys.argv) != 4:
        print('[USAGE]:\n\tforward.py <addr> <token> <msg>')
    else:
        pn_forward(sys.argv[1], sys.argv[2], sys.argv[3])
    

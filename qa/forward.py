#! /bin/env python3

import requests
import sys
import json
import base58
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import hashes


IDKEY_SIZE = 32


def pn_forward(addr, keyset, payload):
    # construct message JSON description
    token = generate_token(keyset)
    message = {"token": token, "payload": payload}
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
        else:
            print('[TOKEN]: ' + token)
    except Exception as e:
        print('[ERROR]: ' + str(e))


def generate_token(keyset):
    data = base58.b58decode(keyset)
    id_key, root_key = data[:IDKEY_SIZE], data[IDKEY_SIZE:]
    digest = hashes.Hash(hashes.SHA256(), backend=default_backend())
    digest.update(root_key)
    return base58.b58encode(id_key + digest.finalize())


if __name__ == '__main__':
    if len(sys.argv) != 4:
        print('[USAGE]:\n\tforward.py <addr> <token> <msg>')
    else:
        pn_forward(sys.argv[1], sys.argv[2], sys.argv[3])

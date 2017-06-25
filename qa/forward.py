#! /bin/env python3

import requests
import sys
import json
from base58 import b58encode, b58decode
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import hashes


ID_SIZE = 8
KEY_SIZE = 32
MASK_SIZE = 32


def pn_forward(addr, keyset, payload):
    # construct message JSON description
    token, new_keyset = generate_token(keyset)
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
            print('[API ERROR]: ' + r.reason)
        else:
            print('[NEW KEYSET]: ' + new_keyset)
    except Exception as e:
        print('[EXCEPTION]: ' + str(e))


def generate_token(keyset):
    data = b58decode(keyset)
    key_end = ID_SIZE + KEY_SIZE
    id, key, mask = data[:ID_SIZE], data[ID_SIZE:key_end], data[key_end:]
    new_key, new_mask = sha256hash(key), sha256hash(mask)
    token = id + bytes([new_key[i] ^ new_mask[i] for i in range(MASK_SIZE)])
    new_keyset = id + new_key + new_mask
    return b58encode(token), b58encode(new_keyset)


def sha256hash(data):
    digest = hashes.Hash(hashes.SHA256(), backend=default_backend())
    digest.update(data)
    return digest.finalize()


if __name__ == '__main__':
    if len(sys.argv) != 4:
        print('[USAGE]:\n\tforward.py <addr> <token> <msg>')
    else:
        pn_forward(sys.argv[1], sys.argv[2], sys.argv[3])

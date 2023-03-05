#!/bin/python3
import codecs, base64, json

def eth_to_fil(eth_key):
    private_key_bytes = codecs.decode(eth_key[2:], 'hex')
    json_sha = "{\"Type\":\"secp256k1\",\"PrivateKey\":\"%s\"}" % base64.b64encode(private_key_bytes).decode()
    print(bytes(json_sha, 'utf-8'))
    fil_key = bytes(json_sha, 'utf-8').hex()
    return fil_key

def fil_to_eth(fil_key):
    json_sha = bytes.fromhex(fil_key).decode('utf-8')
    private_key_sha = json.loads(json_sha)['PrivateKey']
    private_key_bytes = base64.b64decode(private_key_sha)
    private_key = '0x' + codecs.encode(private_key_bytes, 'hex').decode('utf-8')
    return private_key


def main():
    eth_key = '0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9'
    fil_key = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22592b496445503151465632376f4f665439304d6151417545744d4b7348754f496376676b5350342b7a376b3d227d"
    print("eth_key: %s to fil_key: %s" % (eth_key, eth_to_fil(eth_key)))
    print("fil_key: %s to eth_key: %s" % (fil_key, fil_to_eth(fil_key)))

if __name__ == '__main__':
    print("start...")
    main()
    print("end.....")

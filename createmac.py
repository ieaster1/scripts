#!/usr/bin/env python3
import random
import argparse

def create_random_mac(type='qemu'):
    ouis = { 'xen': [ 0x00, 0x16, 0x3E ], 'qemu': [ 0x52, 0x54, 0x00 ] }

    try:
        oui = ouis[type]
    except KeyError:
        oui = ouis['qemu']

    decimal_mac = oui + [ random.randint(0x00, 0xff), random.randint(0x00, 0xff), random.randint(0x00, 0xff) ]
    mac = ':'.join(map(lambda x: "%02x" % x, decimal_mac))
    return decimal_mac, mac

def main():
    parser = argparse.ArgumentParser(description='Create random MAC with Locally Administered Organizational Unique Identifier.')
    parser.add_argument('-c', '--count', help='Number of MACs to create.', type=int, metavar='count')

    args = vars(parser.parse_args())

    # print(f"[debug]: Args {args}")
    if args['count'] is None:
        count = 1
    else:
        count = args['count']

    # print(f"[debug]: create {count} macs")
    while count > 0:
        decimal_mac, mac = create_random_mac()
        print(mac)
        count -= 1

if __name__ == '__main__':
    main()

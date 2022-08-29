#!/usr/bin/env python3
import random
import argparse

def create_random_mac(type='qemu'):
    ouis = { 'xen': [ 0x00, 0x16, 0x3E ], 'qemu': [ 0x52, 0x54, 0x00 ] }

    try:
        oui = ouis[type]
    except KeyError:
        oui = ouis['qemu']

    decimal_mac = oui + random.sample(range(0x00, 0xff), 3)
    mac = ':'.join(map(lambda x: "%02x" % x, decimal_mac))
    return decimal_mac, mac

def main():
    parser = argparse.ArgumentParser(prog='createmac',
        description='Create random MAC with Locally Administered Organizational Unique Identifier.')
    parser.add_argument('-c',
                        '--count',
                        help='Number of MACs to create.',
                        type=int, metavar='count')

    args = vars(parser.parse_args())

    if args['count'] is None:
        count = 1
    else:
        count = args['count']

    while count > 0:
        _, mac = create_random_mac()
        print(mac)
        count -= 1

if __name__ == '__main__':
    main()

#!/usr/bin/env python3
import re
import sys
import random
import argparse

def create_random_mac(type='qemu'):
    ouis = { 'xen': [ 0x00, 0x16, 0x3E ], 'qemu': [ 0x52, 0x54, 0x00 ] }

    if hasattr(create_random_mac, 'oui'):
        str_oui = create_random_mac.oui.split(':')
        oui = [int(i, 16) for i in str_oui]
    else:
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
                        help='number of MACs to create',
                        type=int,
                        metavar='')
    parser.add_argument('-o',
                        '--oui',
                        help='input for OUI, e.g., "00:12:ac"',
                        type=str,
                        metavar='')

    args = vars(parser.parse_args())

    if args['oui'] is not None:
        if re.match("[0-9a-f]{2}([:]?)[0-9a-f]{2}(\\1[0-9a-f]{2})$", args['oui'].lower()):
            create_random_mac.__setattr__('oui', args['oui'])
        else:
            sys.exit(f"{args['oui']} is incorrect format, check help for details")

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

# Networking

## igmp

### One time useage
```bash
wget https://github.com/ieaster1/scripts/raw/main/networking/igmp -O igmp && chmod +x igmp
```

```bash
./igmp
```

### install

```bash
wget https://github.com/ieaster1/scripts/raw/main/networking/igmp -O /usr/local/bin/igmp && chmod +x /usr/local/bin/igmp
```

**example**
```bash
$ igmp

Interface: lo
  Address: 127.0.0.1/8
  Multicast Group: 224.0.0.251
  Multicast Group: 224.0.0.1
  IGMP Groups:
    Group: 2, Version: V3

Interface: eno1
  Address: 192.168.1.1/24
  Multicast Group: 224.0.0.251
  Multicast Group: 224.0.0.1
  IGMP Groups:
    Group: 2, Version: V2

Interface: eno2
  Address: 172.1.1.1/24
  Multicast Group: 239.1.1.1
  Multicast Group: 239.2.2.2
  Multicast Group: 224.0.0.251
  Multicast Group: 224.0.0.1
  IGMP Groups:
    Group: 4, Version: V3
```

## createmac
createmac will create a random MAC using a QEMU OUI by default.  You may pass in a OUI to generate a random MAC.

### Setup
```bash
wget https://github.com/ieaster1/scripts/raw/main/networking/createmac -O createmac && chmod +x createmac
```

**example**
```bash
$ createmac -h
usage: createmac [-h] [-c] [-o]

Create random MAC with Locally Administered Organizational Unique Identifier.

optional arguments:
  -c int
        number of MACs to create (default 1)
  -o string
        input for OUI, e.g., "00:12:ac"
```

```bash
$ createmac
52:54:00:37:26:3e
```

```bash
$ createmac -c 5 -o "00:ac:12"
00:ac:12:d1:96:f1
00:ac:12:3b:33:6a
00:ac:12:74:63:c4
00:ac:12:0f:dc:83
00:ac:12:4e:f6:77
```

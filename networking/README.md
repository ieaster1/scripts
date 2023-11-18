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

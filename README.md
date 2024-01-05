# tic-linky

## Vérifier que ça marche

```bash
screen /dev/ttyAMA0 9600,cs7,parenb,-parodd,-cstopb
```

## Paramètres du port série

- 1 bit de start (0 logique)
- 7 bits
- 1 bit de parité pair
- 1 bit de stop (1 logique)

> Activer le port série sur le rasp : https://www.framboise314.fr/utiliser-le-port-serie-du-raspberry-pi-3-et-du-pi-zero/

## Liste des champs

### Date

- DATE

### Indexes

- EAST
- EASF01
- EASF02
- EASF03
- EASF04
- EASF05
- EASF06
- EASF07
- EASF08
- EASF09
- EASF10
- EASD01
- EASD02
- EASD03
- EASD04

### Puissance

- SINSTS
- URMS1
- IRMS1

### Horodatés

- SMAXSN
- CCASN
- UMOY1

### Annexes

- STGE
- MSG1

## Variables d'environement

- `INFLUXDB_URL="http://mini-pc:8086"`
- `INFLUXDB_USERNAME="tic-linky"`
- `INFLUXDB_PASSWORD="tic-linky"`
- `INFLUXDB_DATABASE="tic-linky"`

## Références

- https://hallard.me/pitinfov12-light/
- https://github.com/JordanMartin/linky-teleinfo/blob/master/README.md
- https://github.com/hallard/teleinfo
- https://www.enedis.fr/media/2035/download

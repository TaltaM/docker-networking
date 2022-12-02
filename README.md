# Build

```
./build.sh
```

# Run with default networking

```
docker compose --profile default-profile up
```

# Run with ipvlan networking

```
docker compose --profile vlan-profile up
```

# Clean up

```
docker compose down
```
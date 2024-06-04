# Sprune

The goal of this project is to prune a Tendermint/CometBFT database of blocks and a Cosmos SDK application database, retaining only the last X versions. This prevents the need for state sync every few days.

This tool works with a subset of modules. While an application may have modules outside the scope of this tool, Sprune will prune the default SDK module and custom application module.

## How to Use
Sprune works on a data directory with the same structure as a normal Cosmos SDK/Tendermint node. By default, it will prune all but 100 blocks from Tendermint and all but 10 versions of application state.

> Note: Application pruning can take a very long time depending on the size of the database.

```
# Clone & build Sprune repo
git clone https://github.com/cakra-labs/sprune
cd sprune
make build

# Stop daemon/cosmovisor
sudo systemctl stop cosmovisor

# Init config
./build/sprune config

# Update config
nano $HOME/.sprune/config.yaml

# Run pruner
./build/sprune prune
```

Flags:
- `data-dir`: set home directory for app config (default "$HOME/.sprune")
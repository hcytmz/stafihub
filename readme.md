# stafihub
**stafihub** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

### Get started for dev

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure for dev

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).


### Running node for production
1. build the chain. `starport build --release -t linux:amd64 -t darwin:amd64 -t darwin:arm64`
2. unzip the release/stafihub_darwin_amd64.tar.gz file to get a execute file, like stafihubd
1. Setting up the keyring, commands:
    - `KEYPASSWD=123456789`
    - `./release/stafihubd config keyring-backend file`
    - `(echo $KEYPASSWD; echo $KEYPASSWD) | ./release/stafihubd keys add kael --keyring-backend file`
    - `MY_VALIDATOR_ADDRESS=$(./release/stafihubd show kael -a --keyring-backend file)`
2. Prepare genesis file and other config files
    - `./release/stafihubd init mynode -o`
    - `./release/stafihubd add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000stake`
    - `./release/stafihubd gentx kael 100000000stake --keyring-backend file --chain-id stafihub`
    - `./release/stafihubd collect-gentxs`
3. Start node: `./release/stafihubd start`
4. There might be this kind of error:
`Error: couldn't get client config: open ~/.stafihub/config/client.toml: permission denied`, just add sudo to your command.
for example, start node: `sudo ./release/stafihubd start`

### Settings

global:

```
stafihubd tx ledger set-protocol-fee-receiver stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx ledger set-relay-fee-receiver uratom stafi1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-stafihub --keyring-backend file

# default 0.1
stafihubd tx ledger set-staking-reward-commission uratom 0.1 --from admin --chain-id local-stafihub --keyring-backend file

# default 0.002
stafihubd tx ledger set-unbond-commission uratom 0.002 --from admin --chain-id local-stafihub --keyring-backend file

# set fis metadata
stafihubd tx sudo add-denom --metadata ./metadata_example.json --chain-id local-stafihub --from admin --keyring-backend file

```

add new rtoken:

```
# set rtoken metadata
stafihubd tx sudo add-denom --metadata ./metadata_example.json --chain-id local-stafihub --from admin --keyring-backend file

# default 1000000ufis
stafihubd tx ledger set-unbond-relay-fee uratom 1000000ufis --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx ledger set-chain-bonding-duration uratom 2 --chain-id local-stafihub --from admin --keyring-backend file

stafihubd tx ledger set-pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx ledger set-init-bond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 uratom --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx relayers add-relayers uratom stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:stafi1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --keyring-backend file --from admin --chain-id local-stafihub

stafihubd tx relayers update-threshold uratom 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd tx ledger set-r-params uratom 0.00001stake 600 0 cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre:cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9nprd --from admin --keyring-backend file --chain-id local-stafihub

```


bridge:

```
stafihubd tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd tx bridge add-relayer 1 stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-stafihub

stafihubd tx bridge set-threshold 1 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 ufis --from admin --keyring-backend file --chain-id local-stafihub

stafihubd tx bridge set-relay-fee-receiver stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-stafihub

stafihubd tx bridge set-relay-fee 1 1000000ufis --from admin --keyring-backend file --chain-id local-stafihub
```

### user operate

iquidity bond:

```
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --memo 1:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

recover:

```
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --memo 2:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```


liquidity unbond:

```
stafihubd tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uratom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/stafi/rtoken-relay-core/keys/stafihub --chain-id local-stafihub
```

deposit:
```
stafihubd tx bridge deposit 1 000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 8000000 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-stafihub --keyring-backend file
```

# Web3 API (GO)
A full Web3 API that allows users to query and interact with the blockchain, as well as generate new wallets.

## API URL
You can test this API for yourself by using [this](https://go-api-378801.wl.r.appspot.com/) base link.
Add the desired endpoint and necessary parameters.

## Endpoints
* /[get_block](https://github.com/aglawson/Web3API-Go/blob/main/api/GetBlock.go)
  * Params: 
     * network - string representing desired network (ex. mainnet, goerli, polygon-mainnet)
  * Returns the latest block on Ethereum mainnet
* /[get_balance](https://github.com/aglawson/Web3API-Go/blob/main/api/GetBalance.go)
  * Params: 
     * network - string representing desired network (ex. mainnet, goerli, polygon-mainnet)
     * wallet - string wallet address
  * Returns the balance of provided wallet address (wallet)
* /[create_wallet](https://github.com/aglawson/Web3API-Go/blob/main/api/GenerateWallet.go)
  * Returns the private key and public address of a newly generated wallet
  * NOTE: do not use except on a secure local server.
  
* /[get_gas_price](https://github.com/aglawson/Web3API-Go/blob/main/api/GetGasPrice.go)
  * Params:
    * network - string representing desired network (ex. mainnet, goerli, polygon-mainnet)
  * Returns the current gas price of a given network

* /[get_token_balance](https://github.com/aglawson/Web3API-Go/blob/main/api/GetTokenBalance.go)
  * Params:
    * network - string representing desired network (ex. mainnet, goerli, polygon-mainnet)
    * wallet - address of asset owner
    * contract - smart contract of asset
      * NOTE: can be either an NFT or an ERC20 token
  * Returns the current number of assets under contract that wallet owns on network
  
* /[get_token_supply](https://github.com/aglawson/Web3API-Go/blob/main/api/GetTokenSupply.go)
  * Params:
    * network - string representing desired network (ex. mainnet, goerli, polygon-mainnet)
    * contract - smart contract of asset
      * NOTE: can be either NFT or ERC20 token

* /[is_token_holder](https://github.com/aglawson/Web3API-Go/blob/main/api/IsTokenHolder.go)
  * Params:
    * wallet - address of wallet in question
    * network - string representing desired network (ex. mainnet, goerli, polygon-mainnet)
    * contract - smart contract of asset
      * NOTE: can be either NFT or ERC20 token
  * Returns boolean value. If wallet owns 1 or more tokens from contract -> true; otherwise -> false;

* /[get_coin_price](https://github.com/aglawson/Web3API-Go/blob/main/api/GetCoinPrice.go)
  * Params:
    * from: name of coin to inquire about (supported: 'ethereum')
    * to: name of currency to convert to (supported: 'usd', 'vnd')
  * Returns float representing the amount of 'to' that equates to 1 'from'

* /[write_coin_price](https://github.com/aglawson/Web3API-Go/blob/main/api/WriteCoinPrice.go)
  * Params:
    * from: name of coin to inquire about (eg. 'ethereum')
    * to: name of currency to convert to (eg. 'usd', 'vnd')
  * Writes to connected Firestore DB the current amount of 'to' that equates to 1 'from'
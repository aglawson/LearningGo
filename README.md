# Web3 API (GO)
A full Web3 API that allows users to query and interact with the blockchain, as well as generate new wallets.

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

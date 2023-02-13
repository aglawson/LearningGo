# Web3 API (GO)
A full Web3 API that allows users to query and interact with the blockchain, as well as generate new wallets.

## Endpoints
* /get_block
  * Returns the latest block on Ethereum mainnet
* /get_balance
  * Returns the balance of provided wallet address (wallet)
* /create_wallet
  * Returns the private key and public address of a newly generated wallet
  * NOTE: do not use except on a secure local server.
  

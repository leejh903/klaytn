// Copyright 2018 The klaytn Authors
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
//
// This file is derived from core/blockchain.go (2018/06/04).
// Modified and improved for the klaytn development.

/*
Package blockchain implements the canonical chain and the state transition model of Klaytn.

A Klaytn node starts to generate a chain from a genesis block and appends validated blocks transferred from other nodes or generated by the node.
Each block containing transactions is validated and processed by each node transiting the global state of the system.
The state transition containing EVM execution is triggered by transactions that containing transition messages for the state.
Transactions are delivered by RPC servers or other nodes, managed by the tx pool and aggregated to a block.
As well as the exact transaction process, Klaytn achieved the efficiency of the transaction processing through a caching logic.
The cacheSender concurrently recovers and caches an address from a signature in the transaction which was one of the main bottlenecks of transaction processing.

Source files

Each file provides the following features
 - bad_blocks.go : keeps block hashes of bad blocks which usually for hard forks.
 - block_validator.go : implements BlockValidator which is responsible for validation block headers and the processed state.
 - blockchain.go : implements the canonical chain of blocks and managing functions to support imports, reverts and reorganisations.
 - chain_indexer.go : implements ChainIndexer.
 - chain_makers.go : generates temporary blocks or chains to support SimulatedBackend.
 - error.go : defines errors frequently used in blockchain package.
 - events.go : defines event structs delivered between go-routines.
 - evm.go : creates an EVM with a given context for use.
 - gaspool.go : defines GasPool which tracks and manages the amount of gas available during the transaction execution.
 - gen_genesis.go : is auto-generated code by gencodec to marshal/unmarshal Genesis as/from JSON.
 - gen_genesis_account.go : is auto-generated code by gencodec to marshal/unmarshal GenesisAccount as/from JSON.
 - genesis.go : defines Genesis which specifies values of a genesis block and initial settings of the chain.
 - genesis_alloc.go : contains the genesis allocation of built-in genesis blocks.
 - headerchain.go : implements HeaderChain which makes a chain with block headers.
 - init_derive_sha.go : initialize a DeriveSha function with a specific type.
 - metrics.go : contains metrics used for blockchain package.
 - mkalloc.go : creates the genesis allocation constants in genesis_alloc.go.
 - state_processor.go : implements StateProcessor which takes care of transitioning state.
 - state_transition.go : implements a state transaction model worked with messages in transactions.
 - tx_cacher.go : recovers senders of transactions from signatures and caches the sender address.
 - tx_journal.go: keeps logs of transactions created by the local node.
 - tx_list.go : provides sorted map and list structures used in tx pool manipulation.
 - tx_pool.go : contains all currently known transactions and manages the transactions.
 - types.go : interfaces Validator and Processor which validate or process block data.
*/
package blockchain

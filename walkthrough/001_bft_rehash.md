# byzantine fault tolerance rehash

Byzantine Generals Problem
https://en.wikipedia.org/wiki/Byzantine_fault

TLDR, 2/3rd of nodes are required to move foward with consensus

there are three main phases of bft
- request
- pre prepare
- prepare
- commit
- reply

## proposal

a responsible node proposes a block based on application conditions and constraints (round robin, determined order, timeout, etc)

broadcasts proposed block to known peers

## pre commit

nodes that received the proposal verify requirements
- is this block (and the txs contained within it) valid
- am i synced enough state wise to validate
- is the proposer of this block eligible to propose (validators / service providers )

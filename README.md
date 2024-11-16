# GO Hexagonal Architecture

## Features
1. Can switch database easily

## Rules
1. Every model should be in its own repository
2. Validate request that don't need to access database at app layer
3. Validate request that need to access database at core layer
4. Return custom error, log real error

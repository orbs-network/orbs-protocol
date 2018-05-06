# Virtual Machine

Executes smart contracts using various processors (languages) and produces state difference as a result.

&nbsp;
## `Processors`

#### Native processor
* Native set of smart contracts included in node core.
* Hard coded in code base and updated through node version updates.
* Implemented in native language of the node for efficiency and performance.
* Contracts are stored in a registry which is loaded on initialization.
* Every contract and method contain configuration for execution.
  * Permission on who can execute the contract externally.
* See contract API [specification](../smart-contracts/native.md).

&nbsp;
## `Data Structures`

#### Transient state cache
* Stores temporary state changes during transaction set execution (until state diff is committed).
* Format is identical to the state store data structure in `StateStorage`.
* No need to be persistent since a new instance is created per execution.
* No limit on max size.

&nbsp;
## `CallContract` (method)
> Execute a read only method in a native contract and return its result

#### Prepare for execution
* Retrieve the contract from the native contract registry.
* Validate that the method exists and has appropriate permissions for execution.

#### Execute
* Contract code cannot update state (state writes must throw an error).
* Execute the contract method and return its result.

&nbsp;
## `ProcessTransactionSet` (method)
> Process a group of transactions together that update state and return the combined state diff

#### Prepare for execution
* Prepare transient state cache that will hold updated state temporarily (across all transactions).
* Retrieve contracts from the native contract registry.
* Validate that the methods exist and have appropriate permissions for execution.

#### Execute
* Contract code can only update state within the transient state cache.
* Execute the contract methods and return their results.

#### Prepare combined state diff
* Retrieve the changes between the transient state cache and the original state.
* Encode changes as state diff.
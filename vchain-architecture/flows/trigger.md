# Trigger

A trigger is a method call that is triggered by the system without a need for an external call.\
Triggers are implemented similarly to regular transaction and generate a receipt, however they do not require a signed client transaction.\
The trigger transaction calls the trigger contract post message.\
Note: The trigger transaction uses a "protocol special" empty signature which is reserved for system. The `triggers` contract should fail if the caller signer is not empty.

## Trigger Transaction
* protocol_version = protocol version
* virtual_chain_id = virtual_chain_id
* timestamp = block.timestamp
* signer = empty
* contract_name = `triggers`
* method_name = `blockPost`
* input_argument_array = empty
* signature = empty

## Flow

#### Participant Services
* `ConsensusContext` - Generates and validates the trigger transaction. Remove trigger transactions from ordering validation by the `TransactionPool`.

#### Trigger Transaction Generation
* If the configuration flag `TRIGGERS_ENABLED` is set to `TRUE`, the `ConsensusContext` generates the trigger transaction as part of `RequestNewTransactionsBlock`.

#### Trigger Transaction Validation
* If the configuration flag `TRIGGERS_ENABLED` is set to `TRUE`, the `ConsensusContext` verifies that a trigger transaction is present and valid, as part of `ValidateTransactionsBlock`. 
* If the configuration flag `TRIGGERS_ENABLED` is set to `FALSE`, the `ConsensusContext` verifies that a trigger transaction is NOT present in the block.

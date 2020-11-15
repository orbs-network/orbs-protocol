# `_Deployments` System Service

The system service charged with deploying new smart contracts (services or libraries) onto a virtual chain.

Implemented on the `Native` processor.

To retrieve the code, `getCodeParts()` in combination with `getCodePart()` should be used - `getCode()` is equivalent to `getCodePart(serviceName, 0)` and only works with single-file contracts.

#### Permissions
* This is a system contract that runs under `System` permissions.

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (potentially changes state).

#### Behavior
* Must be empty to avoid chicken and egg problem since used to deploy all other contracts.

&nbsp;
## `getInfo` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Takes service name as argument.
* For queries about self (service name `_Deployments`) respond with the native processor (chicken and egg).
* Read processor type from state key `<name>.Processor` by calling `State.ReadUint32ByKey`.
* If an empty value is read, contract is not deployed.

&nbsp;
## `deployService` (method)

Deploys the smart contract for the first time.

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (potentially changes state).

#### Behavior
* Takes service name, processor type and code as arguments.
* Has ellipsis syntax, meaning that additional code can be passed as an argument: `deployService(serviceName, processorType, code1, code2, code3, code4)`
* Make sure the **service isn't already deployed** by calling `getInfo`.
* Write processor type to state under key `<name>.Processor` by calling `State.WriteUint32ByKey`.
* Write first part of the code to state under key `<name>.Code` by calling `State.WriteBytesByKey`.
* Write other parts of the code to state under key `<name>.Code.<counter>` by calling `State.WriteBytesByKey`.
* Every time code part is written into state, increment `<name>.CodeParts` counter (except the first time for backwards compatibility)
* Init the service by calling `Service.CallMethod` with method `_init`.
  * This is an `Internal` method, requires `System` permissions.

## `deployVersion` (method)

Deploys new version of already deployed smart contract.

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (potentially changes state).

#### Behavior
* Takes service name, processor type and code as arguments.
* Has ellipsis syntax, meaning that additional code can be passed as an argument: `deployService(serviceName, processorType, code1, code2, code3, code4)`
* Make sure the **service is already deployed** by calling `getInfo`.
* Write processor type to state under key `<name>.Processor.Version.<version>` by calling `State.WriteUint32ByKey`.
* Write first part of the code to state under key `<name>.Code.<version>` by calling `State.WriteBytesByKey`.
* Write other parts of the code to state under key `<name>.Code.<counter>.Version.<version>` by calling `State.WriteBytesByKey`.
* Every time code part is written into state, increment `<name>.CodeParts.<version>` counter (except the first time for backwards compatibility)
* Increments `<name>.Version` value by calling `State.WriteUint32ByKey`.
  * This is an `Internal` method, requires `System` permissions.

&nbsp;
## `getCode` (method) - obsolete for multi-file contracts

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Takes service name as argument.
* Read code from state key `<name>.Code` by calling `State.ReadBytesByKey`.

## `getCodeParts` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Takes service name as argument.
* Takes version as optional argument (if not present defaults to `0`)
* Read code from state key `<name>.CodeParts` by calling `State.ReadBytesByKey`, increments in by `1`.
* Returns error if the contract was not deployed

## `getCodePart` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Takes service name as argument.
* Takes index as argument.
* Takes version as optional argument (if not present defaults to `0`)
* Read code from state key `<name>.Code.<counter>` by calling `State.ReadBytesByKey`.
* Returns error if the contract was not deployed

## `getVersion` (method)

Returns current version of the contract (`0` for a contract deployed with `deployService`).

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Takes service name as argument.
* Read code from state key `<name>.Code.<counter>` by calling `State.ReadUint32ByKey`.
* Returns `0` if the contract was not deployed

## `allowVersioning` (method)

Owner only.

Allows new versions of the contract to be uploaded.

## `disallowVersioning` (method)

Owner only.

Disallows new versions of the contract to be uploaded.

## `tranferOwnership` (method)

Owner only.

Transfers ownership to the new owner. Unless ownership is accepted, the old owner still retains full ownership of the contract.

## `acceptOwnership` (method)

New owner can accept or reject the ownership of the contract.

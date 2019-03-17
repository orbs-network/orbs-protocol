# Orbs Vote Processing Contracts
> Orbs voting is performed on the Ethereum platform and processed on Orbs as described in the [voting flow](../../flows/voting.md)).

&nbsp;
## OrbsVoting
Ownership: none

### Global state
#### Delegation and voting
* agent[`delegator`] - the account the delegator delegated to.
* voted_validators[`guardian`] - voted in the election

#### Election data
* `election_block_height` - The last Ethereum block height for delegation and voting, the reference block_height for the stake read. 

#### Reward data
* reward[`address`] - Cumulative rewards awarded to an address (as a Guardian, Delegator or Validator)

#### Last election data
* received_votes[`validator`] - the total stake that was voted to the validator.
* `total_voting_stake` - the total stake that participated in the election.
* `total_in_circulation` - the total stake in circulation.

#### Elections Parameters
* `VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS` - The number of Ethereum blocks after the `election_block_height` during which mirroring is still allowed for the election.
  * Default: 480 (~ 2 hours)
* `ELECTION_CYCLE_IN_BLOCKS` - The number of Ethereum blocks between elections, determines the next  `election_block_height`.
  * Default: 17280 (~ 3 days)
* `VOTING_VALIDITY_TIME` - The number of Ethereum blocks during which a Guardian's voting is valid
  * Default: 40320 (~ 7 days)
* `TRANSITION_PERIOD_LENGTH_IN_BLOCKS` - The number of **Orbs** blocks (of the virtual chains) until an election decision is applied.
  * Default: 0 - the next new committee is applied in the next block.
* `MAXIMUM_ELECTED_VALIDATORS` - the maximum Validators that can initially be elected
  * Default: 22
* `DISAPPROVAL_THRESHOLD_PERCENT` - the percentage of the total participating stake required to disapprove a Validator
  * Default: 70%
* `FIRST_ELECTION_BLOCK` - the Ethereum block of the first election. 
  * 7467969 (approximately Apr 1 noon UTC)
* `MINIMUM_VALIDATORS` - the minimal number of  

#### Rewards Parameters
* `PARTICIPATION_MAX_ANNUAL_REWARD` - The maximum annual reward awarded to stakeholders (Delegators or Guardians) for participation. (Delegators or Guardians)
  * Default: 60M (ORBS)
* `PARTICIPATION_MAX_STAKE_REWARD_PERCENT` - The maximum award in each election as a percent of the participant stake. (taken into consideration when the total participating stake is low: less than `PARTICIPATION_MAX_REWARD`/`PARTICIPATION_MAX_STAKE_REWARD_PERCENT`).
  * Default: 8%
* `VALIDATOR_INTRODUCTION_PROGRAM_ANNUAL_REWARD` - a fixed reward given to Validators at the initial stage of the network.
  * Default: 1M (ORBS)
* `VALIDATORS_STAKE_REWARD_PERCENT` - the Validator reward as a percent of her stake.
  * Default: 4%
* `GUARDIANS_MAX_ANNUAL_REWARD` - maximum annual award for the Guardians Excellence Program
  * Default: 40M (ORBS)
* `GUARDIANS_MAX_DPOS_REWARD_PERCENT` - The maximum award in each election as a percent of the Guardian delegated stake (including its own). Taken into consideration when the total participating stake in the Excellence Program is low: < `GUARDIANS_MAX_ANNUAL_REWARD`/`GUARDIANS_MAX_DPOS_REWARD_PERCENT`).
  * Default: 10%
* `EXCELLENCE_PROGRAM_NUMBER_OF_GUARDIANS` - number of Guardians with the top delegated stake that participate in the Guardians Excellence Program.
  * Default: 10

#### General
* `ETHEREUM_AVG_BLOCK_TIME_SEC` = 15
* `SEC_IN_A_YEAR` = 31536000
* `ELECTION_PARTICIPATION_MAX_REWARD` = (`PARTICIPATION_MAX_ANNUAL_REWARD` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 493150
* `ELECTION_GUARDIANS_MAX_REWARD` = (`GUARDIANS_MAX_ANNUAL_REWARD` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 328767
* `ELECTION_VALIDATOR_INTRODUCTION_MAX_REWARD` = `VALIDATOR_INTRODUCTION_PROGRAM_ANNUAL_REWARD` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 8220


### mirrorDelegationData(delegator, to, delegation_block_height, delegation_tx_index, updated_by)
> Access: internal
> Process an Ethereum delegation transaction and logs it on Orbs.

#### Check the current account delegation:
* if the current delegator_last_update[`delegator`].`block_height`, `tx_index`} is larger than the {`delegation_block_height`,`delegation_tx_index`} fail the transaction.
* If updated_type[`delegator`] = `VOTING_CONTRACT` and `updated_by` = `TRANSFER` fail the transaction.

#### Update delegation map:
* delegator_last_update[`delegator`] = {`delegation_block_height`, `delegation_tx_index`}
* updated_type[`delegator`] = `updated_by`.
* if (`to` == 0) or (`to` == `delegator`):
  * Set agent[`delegator`] = address(0).
* Else
  * Set agent[`delegator`] = `to`.
  

### mirrorDelegation(bytes txid)
> Access: external
> Mirrors an Ethereum delegation by vote contract transaction by calling `mirrorDelegationData`.
> mirrorDelegation may only be called during the vote mirroring period.

#### Check the that vote processing did not start yet
* If processing has started return an error indicating: Resubmit in the next elections.

#### Read and check the Delegate event
* Read the transaction's `Delegate(delegator, to, delegate_counter)` event emitted by the `OrbsVoting` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.

#### Process the delegation
* Process the delegation by calling `mirrorDelegationData(delegator, to, delegation_block_height, delegation_block_index, VOTING_CONTRACT)`

### mirrorDelegationByTransfer(bytes txid)
> Access: external
> Mirrors an Ethereum delegation by transfer transaction by calling `mirrorDelegationData`.
> mirrorDelegation may only be called during the vote mirroring period.

#### Check the that vote processing did not start yet
* If processing has started return an error indicating: Resubmit in the next elections.

#### Read and check the transfer event
* Read the transaction's `transfer(from, to, tokens` event emitted by the `ERC20` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check `tokens` = 7. 
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.

#### Process the delegation
* Process the delegation by calling `mirrorDelegationData(delegator, to, delegation_block_height, delegation_block_index, TRANSFER)`


### mirrorVote(bytes txid)
> Access: external
> Mirrors a vote transaction that was sent on Ethereum
> mirrorVote may only be called during the vote mirroring period.

#### Check the that vote mirroring period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election.

#### Read and check the vote event
* Read the transaction's `vote(guardian, validators_list, vote_counter)` event emitted by the `OrbsVoting` contract.
  * If no was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `vote_block_height` = `ethereum_block_height`.
  * `vote_tx_index` = `ethereum_tx_index`.

#### Check the current `guardian`'s vote:
* if the current guardian_last_update[`delegator`].{`block_height`, `tx_index`} is larger than the {`vote_block_height`,`vote_tx_index`} fail the transaction.

#### Check that the guardian is in the due-diligence list 
* Call the Ethereum guardians contract: Ethereum.Guardians.isGuardian(`guardian`).

#### Update vote map:
* Update `guardian`'s votes map:
  * guardian_last_update[`guardian`] = {`vote_block_height`, `vote_tx_index`}
  * Set voted_validators[`guardian`] = `validators_list`.


### ProcessVoting() : bool
> Access: external
> Process the election results, updates OrbsValidatorsConfig and the next election data.
#### State
* `current_guardian`
* `current_delegator`
* `processing_state`

#### Check that the vote mirroring period has ended
* `ethereum_block_height` > `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`
  * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`
  * If not return an error indicating: mirroring period has not ended.

#### If processing did not start initiate guardians processing
* Check `election_in_process` != `election_ethereum_block_height`
  * `election_in_process` = `election_ethereum_block_height` 
  * `process_guardians` = TRUE

#### Delegators processing
* If `process_delegators` 
  * stake[`current_delegator`] = Ethereum.balanceOf(`current_delegator`).
  * delegators[agent[`current_delegator`]].add(`current_delegator`). (generate an reverse table)
  * If `current_delegator` is last:
    * Clear `process_delegators`.
    * Set `calculate_votes`
  * Else `current_delegator` = next delegator in list. 

#### Guardians processing
* If `process_guardians`:
  * Check guardian_last_update[`current_guardian`] has not expired (< `election_block_height` - `parameter.VOTING_VALIDITY_TIME`).
  * stake[`current_guardian`] = Ethereum.ERC20.balanceOf(`current_guardian`).
  * If `current_guardian` is last:
    * clear `process_guardians`.
    * Set `process_delegators`.
    * Else `current_guardian` = next guardian in list. 

#### Calculations
* If `calculate_votes`
  * Call `CalculateVotes()`.
  * Call `CalculateElectedValidators()` and update `OrbsValidatorsConfig.UpdateElectionResult`.
  * Set `election_ethereum_block_height` = `election_ethereum_block_height` + `parameter.ELECTION_CYCLE_IN_BLOCKS`.
  * Clear `calculate_votes`.

#### Return processing status (completed) 
* If (`process_delegators` OR `process_guardians` OR `calculate_votes`)
  * Return FALSE
* Else 
  * Return TRUE

### CalculateVotes()
> Calculates the active validators based on the votes:

#### Remove the Guardians from the delegation list (implicitly delegate to themselves)
* Get the `guardians_list` from Ethereum 
  * Call `Ethereum.OrbsGuardians.getGuardians()`
* For every `guardian` in guardians list:
  * `reference_agent` = agent[`guardian`]
  * Remove `guardian` from delegators[`reference_agent`]

#### Calculate the hierarchical voting_stake
* Recursively set the voting_stake of each `guardian` or `delegator`, start with the `guardians_list`: 
  * Add participant to participants_list
  * `total_voting_stake` += stake[participant]
  * If the participant has no `delegator`s that delegated to it then voting_stake = stake[participant]
  * Else voting_stake = sum(delegating `delegator`s voting_stake) + stake[participant]

#### Calculate the elected validators
* Get the validators_list
  * Call `Ethereum.OrbsValidators.getValidators()`
  * Alternatively use `Ethereum.OrbsValidators.isValidator(Validator)` based on the voted addresses.

* For (every `guardian` in guardians list):
  * For every validator in voted_validators[`guardian`]:
    * received_votes[`validator`] += voting_stake[`guardian`].

* Calculate `disapproval_threshold` = `total_voting_stake` * `DISAPPROVAL_THRESHOLD_PERCENT`

* elected_validators = validators in `validators_list` with: 
  * received_votes[`validator`] < `disapproval_threshold`.

#### Minimal number of Validators protection
* If number of elected_validators < `MINIMUM_VALIDATORS`
  * elected_validators = validators in `validators_list`

#### ------------------------------------------------------------------------------
#### TOP22 Calculation - Not Implemented

* Generate `currently_elected_approved_list`
  * `currently_elected_approved_list` = validators in `validators_list` with:
    * received_votes[`validator`] < `disapproval_threshold`.
    * in `currently_elected_validators`.

* Find the top `top_candidate`
  * `top_candidate` is the validator in `validators_list` with the most stake[`validator`] that meets:
    * received_votes[`validator`] < `disapproval_threshold`.
    * not in `currently_elected_validators`.

* If `currently_elected_approved_list` < `MAXIMUM_ELECTED_VALIDATORS`
  * elected_validators.add(`currently_elected_approved_list`)
  * elected_validators.add(`top_candidate`)
* Else
  * Sort the `currently_elected_approved_list` by their stake
  * elected_validators = `MAXIMUM_ELECTED_VALIDATORS` - 1 validators
  * `bottom_validator` = the bottom validator in `currently_elected_approved_list`.
  * If stake[`bottom_validator`] > stake[`top_candidate`]
    * elected_validators.add(`bottom_validator`)
  * Else
    * elected_validators.add(`top_candidate`)

#### ------------------------------------------------------------------------------

#### Calculate the participation reward
* Calculate the participation reward for the election
  * `election_participation_reward` = min(`ELECTION_PARTICIPATION_MAX_REWARD`, `total_voting_stake` * `PARTICIPATION_MAX_REWARD_PERCENT`)
* Calculate the participation reward
  * For every delegator or guardian:
    * participation_reward[participant] += stake[participant] / `total_voting_stake` * `election_participation_reward`

#### Calculate the Guardians Excellence Program reward
* Calculate the participants in the Guardians Excellence Program for the election
  * `excellence_program_participants` = the top `EXCELLENCE_PROGRAM_NUMBER_OF_GUARDIANS` with the most `voting_stake`.
  * `total_excellence_program_stake` = total `voting_stake` of the excellence_program_participants.  
* Calculate the Guardians Excellence Program reward for the election
  * `election_guardians_reward` = min(`ELECTION_GUARDIANS_MAX_REWARD`, `total_excellence_program_stake` * `GUARDIANS_MAX_DPOS_REWARD_PERCENT`).
* Calculate the Guardians reward
  * For every `guardian` in excellence_program_participants:
    * guardian_excellence_reward[`guardian`] += voting_stake[`guardian`] / `total_excellence_program_stake` * `election_guardians_reward`

#### Calculate the Validators reward
* For every `elected_validator` in `elected_validators`
  * validator_reward[`elected_validator`] += stake[`elected_validator`] * `VALIDATORS_STAKE_REWARD_PERCENT` + `VALIDATOR_INTRODUCTION_PROGRAM_REWARD`


&nbsp;
### Elected Validators Database (part of the voting contract)

#### UpdateElectionResult(election_ethereum_block_height, elected_validators)
> Stores the list of elected validators per block height
* `election_index` += 1
* Validators[`election_index`] = `elected_validators`
* ElectionEthereumBlockHeight[`election_index`] = `election_block_height`
* ValidatorsApplyBlockHeight[`election_index`] = current **Orbs** block_height` + `parameter.TRANSITION_PERIOD_LENGTH_IN_BLOCKS`


#### GetElectedValidatorsByHeight(orbs_block_height) : elected_validators
> Return the list of elected validators per requested block_height
* Find the `reference_election_index` = the largest election_index, such that ValidatorsApplyBlockHeight[election_index] is smaller than the provided block_height.
* Return Validators[`reference_election_index`]

#### GetNumberOfElections : uint
> Returns `election_index`, used for migrations.

#### GetElectedValidatorsByIndex(election_index) : (election_ethereum_block_height, apply_block_height, validators_list)
> Returns the list of elected validators and relevant election and apply block_height by election_index, used for migrations.


&nbsp;
### OrbsRewards
> Holds the guardians, delegators nad validators rewards


&nbsp;
## Getters Interface

#### General
* getNextElectionBlockNumber() : Ethereum_block
* getEffectiveElectionBlockNumber() : Ethereum_block
* getElectionPeriod() : Number Of Ethereum_blocks
* getNumberOfElections()
* getElectedValidatorsOrbsAddress() : list of Validators
* getElectedValidatorsEthereumAddress() : list of validators 
  
#### Extended Last Election Results
* GetGuardianVotingWeight(Guardian) : uint
* GetGuardianStake(Guardian) : uint (Integer ORBS)
* GetValidatorVotes(ValidatorEthereumAddress) : uint
* GetValidatorStake(ValidatorEthereumAddress) : uint (Integer ORBS)
* GetTotalStake() : uint
* GetExcellenceProgramGuardians() : list of guardians 

#### Historical data
* getElectedValidatorsOrbsAddressByBlockHeight(orbs_block_height) : elected_validators
* getElectedValidatorsEthereumAddressByBlockHeight() : elected_validators
* getElectedValidatorsOrbsAddressByIndex(index) : elected_validators (Ethereum addresses).
* getElectedValidatorsOrbsEthereumByIndex(index) : elected_validators (Ethereum addresses).
* getElectedValidatorsBlockNumberByIndex(index) : Election event Ethereum block number
* getElectedValidatorsBlockHeightByIndex(index) : Apply results orbs_block_height

#### Rewards
* GetCumulativeParticipationReward(delegator) : uint (Integer ORBS)
* GetCumulativeValidatorReward(validatorEthereumAddress) : uint (Integer ORBS)
* GetCumulativeGuardiansExcellenceReward(guardian) : uint (Integer ORBS)
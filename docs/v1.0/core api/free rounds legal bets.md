---
title: Legal bets
excerpt: Legal bets term, setup and values.
category: 65bc99b3a8f20d00395a35e7
order: -1
hidden: false
slug: core-free-rounds-legal-bet
---

Legal bets are related to operator's free rounds settings. It contains from the following table, which operator fills up and provided before the integration:

| Parameter      | Value               | Description           |
| -------------- | ------------------- | --------------------- |
| provider_id    | [string] or null    | operator's ID         |
| game           | [string]            | game's ID             |
| currency       | [string]            | currency's ID         |
| bet_value      | [array or string]   | bet's values          |

The `bet_value` parameter can be an array or a string. It contains the following structure:
<max_bet>:<max_target_multiplier>
where:
- `max_bet` - maximum bet value for this currency (full value, not in cents), e.g. `1.00`
- `max_target_multiplier` - maximum target multiplier for the game, e.g. `100.00`

When operator controls free rounds on its side, he just provides <bet_id> which is a number from 1 to N (where N is the number of legal bets) and the game will use the bet value from the list of legal bets.

When operator controls free rounds on the game side, he provides the bet value and the game will use it.
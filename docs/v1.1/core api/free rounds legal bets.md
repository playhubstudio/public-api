---
title: Legal bets
excerpt: Legal bets term, setup and values.
category: 671f27ffff7f7a00128379aa
order: -1
hidden: false
slug: core-free-rounds-legal-bet
---

Legal bets are related to the operator's free rounds settings. The operator fills out the following table and provides it before the integration:

| Parameter   | Value             | Description   |
| ----------- | ----------------- | ------------- |
| provider_id | [string] or null  | Operator's ID |
| game        | [string]          | Game's ID     |
| currency    | [string]          | Currency's ID |
| bet_value   | [array or string] | Bet's values  |

### bet_value Parameter

The `bet_value` parameter can be an array or a string and follows this structure:

```
<max_bet>:<max_target_multiplier>
```

Where:

- **max_bet**: Maximum bet value for this currency (full value, not in cents), e.g., 1.00
- **max_target_multiplier**: Maximum target multiplier for the game, e.g., 100.00

### Free Rounds Control

- **Operator Controls Free Rounds on Their Side**:
  - The operator provides a `<bet_id>`, which is a number from 1 to N (where N is the number of legal bets).
  - The game uses the bet value from the list of legal bets.

- **Operator Controls Free Rounds on the Game Side**:
  - The operator provides the bet value, and the game will use it.

By following these guidelines, operators can effectively manage free rounds and legal bets within the game.
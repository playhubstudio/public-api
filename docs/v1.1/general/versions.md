---
title: Versions
excerpt: Versions.
category: 671b459dfe48cf0030150ea1
order: -1
hidden: false
slug: general-versions
---

# 1.1.3

- Added `/games` request to Core API to provide list of supported games and related settings. Now you can retrieve list of supported games and their settings from Core API. This request is useful for integrations where you need to get the list of games and their settings dynamically. For example, you can use this request to get the list of games and their settings for the game lobby, or to get the list of games and their settings for the game configuration page in the admin panel.

# 1.1.2

- Set 'user' field as optional in `game-create-new` request. This allows to create a game without a user (for example, for demo game). In this case we will generate suitable user's id on our side.

# 1.1.1

- Added 'round_id' for Bet and Win requests in Playhub Adapter. This allows to identify/link free rounds config used for the bet/win free rounds requests in external integrations.
- Added 'orig_tx_id' for Refund requests in Playhub Adapter. This allows to identify/link the original transaction for the refund requests in external integrations.

For now 'round_id' and tx_id' are the same for Bet and Win requests. In the future, they may be different.  
'orig_tx_id' is the same as 'tx_id' for Refund requests. In future, they may be different.
We added these parameters to make it easier to track the requests in external systems and also to change 'tx_id' logic in future version without breaking the existing integrations.
'tx_id' for now represents internal 'round_id' in the system. In the future, it may be changed to represent the external transaction id (unique per integration).

# 1.1.0

- Added 'frc_id' field to Bet and Win requests (for type='free') in Playhub Adapter. This allows to identify/link free rounds config used for the bet/win free rounds requests in external integrations.

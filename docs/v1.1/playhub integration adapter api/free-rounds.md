---
title: Free rounds
excerpt: Playhub Integration Adapter Free rounds.
category: 671f28042b66d5006370115e
order: -1
hidden: false
slug: playhub-integration-adapter-free-rounds
---

# Bonus balance

When player plays Free Rounds its main balance is not affected. All wins are going to the bonus balance.
Bonus balance is a separate balance that can be used only for Free Rounds. It can't be withdrawn.
Once Free Rounds are done we transfer all wins from the bonus balance to the main balance.

# "Free" request type

In every request from our PHIA (Playhub Integration Adapter) we send "type" parameter. If it's "free" it means that player is playing Free Rounds and amount we send it's virtual bonus balance. It's not real money. So you don't need to process it as a real money transaction but you need to return real money balance to the PHIA. Usually clients just skip this request and return real money balance to the PHIA.

# Free rounds finalization

When free rounds are over on our side we send a request to finalize free rounds ( notice that type is not "free" anymore). In this request we send all wins from the bonus balance. You need to transfer this amount to the main balance and return the new balance to the PHIA. In this case we add `frc_id` parameter to the request to identify the free rounds session. This means real money transaction and you need to process it as a real money transaction.


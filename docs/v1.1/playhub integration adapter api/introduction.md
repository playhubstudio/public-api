---
title: Introduction
excerpt: Playhub Integration Adapter Introduction.
category: 671f28042b66d5006370115e
order: -1
hidden: false
slug: playhub-integration-adapter-intro
---

# Initial setup

To start integration with Playhub, you need to provide us with your endpoints information. It might be single endpoint for all requests or different endpoints for different requests. We support HTTPS protocol only.

# Round, transaction and money processing.
Every request through the Playhub Integration Adapter API contains transaction ID (`tx_id`) and round ID (`round_id`).
All transactions from the same round (`bet`, `win`, `refund`, etc.) have the same round_id but might differ in `tx_id`. `tx_id` is mostly used as idempotency key for requests, means
same `bet` request with same `tx_id` should be processed only once.
Also please note that `round_id` might be non unique for different games (e.g. `crash` and `limbo` games might have the same `round_id`).
In additional you can use `game_session_id` along with `round_id` to make it a unique combination on your side, since `game_session_id` is unique per game (for the same user and currency).


# Errors processing. Refund policy.
All error codes except 200 (OK) are considered as errors. Every bet request with non 200 (OK) response will be refunded.

# Idempotency Requirement

All requests from our platform to yours should be processed in an idempotent way. This means that if we send the same request multiple times, the result should be the same as if we sent the request only once. This is crucial to avoid any issues with your system in case of network problems or other issues.

# Free rounds

When you receive a request with `isFree` parameter in additional we send `frc_id` field which is related to free rounds configuration.

# IP Addresses

We use the following IPs to access your endpoints:

- **Staging Environment**: 34.76.77.205
- **Production Environment**: 34.140.75.241

By ensuring idempotency and allowing access from these IP addresses, you can maintain the reliability and stability of the integration.
---
title: Introduction
excerpt: Playhub Integration Adapter Introduction.
category: 671f28042b66d5006370115e
order: -1
hidden: false
slug: playhub-integration-adapter-intro
---

# General

Playhub integration adapter (PHIA) is a service that allows you to integrate your game with Playhub platform.
It's located in our cluster and sends requests to your endpoints on every game action (mostly related to money change).
You integrate your backend with our PHIA and we take care of the rest.
Your integration should be able to handle requests from our PHIA and respond to them in a timely manner (as quicker as possile).
We don't use any retry mechanisms to make requests as quick as possible, so you should be able to handle requests in a timely manner and respond to them as soon as possible.

# Initial setup and security

To start integration with Playhub, you need to provide us with your endpoints information.
It might be single endpoint for all requests or different endpoints for different requests. We support HTTPS protocol only.
We use request signature with default key, you can either share your signature key with us or whitelist our IP addresses and avoid signature check.

# Round, transaction and money processing.

Every request through the Playhub Integration Adapter API contains transaction ID (`tx_id`) and round ID (`round_id`).
All transactions from the same round (`bet`, `win`, `refund`, etc.) have the same round_id but might differ in `tx_id`. `tx_id` is mostly used as idempotency key for requests, means
same `bet` request with same `tx_id` should be processed only once.
Also please note that `round_id` might be non unique for different games (e.g. `crash` and `limbo` games might have the same `round_id`).
In additional you can use `game_session_id` along with `round_id` to make it a unique combination on your side, since `game_session_id` is unique per game (for the same user and currency).

# Errors processing. Refund policy.

All error codes except 200 (OK) are considered as errors. We decline game's action in case of any error.
No refund will be sent in this case.
Timeout over than 6 seconds will be considered as an error as well and refund will be issued afterwards.
In error your can return any useful information to our side, later this could be useful to debug some issues.
We don't show any errors from integrations to the end user, so you can return any error message you want ( for debugging purposes).
Otherwise we have special errors which allow our client to execute some actoins on client's side, see more [Client actions](playhub-integration-adapter-client-actions)

# Idempotency Requirement

All requests from our platform to yours should be processed in an idempotent way.
This means that if we send the same request multiple times, the result should be the same as if we sent the request only once.
This is crucial to avoid any issues with your system in case of network problems or other issues.
We use `tx_id` field to ensure idempotency.

# IP Addresses

We use the following IPs to access your endpoints:

**Staging Environment**

- 34.22.222.37
  **Production Environment**
- 34.140.75.241 (deprecated and will be removed soon)
- 104.155.124.243

By ensuring idempotency and allowing access from these IP addresses, you can maintain the reliability and stability of the integration.

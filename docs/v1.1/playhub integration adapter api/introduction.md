---
title: Introduction
excerpt: Playhub Integration Adapter Introduction.
category: 671f28042b66d5006370115e
order: -1
hidden: false
slug: playhub-integration-adapter-intro
---

## Round, transaction and money processing.
Every request through the Playhub Integration Adapter API contains transaction ID (tx_id) which is a unique for every round.
So all money processing requests/transactions within a single round is processed with the same tx_id.
It means you can use it to identify all transactions related to a single round.


## Errors processing. Refund policy.
All error codes except 200 (OK) are considered as errors. Every bet request with non 200 (OK) response will be refunded.

## Idempotency Requirement

All requests from our platform to yours should be processed in an idempotent way. This means that if we send the same request multiple times, the result should be the same as if we sent the request only once. This is crucial to avoid any issues with your system in case of network problems or other issues.

## Free rounds

When you receive a request with `isFree` parameter in additional we send `frc_id` field which is related to free rounds configuration.

## IP Addresses

We use the following IPs to access your endpoints:

- **Staging Environment**: 34.76.77.205
- **Production Environment**: 34.140.75.241

By ensuring idempotency and allowing access from these IP addresses, you can maintain the reliability and stability of the integration.
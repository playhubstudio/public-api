---
title: Introduction
excerpt: Playhub Integration Adapter Introduction.
category: 671f23bd5e2e790010cf4879
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

## IP Addresses

We use the following IPs to access your endpoints:

**Staging Environment** 
- 34.22.222.37
**Production Environment** 
- 34.140.75.241 (deprecated and will be removed soon)
- 104.155.124.243

By ensuring idempotency and allowing access from these IP addresses, you can maintain the reliability and stability of the integration.
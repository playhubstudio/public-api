---
title: Introduction
excerpt: Playhub Integration Adapter Introduction.
category: 65bc99c69969910010b1426c
order: -1
hidden: false
slug: playhub-integration-adapter-intro
---

## Errors processing. Refund policy.

All error codes except 200 (OK) are considered as errors. Every bet request with non 200 (OK) response will be refunded.

## Idempotency Requirement

All requests from our platform to yours should be processed in an idempotent way. This means that if we send the same request multiple times, the result should be the same as if we sent the request only once. This is crucial to avoid any issues with your system in case of network problems or other issues.

## IP Addresses

We use the following IPs to access your endpoints:

- **Staging Environment**: 34.76.77.205
- **Production Environment**: 34.140.75.241

By ensuring idempotency and allowing access from these IP addresses, you can maintain the reliability and stability of the integration.
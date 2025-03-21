---
title: Client actions
excerpt: Playhub Integration Adapter Client Actions.
category: 671f28042b66d5006370115e
order: 2
hidden: false
slug: playhub-integration-adapter-client-actions
---

# General

When you return any error from Playhub Integration Adapter request, there are some special cases you can use to manage how client will behave.  
See the list below.

## Page reload behavior

If you return error message which cointains any of

- `session not found`
- `wrong session ID`
- `cannot find external session`
- `unauthorized`

Client will stop the game immediately and show modal dialogue "Session expired, please reload the page"

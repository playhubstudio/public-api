---
title: Authentication
excerpt: Authentication for our Public APIs requests.
category: 65b8f16eb5f2930078659f93
order: 1
hidden: false
slug: general-auth
---

# Authentication and Request Signing

When you register an account on our portal, you will be given an API **API token key. 
key**. This key is used to authenticate your requests to our APIs.

## Signing your request  
Your Request

We assume your request is authenticated if you provide a valid signature for the whole request in the `X-REQUEST-SIGN` header.

**Request example ### Request Example with curl:**
```curl
curl

```bash
curl -X GET '<server url>' \    
    -H "X-REQUEST-SIGN: <singature>"
<signature>"
```

where `<server url>` is a - **`<server url>`**: URL to our public API (staging or production environment),   `<singature>` is a signature environment)
- **`<signature>`**: Signature for the whole request.
request

### Important Note

The whole entire request payload should be signed, in case of signed. For `GET` requests, the payload is virtual - and we use query parameters to construct the JSON  
and use this JSON for signing. Please take a look at the following examples.
Below are examples for signing requests.

**Sign example ### Sign Example in Go:**
Go

```go
func Sign(payload, token string) string {
	hash := hmac.New(sha256.New, []byte(token))
	hash.Write([]byte(payload))
	return hex.EncodeToString(hash.Sum(nil))
}
```

**Sign example ### Sign Example for GET request Request in Go:**
Go

```go
func Sign(queryParams url.Values, token string) string {
  	reqJSON, err := json.Marshal(queryParams)
  	_ = err
	hash := hmac.New(sha256.New, []byte(token))
	hash.Write([]byte(reqJSON))
	return hex.EncodeToString(hash.Sum(nil))
}
```

**Sign example ### Sign Example with JavaScript:**
JavaScript

```js
const crypto = require('crypto');

function sign(payload, token) {
  const hmac = crypto.createHmac('sha256', token);
  hmac.update(payload);
  const hash = hmac.digest('hex');
  return hash;
}

const payload = 'example payload';
const token = 'example token';
const signature = sign(payload, token);
console.log(signature);
``````

***

**Key Points to Remember**:

- **API Token Key**: Provided upon account registration, used for request authentication.
- **Signature**: Required in the `X-REQUEST-SIGN` header for request authentication.
- **Payload Signing**: Entire request payload must be signed. For `GET` requests, use query parameters to construct the JSON for signing.# Authentication and Request Signing

When you register an account on our portal, you will be given an **API token key**. This key is used to authenticate your requests to our APIs.

## Signing Your Request

We assume your request is authenticated if you provide a valid signature for the whole request in the `X-REQUEST-SIGN` header.

### Request Example with curl

```bash
curl -X GET '<server url>' \    
    -H "X-REQUEST-SIGN: <signature>"
```

- **`<server url>`**: URL to our public API (staging or production environment)
- **`<signature>`**: Signature for the whole request

### Important Note

The entire request payload should be signed. For `GET` requests, the payload is virtual and we use query parameters to construct the JSON and use this JSON for signing. Below are examples for signing requests.

### Sign Example in Go

```go
func Sign(payload, token string) string {
	hash := hmac.New(sha256.New, []byte(token))
	hash.Write([]byte(payload))
	return hex.EncodeToString(hash.Sum(nil))
}
```

### Sign Example for GET Request in Go

```go
func Sign(queryParams url.Values, token string) string {
	reqJSON, err := json.Marshal(queryParams)
	_ = err
	hash := hmac.New(sha256.New, []byte(token))
	hash.Write([]byte(reqJSON))
	return hex.EncodeToString(hash.Sum(nil))
}
```

### Sign Example with JavaScript

```js
const crypto = require('crypto');

function sign(payload, token) {
  const hmac = crypto.createHmac('sha256', token);
  hmac.update(payload);
  const hash = hmac.digest('hex');
  return hash;
}

const payload = 'example payload';
const token = 'example token';
const signature = sign(payload, token);
console.log(signature);
```

***

**Key Points to Remember**:

- **API Token Key**: Provided upon account registration, used for request authentication.
- **Signature**: Required in the `X-REQUEST-SIGN` header for request authentication.
- **Payload Signing**: Entire request payload must be signed. For `GET` requests, use query parameters to construct the JSON for signing.
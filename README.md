# CrisisCoreFusion SDK

Simulate materia fusion from Crisis Core: Final Fantasy VII and look up possible fusion outcomes

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Crisis Core Fusion API

The Crisis Core Fusion API is a small Go service that models the materia fusion mechanics from the video game *Crisis Core: Final Fantasy VII*. It is maintained as an open-source side project by [RayMathew](https://github.com/RayMathew/crisis-core-materia-fusion-api) and runs on Google Cloud Run.

What you get from the API:

- `GET /materia` — list every materia known to the service.
- `POST /fusion` — submit two materia (with their mastery status) and receive the resulting fusion outcome computed from the game's rules.
- `GET /status` — lightweight health check for uptime monitoring.

No authentication, API key, or documented rate limit is required to call the public endpoint.

## Try it

**TypeScript**
```bash
npm install crisis-core-fusion
```

**Python**
```bash
pip install crisis-core-fusion-sdk
```

**PHP**
```bash
composer require voxgig/crisis-core-fusion-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/crisis-core-fusion-sdk/go
```

**Ruby**
```bash
gem install crisis-core-fusion-sdk
```

**Lua**
```bash
luarocks install crisis-core-fusion-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { CrisisCoreFusionSDK } from 'crisis-core-fusion'

const client = new CrisisCoreFusionSDK({})

```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o crisis-core-fusion-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "crisis-core-fusion": {
      "command": "/abs/path/to/crisis-core-fusion-mcp"
    }
  }
}
```

## Entities

The API exposes 3 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Fusion** | Fusion operations that combine two input materia and return the resulting materia, exposed at `POST /fusion`. | `/fusion` |
| **Materia** | The catalogue of materia items available as fusion inputs, listed via `GET /materia`. | `/materia` |
| **System** | Service-level utilities such as the `GET /status` health check. | `/health` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from crisiscorefusion_sdk import CrisisCoreFusionSDK

client = CrisisCoreFusionSDK({})

```

### PHP

```php
<?php
require_once 'crisiscorefusion_sdk.php';

$client = new CrisisCoreFusionSDK([]);

```

### Golang

```go
import sdk "github.com/voxgig-sdk/crisis-core-fusion-sdk/go"

client := sdk.NewCrisisCoreFusionSDK(map[string]any{})

```

### Ruby

```ruby
require_relative "CrisisCoreFusion_sdk"

client = CrisisCoreFusionSDK.new({})

```

### Lua

```lua
local sdk = require("crisis-core-fusion_sdk")

local client = sdk.new({})

```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = CrisisCoreFusionSDK.test()
const result = await client.Fusion().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = CrisisCoreFusionSDK.test(None, None)
result, err = client.Fusion(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = CrisisCoreFusionSDK::test(null, null);
[$result, $err] = $client->Fusion(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Fusion(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = CrisisCoreFusionSDK.test(nil, nil)
result, err = client.Fusion(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Fusion(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Crisis Core Fusion API

- Upstream: [https://github.com/RayMathew/crisis-core-materia-fusion-api](https://github.com/RayMathew/crisis-core-materia-fusion-api)
- API docs: [https://freepublicapis.com/crisis-core-fusion-api](https://freepublicapis.com/crisis-core-fusion-api)

- Distributed under the GNU General Public License v3.0 (GPL-3.0).
- Source-available; redistributions and derivative works must remain under the same licence.
- Attribution to the upstream project is expected.
- Final Fantasy VII and Crisis Core are trademarks of Square Enix; this is an unofficial fan project and is not endorsed by the rights holders.

---

Generated from the Crisis Core Fusion API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

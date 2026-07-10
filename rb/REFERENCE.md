# CrisisCoreFusion Ruby SDK Reference

Complete API reference for the CrisisCoreFusion Ruby SDK.


## CrisisCoreFusionSDK

### Constructor

```ruby
require_relative 'CrisisCoreFusion_sdk'

client = CrisisCoreFusionSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CrisisCoreFusionSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = CrisisCoreFusionSDK.test
```


### Instance Methods

#### `Fusion(data = nil)`

Create a new `Fusion` entity instance. Pass `nil` for no initial data.

#### `Materia(data = nil)`

Create a new `Materia` entity instance. Pass `nil` for no initial data.

#### `System(data = nil)`

Create a new `System` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## FusionEntity

```ruby
fusion = client.Fusion
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `materia1` | `String` | Yes |  |
| `materia1_mastered` | `Boolean` | Yes |  |
| `materia2` | `String` | Yes |  |
| `materia2_mastered` | `Boolean` | Yes |  |
| `result` | `Hash` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Fusion.create({
  "materia1" => "example_materia1", # String
  "materia1_mastered" => true, # Boolean
  "materia2" => "example_materia2", # String
  "materia2_mastered" => true, # Boolean
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FusionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MateriaEntity

```ruby
materia = client.Materia
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `String` | No |  |
| `id` | `Integer` | Yes |  |
| `max_level` | `Integer` | No |  |
| `name` | `String` | Yes |  |
| `rarity` | `String` | No |  |
| `type` | `String` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Materia.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Materia.load({ "id" => "materia_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MateriaEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SystemEntity

```ruby
system = client.System
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `status` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.System.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SystemEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = CrisisCoreFusionSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


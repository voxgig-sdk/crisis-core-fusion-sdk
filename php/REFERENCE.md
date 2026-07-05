# CrisisCoreFusion PHP SDK Reference

Complete API reference for the CrisisCoreFusion PHP SDK.


## CrisisCoreFusionSDK

### Constructor

```php
require_once __DIR__ . '/crisiscorefusion_sdk.php';

$client = new CrisisCoreFusionSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CrisisCoreFusionSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = CrisisCoreFusionSDK::test();
```


### Instance Methods

#### `Fusion($data = null)`

Create a new `FusionEntity` instance. Pass `null` for no initial data.

#### `Materia($data = null)`

Create a new `MateriaEntity` instance. Pass `null` for no initial data.

#### `System($data = null)`

Create a new `SystemEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): CrisisCoreFusionUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## FusionEntity

```php
$fusion = $client->Fusion();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `materia1` | `string` | Yes |  |
| `materia1_mastered` | `bool` | Yes |  |
| `materia2` | `string` | Yes |  |
| `materia2_mastered` | `bool` | Yes |  |
| `result` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Fusion()->create([
  "materia1" => null, // string
  "materia1_mastered" => null, // bool
  "materia2" => null, // string
  "materia2_mastered" => null, // bool
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FusionEntity`

Create a new `FusionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MateriaEntity

```php
$materia = $client->Materia();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | No |  |
| `id` | `int` | Yes |  |
| `max_level` | `int` | No |  |
| `name` | `string` | Yes |  |
| `rarity` | `string` | No |  |
| `type` | `string` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Materia()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Materia()->load(["id" => "materia_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MateriaEntity`

Create a new `MateriaEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SystemEntity

```php
$system = $client->System();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `status` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->System()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SystemEntity`

Create a new `SystemEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new CrisisCoreFusionSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```


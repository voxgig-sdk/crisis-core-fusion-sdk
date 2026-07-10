# CrisisCoreFusion Python SDK Reference

Complete API reference for the CrisisCoreFusion Python SDK.


## CrisisCoreFusionSDK

### Constructor

```python
from crisiscorefusion_sdk import CrisisCoreFusionSDK

client = CrisisCoreFusionSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CrisisCoreFusionSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = CrisisCoreFusionSDK.test()
```


### Instance Methods

#### `Fusion(data=None)`

Create a new `FusionEntity` instance. Pass `None` for no initial data.

#### `Materia(data=None)`

Create a new `MateriaEntity` instance. Pass `None` for no initial data.

#### `System(data=None)`

Create a new `SystemEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## FusionEntity

```python
fusion = client.Fusion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `materia1` | `str` | Yes |  |
| `materia1_mastered` | `bool` | Yes |  |
| `materia2` | `str` | Yes |  |
| `materia2_mastered` | `bool` | Yes |  |
| `result` | `dict` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Fusion().create({
    "materia1": "example_materia1",  # str
    "materia1_mastered": True,  # bool
    "materia2": "example_materia2",  # str
    "materia2_mastered": True,  # bool
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FusionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MateriaEntity

```python
materia = client.Materia()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `str` | No |  |
| `id` | `int` | Yes |  |
| `max_level` | `int` | No |  |
| `name` | `str` | Yes |  |
| `rarity` | `str` | No |  |
| `type` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Materia().list()
for materia in results:
    print(materia)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Materia().load({"id": "materia_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MateriaEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SystemEntity

```python
system = client.System()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `status` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.System().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SystemEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = CrisisCoreFusionSDK({
    "feature": {
        "test": {"active": True},
    },
})
```


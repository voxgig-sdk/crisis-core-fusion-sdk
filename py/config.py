# CrisisCoreFusion SDK configuration


def make_config():
    return {
        "main": {
            "name": "CrisisCoreFusion",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://crisis-core-materia-fusion-api-546461677134.us-central1.run.app",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "fusion": {},
                "materia": {},
                "system": {},
            },
        },
        "entity": {
      "fusion": {
        "fields": [
          {
            "name": "materia1",
            "req": True,
            "type": "`$STRING`",
            "active": True,
            "index$": 0,
          },
          {
            "name": "materia1_mastered",
            "req": True,
            "type": "`$BOOLEAN`",
            "active": True,
            "index$": 1,
          },
          {
            "name": "materia2",
            "req": True,
            "type": "`$STRING`",
            "active": True,
            "index$": 2,
          },
          {
            "name": "materia2_mastered",
            "req": True,
            "type": "`$BOOLEAN`",
            "active": True,
            "index$": 3,
          },
          {
            "name": "result",
            "req": False,
            "type": "`$OBJECT`",
            "active": True,
            "index$": 4,
          },
        ],
        "name": "fusion",
        "op": {
          "create": {
            "name": "create",
            "points": [
              {
                "method": "POST",
                "orig": "/fusion",
                "parts": [
                  "fusion",
                ],
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "args": {},
                "select": {},
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "materia": {
        "fields": [
          {
            "name": "description",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 0,
          },
          {
            "name": "id",
            "req": True,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 1,
          },
          {
            "name": "max_level",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 2,
          },
          {
            "name": "name",
            "req": True,
            "type": "`$STRING`",
            "active": True,
            "index$": 3,
          },
          {
            "name": "rarity",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 4,
          },
          {
            "name": "type",
            "req": True,
            "type": "`$STRING`",
            "active": True,
            "index$": 5,
          },
        ],
        "name": "materia",
        "op": {
          "list": {
            "name": "list",
            "points": [
              {
                "method": "GET",
                "orig": "/materia",
                "parts": [
                  "materia",
                ],
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "args": {},
                "select": {},
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "list",
          },
          "load": {
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "name",
                      "reqd": True,
                      "type": "`$STRING`",
                      "active": True,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/materia/{name}",
                "parts": [
                  "materia",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "name": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "system": {
        "fields": [
          {
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 0,
          },
        ],
        "name": "system",
        "op": {
          "load": {
            "name": "load",
            "points": [
              {
                "method": "GET",
                "orig": "/health",
                "parts": [
                  "health",
                ],
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "args": {},
                "select": {},
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "load",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }

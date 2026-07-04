package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/crisis-core-fusion-sdk/go"
	"github.com/voxgig-sdk/crisis-core-fusion-sdk/go/core"

	vs "github.com/voxgig-sdk/crisis-core-fusion-sdk/go/utility/struct"
)

func TestMateriaEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Materia(nil)
		if ent == nil {
			t.Fatal("expected non-nil MateriaEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := materiaBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "materia." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set CRISISCOREFUSION_TEST_MATERIA_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		materiaRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.materia", setup.data)))
		var materiaRef01Data map[string]any
		if len(materiaRef01DataRaw) > 0 {
			materiaRef01Data = core.ToMapAny(materiaRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = materiaRef01Data

		// LIST
		materiaRef01Ent := client.Materia(nil)
		materiaRef01Match := map[string]any{}

		materiaRef01ListResult, err := materiaRef01Ent.List(materiaRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, materiaRef01ListOk := materiaRef01ListResult.([]any)
		if !materiaRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", materiaRef01ListResult)
		}

		// LOAD
		materiaRef01MatchDt0 := map[string]any{
			"id": materiaRef01Data["id"],
		}
		materiaRef01DataDt0Loaded, err := materiaRef01Ent.Load(materiaRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		materiaRef01DataDt0LoadResult := core.ToMapAny(materiaRef01DataDt0Loaded)
		if materiaRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if materiaRef01DataDt0LoadResult["id"] != materiaRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func materiaBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "materia", "MateriaTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read materia test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse materia test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"materia01", "materia02", "materia03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("CRISISCOREFUSION_TEST_MATERIA_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CRISISCOREFUSION_TEST_MATERIA_ENTID": idmap,
		"CRISISCOREFUSION_TEST_LIVE":      "FALSE",
		"CRISISCOREFUSION_TEST_EXPLAIN":   "FALSE",
	})

	idmapResolved := core.ToMapAny(env["CRISISCOREFUSION_TEST_MATERIA_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["CRISISCOREFUSION_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
			},
			extra,
		})
		client = sdk.NewCrisisCoreFusionSDK(core.ToMapAny(mergedOpts))
	}

	live := env["CRISISCOREFUSION_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["CRISISCOREFUSION_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}

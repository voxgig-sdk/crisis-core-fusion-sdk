
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { CrisisCoreFusionSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('MateriaEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CRISISCOREFUSION_TEST_LIVE=TRUE.
  afterEach(liveDelay('CRISISCOREFUSION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CrisisCoreFusionSDK.test()
    const ent = testsdk.Materia()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CRISIS_CORE_FUSION_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (maybeSkipControl(t, 'entityOp', 'materia.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set CRISIS_CORE_FUSION_TEST_MATERIA_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let materia_ref01_data = Object.values(setup.data.existing.materia)[0] as any

    // LIST
    const materia_ref01_ent = client.Materia()
    const materia_ref01_match: any = {}

    const materia_ref01_list = await materia_ref01_ent.list(materia_ref01_match)


    // LOAD
    const materia_ref01_match_dt0: any = {}
    materia_ref01_match_dt0.id = materia_ref01_data.id
    const materia_ref01_data_dt0 = await materia_ref01_ent.load(materia_ref01_match_dt0)
    assert(materia_ref01_data_dt0.id === materia_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/materia/MateriaTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = CrisisCoreFusionSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['materia01','materia02','materia03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['CRISIS_CORE_FUSION_TEST_MATERIA_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'CRISIS_CORE_FUSION_TEST_MATERIA_ENTID': idmap,
    'CRISIS_CORE_FUSION_TEST_LIVE': 'FALSE',
    'CRISIS_CORE_FUSION_TEST_EXPLAIN': 'FALSE',
    'CRISIS_CORE_FUSION_APIKEY': 'NONE',
  })

  idmap = env['CRISIS_CORE_FUSION_TEST_MATERIA_ENTID']

  const live = 'TRUE' === env.CRISIS_CORE_FUSION_TEST_LIVE

  if (live) {
    client = new CrisisCoreFusionSDK(merge([
      {
        apikey: env.CRISIS_CORE_FUSION_APIKEY,
      },
      extra
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.CRISIS_CORE_FUSION_TEST_EXPLAIN,
    live,
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  

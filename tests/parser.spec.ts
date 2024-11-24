import { beforeAll, expect, it } from 'vitest'
import { $, fs } from 'zx'
import '../wasm_exec.js'

declare global {
  interface Window {
    __GO_PARSE_FILE__: (content: string) => string
  }
}

beforeAll(async () => {
  await $`tsx scripts/prepare-test.mts`
})

it('can parse generic', async () => {
  const { stdout } = await $`./astexplorer-go tests/examples/generic.go`
  expect(stdout).toContain('SumIntsOrFloats')
  expect(JSON.parse(stdout)).toMatchSnapshot()
})

it('can parse with wasm', async () => {
  const go = new Go()
  const resolvedUrl = (await import('../dist/parser.wasm?url')).default
  const buffer = await fs.readFile('.' + resolvedUrl)
  const { instance } = await WebAssembly.instantiate(buffer, go.importObject)
  go.run(instance)
  const content = await fs.readFile('tests/examples/generic.go', 'utf-8')
  expect(JSON.parse(window.__GO_PARSE_FILE__(content))).toMatchSnapshot()
})

import { beforeAll, expect, it } from 'vitest'
import { $ } from 'zx'

beforeAll(async () => {
  await $`tsx scripts/prepare-test.mts`
})

it('can parse generic', async () => {
  const { stdout } = await $`./astexplorer-go tests/examples/generic.go`
  expect(stdout).toContain('SumIntsOrFloats')
  expect(JSON.parse(stdout)).toMatchSnapshot()
})

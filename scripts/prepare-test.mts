import { fs, $ } from 'zx'

await fs.remove('astexplorer-go')
await $`go build -o astexplorer-go`
await $`yarn run build`

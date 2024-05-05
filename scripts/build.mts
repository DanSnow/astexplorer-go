import { fs, $, glob, path } from 'zx'

await fs.remove('dist')
await fs.mkdir('dist', { recursive: true })
const files = await glob('*.js')
for (const file of files) {
  await fs.copy(file, path.join('dist', file))
}

const pkg = JSON.parse(await fs.readFile('package.json', 'utf8'))
delete pkg.private
delete pkg.devDependencies
delete pkg.scripts
await fs.writeJSON('./dist/package.json', pkg, { spaces: 2 })

await $({ env: { ...process.env, GOOS: 'js', GOARCH: 'wasm' } })`go build -o dist/parser.wasm`

const { readFileSync, writeFileSync } = require('fs')
const { cp, rm, mkdir, exec } = require('shelljs')

rm('-rf', 'dist')
mkdir('-p', 'dist')
cp('*.js', 'dist')
const pkg = JSON.parse(readFileSync('package.json', 'utf8'))
delete pkg.private
delete pkg.devDependencies
delete pkg.scripts
writeFileSync('./dist/package.json', JSON.stringify(pkg, null, 2))
exec('go build -o dist/parser.wasm', { env: { ...process.env, GOOS: 'js', GOARCH: 'wasm' } })

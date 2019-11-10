let _init = async function () {
  const exports = await import('./parser.wasm')
  window._go.run({exports})
}

export async function init() {
  if (_init) {
    await _init()
    _init = undefined
  }
}


export function parseFile (code) {
  return JSON.parse(window.__GO_PARSE_FILE__(code))
}

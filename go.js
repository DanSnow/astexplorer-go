require('./wasm_exec')

window._go = new Go()
window._go.argv = []
window._go.env = []
window._go.exit = () => console.log('EXIT CALLED')
module.exports = window._go.importObject

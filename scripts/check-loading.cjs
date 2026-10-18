const fs = require('node:fs');
const vm = require('node:vm');
const ts = require('typescript');
const assert = require('node:assert/strict');
const calls = [];
const moduleObject = { exports: {} };
const compiled = ts.transpileModule(fs.readFileSync('src/utils/loading.ts', 'utf8'), {
  compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 },
}).outputText;
vm.runInNewContext(compiled, {
  module: moduleObject, exports: moduleObject.exports,
  uni: { showLoading: options => calls.push(['show', options.title, options.mask]), hideLoading: () => calls.push(['hide']) },
});
const { withNativeLoading } = moduleObject.exports;
(async () => {
  let finish;
  const first = withNativeLoading('正在登录', () => new Promise(resolve => { finish = resolve; }));
  await assert.rejects(withNativeLoading('正在保存', async () => { throw Error('failed'); }), /failed/);
  assert.deepEqual(calls, [['show', '正在登录', true]]);
  finish('ok');
  assert.equal(await first, 'ok');
  assert.deepEqual(calls.at(-1), ['hide']);
  await assert.rejects(withNativeLoading('重试', () => { throw Error('sync failure'); }), /sync failure/);
  assert.deepEqual(calls.slice(-2), [['show', '重试', true], ['hide']]);
  console.log('Native loading: concurrent ownership, result forwarding, async/sync failure cleanup passed.');
})().catch(error => { console.error(error); process.exitCode = 1; });

const fs = require('node:fs');
const vm = require('node:vm');
const ts = require('typescript');
const assert = require('node:assert/strict');
const source = fs.readFileSync('src/pages/login/index.vue', 'utf8');
const script = source.match(/<script setup lang="ts">([\s\S]*?)<\/script>/)[1];
const compiled = ts.transpileModule(script + '\nexport { onGetPhoneNumber, busy, error };', {compilerOptions: {module: ts.ModuleKind.CommonJS}}).outputText;
const calls = [], redirects = [];
let fail = false, finish;
const deps = {
  vue: {ref: value => ({value})}, '@dcloudio/uni-app': {onLoad() {}},
  '../../components/SurfaceHeader.vue': {}, '../../theme': {theme: {}},
  '../../utils/auth-navigation': {safeReturnURL: value => value},
  '../../stores/auth': {isAuthenticated: {value: true}, wechatLogin: async code => {calls.push(code); if (fail) throw Error('登录失败'); await new Promise(resolve => {finish = resolve;});}},
};
const moduleObject = {exports: {}};
vm.runInNewContext(compiled, {module: moduleObject, exports: moduleObject.exports, Error, require: name => {assert.ok(name in deps); return deps[name];}, uni: {redirectTo: value => redirects.push(value), reLaunch() {}}});
const page = moduleObject.exports;
(async () => {
  for (const detail of [{errMsg: 'getPhoneNumber:fail user deny'}, {errMsg: 'getPhoneNumber:fail cancel'}, {errMsg: 'getPhoneNumber:ok'}]) {
    await page.onGetPhoneNumber({detail}); assert.equal(calls.length, 0); assert.equal(redirects.length, 0); assert.ok(page.error.value);
  }
  const login = page.onGetPhoneNumber({detail: {errMsg: 'getPhoneNumber:ok', code: 'native-phone-code'}});
  assert.equal(page.busy.value, true);
  await page.onGetPhoneNumber({detail: {errMsg: 'getPhoneNumber:ok', code: 'duplicate'}});
  assert.deepEqual(calls, ['native-phone-code']); finish(); await login;
  assert.equal(page.busy.value, false); assert.equal(redirects.length, 1);
  fail = true;
  await page.onGetPhoneNumber({detail: {errMsg: 'getPhoneNumber:ok', code: 'failed-code'}});
  assert.equal(redirects.length, 1); assert.equal(page.error.value, '登录失败'); assert.equal(page.busy.value, false);
  assert.ok(!source.includes('WeChatProfileSheet'));
  console.log('Phone login: native authorization required, denial/cancel/missing code do not register, duplicate submit blocked, success redirects, failure stays retryable, no login profile popup.');
})().catch(error => {console.error(error); process.exitCode = 1;});

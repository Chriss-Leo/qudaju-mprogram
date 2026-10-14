const fs = require('node:fs');
const vm = require('node:vm');
const ts = require('typescript');
const assert = require('node:assert/strict');
const vue = require('vue');
class UnauthorizedError extends Error {}
const storage = new Map();
const api = { me: async () => user, login: async () => result, logout: async () => {} };
const user = { id: '42', nickname: '微信用户', avatarUrl: '' };
const token = 'a'.repeat(43);
const result = { accessToken: token, tokenType: 'Bearer', expiresIn: 7200, user };
let codeCalls = 0;
function load(file, deps) {
  const compiled = ts.transpileModule(fs.readFileSync(file, 'utf8'), { compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 } }).outputText;
  const module = { exports: {} };
  vm.runInNewContext(compiled, { module, exports: module.exports, require: name => { if (!(name in deps)) throw Error(name); return deps[name]; }, uni: { getStorageSync: k => storage.get(k), setStorageSync: (k,v) => storage.set(k,v), removeStorageSync: k => storage.delete(k), showToast() {} }, Date, Math, console, setTimeout, clearTimeout });
  return module.exports;
}
const auth = load('src/stores/auth.ts', { vue, '../api/identity': { identityAPI: api, UnauthorizedError, validUser: u => !!u && /^\d+$/.test(u.id) && typeof u.nickname === 'string' && typeof u.avatarUrl === 'string' }, '../platforms': { getPlatformAdapter: () => ({ loginCode: async () => { codeCalls++; return 'wx-code'; } }) }, '../config/identity': { requireIdentityURL() {} } });
const key = 'qudaju-auth-session-v1';
const cache = () => storage.set(key, { accessToken: token, expiresAt: Date.now()+7200000, user });
(async () => {
  assert.equal(await auth.restoreAuth(), false);
  assert.throws(() => auth.assertAuthenticated());
  cache(); api.me = async () => { throw new UnauthorizedError(); };
  assert.equal(await auth.restoreAuth(), false); assert.equal(storage.has(key), false);
  cache(); api.me = async () => { throw Error('offline'); };
  assert.equal(await auth.restoreAuth(), false); assert.equal(auth.isAuthenticated.value, false);
  let resolve; api.me = () => new Promise(r => { resolve = r; });
  const pending = auth.restoreAuth(); auth.clearAuth(); resolve(user);
  assert.equal(await pending, false); assert.equal(auth.isAuthenticated.value, false);
  api.me = async () => user;
  await Promise.all([auth.wechatLogin(), auth.wechatLogin()]);
  assert.equal(codeCalls, 1); assert.equal(auth.isAuthenticated.value, true);
  await auth.logout(); assert.equal(auth.isAuthenticated.value, false);
  cache(); assert.equal(await auth.restoreAuth(), true);
  const groups = load('src/mocks/group-cards.ts', {});
  const meetups = load('src/stores/meetups.ts', { vue, '../mocks/group-cards': groups, './auth': auth });
  assert.equal(meetups.meetupState.registrations.length, 0);
  assert.equal(meetups.getGroup('my-game'), undefined);
  meetups.applyToGroup('outdoor-01', 'hello');
  assert.ok(meetups.myRegistration('outdoor-01'));
  auth.authState.session = { accessToken: token, expiresAt: Date.now()+7200000, user: {...user,id:'43'} };
  assert.equal(meetups.meetupState.registrations.length, 0, 'another account must not inherit records');
  auth.clearAuth(); assert.equal(meetups.myRegistration('outdoor-01'), undefined);
  assert.throws(() => meetups.applyToGroup('game-01','guest'));
  console.log('Auth: cache verification, 401/offline rejection, late response, login dedup, logout, account isolation and mutation gates passed.');
})().catch(e => { console.error(e); process.exitCode = 1; }).finally(() => auth.clearAuth());

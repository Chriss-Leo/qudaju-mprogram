const fs = require('node:fs');
const vm = require('node:vm');
const ts = require('typescript');
const assert = require('node:assert/strict');
const vue = require('vue');
const storage = new Map();
// Test-only identity and fixtures; production login cannot be bypassed this way.
const authState = vue.reactive({ session: null });
const isAuthenticated = vue.computed(() => !!authState.session);
const auth = { authState, isAuthenticated, assertAuthenticated() { assert.ok(isAuthenticated.value, 'login required'); } };
function moduleAt(file, dependencies) {
  const source = ts.transpileModule(fs.readFileSync(file, 'utf8'), { compilerOptions: { module: ts.ModuleKind.CommonJS, target: ts.ScriptTarget.ES2020 } }).outputText;
  const module = { exports: {} };
  vm.runInNewContext(source, { exports: module.exports, module, require: name => dependencies[name], uni: { getStorageSync: key => storage.get(key), setStorageSync: (key, value) => storage.set(key, value), showToast() {} }, Date, Math, console });
  return module.exports;
}
function create() {
  const groups = moduleAt('src/mocks/group-cards.ts', {});
  return moduleAt('src/stores/meetups.ts', { vue, '../mocks/group-cards': groups, './auth': auth });
}
(async () => {
  const s = create();
  storage.set('qudaju-local-meetups-v2:fixture', JSON.parse(JSON.stringify(s.meetupState)));
  authState.session = { user: { id: 'fixture', nickname: 'Test', avatarUrl: '' } };
  assert.equal(s.getGroup('__proto__'), undefined);
  assert.equal(s.getGroup('outdoor-01').joined, 6);
  assert.throws(() => s.requestInvitation('game-01', 'sample-1', 'hello'));
  assert.throws(() => s.requestInvitation('my-game', 'me', 'hello'));
  assert.throws(() => s.requestInvitation('my-game', 'not-registered', 'hello'));
  const inviteId = s.requestInvitation('my-game', 'sample-1', 'hello');
  assert.equal(s.requestInvitation('my-game', 'sample-1', 'again'), inviteId);
  assert.throws(() => s.respondInvitation(inviteId, true));
  assert.throws(() => s.sendMessage('invite-me', 'before acceptance'));
  const invitation = s.meetupState.invitations.find(i => i.id === 'invite-me');
  s.respondInvitation('invite-me', true);
  assert.equal(s.canChat(invitation), true);
  s.sendMessage('invite-me', '  hello  ');
  assert.equal(s.meetupState.messages.at(-1).text, 'hello');
  const count = s.meetupState.messages.length;
  s.withdraw('r-joined');
  assert.equal(s.canChat(invitation), false);
  assert.throws(() => s.sendMessage('invite-me', 'after withdrawal'));
  assert.equal(s.meetupState.messages.length, count);
  assert.equal(s.getGroup('outdoor-01').joined, 5);
  s.applyToGroup('outdoor-01', 'reapply');
  assert.equal(s.myRegistration('outdoor-01').status, 'joined');
  assert.equal(s.canChat(invitation), false, 'rejoining must not revive the old conversation');
  assert.throws(() => s.sendMessage('invite-me', 'old conversation after rejoin'));
  s.withdraw('r-joined');
  s.simulatePayment('r-payment');
  assert.equal(s.myRegistration('photo-01').status, 'joined');
  assert.equal(s.getGroup('photo-01').joined, 4);
  assert.throws(() => s.simulatePayment('r-payment'));
  s.reviewMember('host-review', true);
  assert.equal(s.getGroup('my-game').joined, 5);
  assert.throws(() => s.reviewMember('r-review', true));
  s.endGroup('my-game');
  assert.equal(s.canChat(s.meetupState.invitations.find(i => i.id === 'invite-host-accepted')), false);
  assert.throws(() => s.requestInvitation('my-game', 'sample-1', 'closed'));
  await vue.nextTick();
  const restored = create();
  assert.equal(restored.myRegistration('photo-01').status, 'joined');
  assert.equal(restored.myRegistration('outdoor-01').status, 'withdrawn');
  storage.set('qudaju-local-meetups-v2:fixture', { version: 2, registrations: null });
  assert.equal(create().myRegistration('outdoor-01'), undefined);
  console.log('Local rules: ownership, relationship, invitation acceptance, message gating, withdrawal, payment seat count, host review, closure, persistence and corrupt storage fallback passed.');
})().catch(error => { console.error(error); process.exitCode = 1; });

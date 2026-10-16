import { chromium } from '/Users/chris/.cache/codex-runtimes/codex-primary-runtime/dependencies/node/node_modules/playwright/index.mjs';
import assert from 'node:assert/strict';
import fs from 'node:fs';
const base=process.env.QUDADJU_BASE_URL || 'http://127.0.0.1:5186/';
const output='local-tests/profile-20261015';fs.mkdirSync(output,{recursive:true});
const browser=await chromium.launch();
try {
 const errors=[];
 for(const width of [320,393,1200]) {
  const page=await browser.newPage({viewport:{width,height:852}});page.on('pageerror',e=>errors.push(e.message));
  let user={id:'42',nickname:'微信用户',avatarUrl:''};let failSave=true;let uploaded=false;
  await page.route('http://127.0.0.1:8081/v1/auth/me',route=>route.fulfill({json:{code:'1000000',data:user}}));
  await page.route('http://127.0.0.1:8081/v1/auth/me/profile',route=>{ uploaded=route.request().headers()['content-type']?.startsWith('multipart/form-data') || false; if(failSave)return route.fulfill({json:{code:'3010107',data:null}}); user={id:'42',nickname:'小搭子',avatarUrl:'/v1/media/avatars/'+'a'.repeat(48)+'.jpg'};return route.fulfill({json:{code:'1000000',data:user}}); });
  await page.route('http://127.0.0.1:8081/v1/media/avatars/**',route=>route.fulfill({path:'src/static/profile-a/profile-avatar.jpg'}));
  await page.goto(base);await page.waitForFunction(()=>typeof uni !== 'undefined');
  await page.evaluate(()=>uni.setStorageSync('qudaju-auth-session-v1',{accessToken:'a'.repeat(43),expiresAt:Date.now()+7200000,user:{id:'42',nickname:'微信用户',avatarUrl:''}}));
  await page.goto(base+'#/pages/login/index');await page.locator('.login-action').click();await page.getByText('请在微信小程序中使用手机号一键登录。',{exact:true}).waitFor();assert.equal(await page.locator('.profile-form').count(),0,'login must not open custom profile sheet');
  await page.goto(base+'#/pages/profile/index');await page.locator('.edit-profile').click();await page.locator('.profile-form').waitFor();assert.ok(page.url().includes('/pages/profile/index'),'profile editing stays optional in My');
  await page.screenshot({path:`${output}/empty-${width}.png`});
  assert.equal(await page.evaluate(()=>document.documentElement.scrollWidth>innerWidth),false);
  await page.locator('.profile-save').click();await page.locator('.profile-error').waitFor();
  await page.locator('.nickname-field input').fill('小搭子');
  const chooser=page.waitForEvent('filechooser');await page.locator('.avatar-picker').click();await (await chooser).setFiles('src/static/profile-a/profile-avatar.jpg');
  await page.locator('.avatar-picker img').waitFor();
  await page.locator('.profile-save').click();await page.getByText('服务繁忙，请稍后重试',{exact:true}).waitFor();assert.ok(uploaded);
  assert.equal(await page.locator('.nickname-field input').inputValue(),'小搭子');
  await page.waitForFunction(()=>!document.querySelector('.profile-save').hasAttribute('disabled'));await page.screenshot({path:`${output}/retry-${width}.png`});failSave=false;
  await page.locator('.profile-save').click();await page.locator('.identity-card__name').waitFor();await page.getByText('小搭子',{exact:true}).waitFor();
  await page.locator('.edit-profile').click();await page.locator('.profile-form').waitFor();assert.ok(page.url().includes('/pages/profile/index'),'editing should not navigate away');await page.waitForFunction(()=>document.querySelector('.nickname-field input')?.value==='小搭子');assert.equal(await page.locator('.nickname-field input').inputValue(),'小搭子');
  await page.screenshot({path:`${output}/saved-${width}.png`});await page.close();
 }
 assert.deepEqual(errors,[]);console.log('Profile UI: 320/393/1200, anonymous fixtures only in browser, empty validation, upload, failed save retains input, success updates profile, edit prefill passed.');
} finally {await browser.close();}

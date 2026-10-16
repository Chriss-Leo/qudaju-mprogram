-- Additive only: preserve users and their OpenID identity in public.
BEGIN;
ALTER TABLE public.auth_user ADD COLUMN IF NOT EXISTS phone_number VARCHAR(16);
ALTER TABLE public.auth_user ADD COLUMN IF NOT EXISTS phone_verified_time TIMESTAMPTZ;
COMMENT ON COLUMN public.auth_user.phone_number IS '微信服务端验证的E.164手机号；非公开个人资料，不按手机号自动合并账号';
COMMENT ON COLUMN public.auth_user.phone_verified_time IS '最近一次微信手机号验证时间；历史账号可为空';
INSERT INTO public.base_language(code,en,zh) VALUES
 ('wechat_phone_failed','Phone authorization failed; authorize again','手机号验证失败，请重新授权；若持续失败，请检查小程序手机号能力配置')
ON CONFLICT (code) DO NOTHING;
COMMIT;

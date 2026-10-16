-- Preserve existing rows, indexes and the owned identity sequence; never DROP CASCADE.
BEGIN;
SET LOCAL lock_timeout = '5s';
ALTER TABLE IF EXISTS qudaju_identity.auth_user SET SCHEMA public;
ALTER TABLE IF EXISTS qudaju_base.base_language SET SCHEMA public;
DROP SCHEMA IF EXISTS qudaju_identity;
DROP SCHEMA IF EXISTS qudaju_base;
COMMIT;

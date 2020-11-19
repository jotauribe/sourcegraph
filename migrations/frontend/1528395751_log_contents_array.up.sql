BEGIN;

--
-- Drop dependent views

DROP VIEW lsif_indexes_with_repository_name;
DROP VIEW external_service_sync_jobs_with_next_sync_at;

--
-- Swap out column

ALTER TABLE lsif_indexes ADD COLUMN temp json[];
UPDATE lsif_indexes SET temp = (temp || json_build_object('command', '{}'::text[], 'out', log_contents)) WHERE log_contents IS NOT NULL;
ALTER TABLE lsif_indexes DROP COLUMN log_contents;
ALTER TABLE lsif_indexes RENAME COLUMN temp TO log_contents;

ALTER TABLE changesets ADD COLUMN temp json[];
ALTER TABLE changesets DROP COLUMN log_contents;
ALTER TABLE changesets RENAME COLUMN temp TO log_contents;

ALTER TABLE external_service_sync_jobs ADD COLUMN temp json[];
ALTER TABLE external_service_sync_jobs DROP COLUMN log_contents;
ALTER TABLE external_service_sync_jobs RENAME COLUMN temp TO log_contents;

--
-- Recreate views with new columns

CREATE VIEW lsif_indexes_with_repository_name AS
    SELECT u.*, r.name as repository_name FROM lsif_indexes u
    JOIN repo r ON r.id = u.repository_id
    WHERE r.deleted_at IS NULL;

CREATE VIEW external_service_sync_jobs_with_next_sync_at AS SELECT
    j.id,
    j.state,
    j.failure_message,
    j.started_at,
    j.finished_at,
    j.process_after,
    j.num_resets,
    j.num_failures,
    j.log_contents,
    j.external_service_id,
    e.next_sync_at
FROM external_services e JOIN external_service_sync_jobs j ON e.id = j.external_service_id;

COMMIT;

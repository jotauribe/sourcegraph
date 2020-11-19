BEGIN;

--
-- Drop dependent views

DROP VIEW lsif_indexes_with_repository_name;
DROP VIEW external_service_sync_jobs_with_next_sync_at;

--
-- Swap out column

ALTER TABLE lsif_indexes ADD COLUMN temp text;

WITH
    t1 AS (SELECT id, log_contents[1]->>'out' as text FROM lsif_indexes),
    t2 AS (SELECT id, string_agg(t1.text, '\n\n') as text FROM t1 GROUP BY id)
UPDATE lsif_indexes idx SET temp = t2.text FROM t2 WHERE idx.id = t2.id;

ALTER TABLE lsif_indexes DROP COLUMN log_contents;
ALTER TABLE lsif_indexes RENAME COLUMN temp TO log_contents;

ALTER TABLE changesets ADD COLUMN temp text;
ALTER TABLE changesets DROP COLUMN log_contents;
ALTER TABLE changesets RENAME COLUMN temp TO log_contents;

ALTER TABLE external_service_sync_jobs ADD COLUMN temp text;
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

WITH website_parts AS (
    SELECT
        "MY_TABLE".*,
        CASE
            WHEN POSITION('www.' IN SUBSTRING(website FROM POSITION('//' IN website) + 2)) = 1 THEN
                SUBSTRING(website FROM POSITION('//' IN website) + 6)
            ELSE
                SUBSTRING(website FROM POSITION('//' IN website) + 2)
            END AS domain_part
    FROM
        "MY_TABLE"
    WHERE
        website IS NOT NULL AND
            website != '' AND
            POSITION('//' IN website) > 0
)
SELECT
    CASE
        WHEN POSITION('/' IN domain_part) > 0 THEN
            SUBSTRING(domain_part FOR POSITION('/' IN domain_part) - 1)
        ELSE
            domain_part
        END AS domain,
    COUNT(*)
FROM
    website_parts
GROUP BY
    domain
HAVING
        COUNT(*) > 1;

-- less refactored version:
-- SELECT
--     CASE
--         WHEN POSITION('www.' IN SUBSTRING(website FROM POSITION('//' IN website) + 2)) = 1 THEN
--             CASE
--                 WHEN POSITION('/' IN SUBSTRING(website FROM POSITION('//' IN website) + 6)) > 0 THEN
--                     SUBSTRING(website FROM POSITION('//' IN website) + 6 FOR POSITION('/' IN SUBSTRING(website FROM POSITION('//' IN website) + 6)) - 1)
--                 ELSE
--                     SUBSTRING(website FROM POSITION('//' IN website) + 6)
--                 END
--         WHEN POSITION('/' IN SUBSTRING(website FROM POSITION('//' IN website) + 2)) > 0 THEN
--             SUBSTRING(website FROM POSITION('//' IN website) + 2 FOR POSITION('/' IN SUBSTRING(website FROM POSITION('//' IN website) + 2)) - 1)
--         ELSE
--             SUBSTRING(website FROM POSITION('//' IN website) + 2)
--         END AS domain,
--     COUNT(*)
-- FROM
--     "MY_TABLE"
-- WHERE
--     website IS NOT NULL AND
--         website != '' AND
--         POSITION('//' IN website) > 0
-- GROUP BY
--     domain
-- HAVING
--         COUNT(*) > 1;
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
),
     domains AS (
         SELECT
             CASE
                 WHEN POSITION('/' IN domain_part) > 0 THEN
                     SUBSTRING(domain_part FOR POSITION('/' IN domain_part) - 1)
                 ELSE
                     domain_part
                 END AS domain,
             COUNT(*) AS domain_count
         FROM
             website_parts
         GROUP BY
             domain
         HAVING
                 COUNT(*) > 1
     )
SELECT
    website_parts.name,
    domains.domain,
    domains.domain_count
FROM
    website_parts
        JOIN
    domains ON CASE
                   WHEN POSITION('/' IN website_parts.domain_part) > 0 THEN
                       SUBSTRING(website_parts.domain_part FOR POSITION('/' IN website_parts.domain_part) - 1)
                   ELSE
                       website_parts.domain_part
                   END = domains.domain
WHERE
        POSITION(domains.domain IN website_parts.website) > 0;

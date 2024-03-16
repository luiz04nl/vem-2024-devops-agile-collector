SELECT * from repositories;

-- Total 500
SELECT count(*) from repositories;

-- Clonados 490
SELECT count(*) from repositories where wasCloned = 1;

-- Identificados como java - 19
SELECT count(*) from repositories where projectType in ('maven', 'gradle', 'ant')

-- Projetos com componentes do sonar identificados - 17
SELECT count(*) from repositories where projectType in ('maven', 'gradle', 'ant')
and repositories.linesOfCodesFromSonar > 0;

-- Ferramentas de CI/CD
SELECT
    SUM(CASE WHEN useGithubPipelines = 1 THEN 1 ELSE 0 END) AS GithubPipelines,
    SUM(CASE WHEN useCircleCI = 1 THEN 1 ELSE 0 END) AS CircleCI,
    SUM(CASE WHEN useJenkins = 1 THEN 1 ELSE 0 END) AS Jenkins,
    SUM(CASE WHEN useGitLabPipelines = 1 THEN 1 ELSE 0 END) AS GitLabPipelines,
    SUM(CASE WHEN useAzureDevops = 1 THEN 1 ELSE 0 END) AS AzureDevops,
    SUM(CASE WHEN useTravisCI = 1 THEN 1 ELSE 0 END) AS TravisCI,
    SUM(CASE WHEN useHarness = 1 THEN 1 ELSE 0 END) AS Harness,
    SUM(CASE WHEN useBitBucketPipelines = 1 THEN 1 ELSE 0 END) AS BitBucketPipelines
FROM
    repositories
WHERE
    wasCloned = 1
    AND useDevOps = 1;

-- Projetos com Agile/Devops
SELECT
    SUM(CASE WHEN useAgile = 1 THEN 1 ELSE 0 END) AS Agile,
    SUM(CASE WHEN useDevops = 1 THEN 1 ELSE 0 END) AS DevOps,
    SUM(CASE WHEN useDevops = 1 AND useAgile = 1 THEN 1 ELSE 0 END) AS AgileAndDevOps,
    SUM(CASE WHEN useDevops = 0 AND useAgile = 0 THEN 1 ELSE 0 END) AS NeitherAgileAndDevOpsm,
    SUM(CASE WHEN useDevops = 1 AND useAgile = 0 THEN 1 ELSE 0 END) AS DevopsButNonAgile,
    SUM(CASE WHEN useAgile = 1 AND useDevops = 0 THEN 1 ELSE 0 END) AS AgileButNonDovOps
FROM
    repositories
WHERE
    wasCloned = 1;


--######################################### Correlações
-- Projetos considerados nas correlações

 -- Total 17
SELECT count(*) from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')


 -- Projetos com Agile/Devops para correlações
SELECT
    SUM(CASE WHEN useAgile = 1 THEN 1 ELSE 0 END) AS Agile,
    SUM(CASE WHEN useDevops = 1 THEN 1 ELSE 0 END) AS DevOps,
    SUM(CASE WHEN useDevops = 1 AND useAgile = 1 THEN 1 ELSE 0 END) AS AgileAndDevOps,
    SUM(CASE WHEN useDevops = 0 AND useAgile = 0 THEN 1 ELSE 0 END) AS NeitherAgileAndDevOpsm,
    SUM(CASE WHEN useDevops = 1 AND useAgile = 0 THEN 1 ELSE 0 END) AS DevopsButNonAgile,
    SUM(CASE WHEN useAgile = 1 AND useDevops = 0 THEN 1 ELSE 0 END) AS AgileButNonDovOps
FROM
    repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')


 -- Projetos por build type para correlações
SELECT
    projectType AS Name,
    COUNT(id) AS Value
FROM
    repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP BY
    projectType


-- Relação de agile e codeSmells para correlações
select
	useAgile,
	count(useAgile) as useAgileCount,
--	CASE
--    	WHEN CAST(sqaleRating AS REAL) BETWEEN 0 AND 0.05 THEN 'A'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.06 AND 0.1 THEN 'B'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.11 AND 0.2 THEN 'C'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.21 AND 0.5 THEN 'D'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.51 AND 1 THEN 'E'
--        ELSE null
--    END AS maintainabilityRating,
--    AVG(sqaleDebtRatio) AS technicalDebtRatio,
    AVG(codeSmells) AS codeSmellsCount
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
group by useAgile

-- Relação de Devops e codeSmells para correlações
select
	useDevOps,
	count(useDevOps) as useDevOpsCount,
--	CASE
--    	WHEN CAST(sqaleRating AS REAL) BETWEEN 0 AND 0.05 THEN 'A'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.06 AND 0.1 THEN 'B'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.11 AND 0.2 THEN 'C'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.21 AND 0.5 THEN 'D'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.51 AND 1 THEN 'E'
--        ELSE null
--    END AS maintainabilityRating,
--    AVG(sqaleDebtRatio) AS technicalDebtRatio,
    AVG(codeSmells) AS codeSmellsCount
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
group by useDevOps


-- Relação de Agile e Devops e codeSmells para correlações
select
    useDevOps,
    useAgile,
    COUNT(*) AS useDevOpsCount,
--	CASE
--    	WHEN CAST(sqaleRating AS REAL) BETWEEN 0 AND 0.05 THEN 'A'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.06 AND 0.1 THEN 'B'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.11 AND 0.2 THEN 'C'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.21 AND 0.5 THEN 'D'
--        WHEN CAST(sqaleRating AS REAL) BETWEEN 0.51 AND 1 THEN 'E'
--        ELSE null
--    END AS maintainabilityRating,
--    AVG(sqaleDebtRatio) AS technicalDebtRatio,
    AVG(codeSmells) AS codeSmellsCount
from repositories
where projectType in ('maven', 'gradle', 'ant')
and wasCloned = 1 and linesOfCodesFromSonar > 0
GROUP BY useDevOps, useAgile;




  -- spearman
select
	useAgile,
	linesOfCodesFromSonar,
    codeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')


select
	useDevops,
    codeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')


SELECT
--    CASE WHEN useAgile = 1 THEN 1 ELSE 0 END AS Agile,
--    CASE WHEN useDevops = 1 THEN 1 ELSE 0 END AS DevOps,
    CASE WHEN useDevops = 1 AND useAgile = 1 THEN 1 ELSE 0 END AS AgileAndDevOps,
--    CASE WHEN useDevops = 0 AND useAgile = 0 THEN 1 ELSE 0 END AS NeitherAgileAndDevOpsm,
--    CASE WHEN useDevops = 1 AND useAgile = 0 THEN 1 ELSE 0 END AS DevopsButNonAgile,
--    CASE WHEN useAgile = 1 AND useDevops = 0 THEN 1 ELSE 0 END AS AgileButNonDovOps,
    codeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
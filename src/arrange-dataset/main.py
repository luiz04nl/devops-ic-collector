import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def arrange_dataset(consulta_sql, title, descX, descY):
  conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')

  query = consulta_sql
  df = pd.read_sql_query(query, conn)

  conn.close()

  sns.barplot(data=df, x=descX, y=descY)

  plt.title(title)
  plt.xlabel(descX)
  plt.ylabel(descY)

  plt.savefig(f"../../out/arrange-dataset/{descX}-{descY}.png")

consulta_sql_1 = """
SELECT
    (
    	CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Agile Only'
    	ELSE
    		CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Both'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Devops Only'
    			ELSE
    				CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Neither'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) as CategoryGroup,
    count(id) as Quantity
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by CategoryGroup
  """
arrange_dataset(consulta_sql_1, 'Categories', 'CategoryGroup', 'Quantity')

consulta_sql_2 = """
SELECT
    (
    	CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Agile Only'
    	ELSE
    		CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Both'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Devops Only'
    			ELSE
    				CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Neither'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) as CategoryGroup,
    avg(bugs) as Bugs
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by CategoryGroup
  """
arrange_dataset(consulta_sql_2, 'Mean of Bugs by Group', 'CategoryGroup', 'Bugs')



consulta_sql_3 = """
SELECT
    (
    	CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Agile Only'
    	ELSE
    		CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Both'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Devops Only'
    			ELSE
    				CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Neither'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) as CategoryGroup,
    avg(vulnerabilities) as Vulnerabilities
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by CategoryGroup
  """
arrange_dataset(consulta_sql_3, 'Mean of Vulnerabilities by Group', 'CategoryGroup', 'Vulnerabilities')


consulta_sql_4 = """
SELECT
    (
    	CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Agile Only'
    	ELSE
    		CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Both'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Devops Only'
    			ELSE
    				CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Neither'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) as CategoryGroup,
    avg(securityRating) as SecurityRating
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by CategoryGroup
  """
arrange_dataset(consulta_sql_4, 'Mean of Security Rating by Group', 'CategoryGroup', 'SecurityRating')


consulta_sql_5 = """
SELECT
    (
    	CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Agile Only'
    	ELSE
    		CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Both'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Devops Only'
    			ELSE
    				CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Neither'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) as CategoryGroup,
    avg(codeSmells) as CodeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by CategoryGroup
  """
arrange_dataset(consulta_sql_5, 'Mean of Code Smells by Group', 'CategoryGroup', 'CodeSmells')

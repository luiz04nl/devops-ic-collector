import pandas as pd
import statsmodels.api as sm
from statsmodels.formula.api import ols
import sqlite3

def regression(consulta_sql, value):
	title=f"Regression analysis for {value} related with use of Agile methodologies and DevOps"

	conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')
	query = consulta_sql
	df = pd.read_sql_query(query, conn)
	conn.close()

	print(f"### {title} ###")

	modelo = ols(f"{value} ~ CategoryGroup", data=df).fit()

	print(modelo.summary())

	p_values = modelo.pvalues
	print("Valores p (P>|t|):")
	print(p_values)

	alpha = 0.05
	significativos = p_values < alpha
	for coef, p_val, sig in zip(p_values.index, p_values, significativos):
		if sig:
			print(f"Para {value} o coeficiente para '{coef}' é estatisticamente significativo (p = {p_val:.5f}).")
		else:
			print(f"Para {value} o coeficiente para '{coef}' NÃO é estatisticamente significativo (p = {p_val:.5f}).")

	print("###############")

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
    CAST(bugs as INTEGER) as Bugs
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
  """
regression(consulta_sql_2, 'Bugs')


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
    CAST(vulnerabilities as INTEGER) as Vulnerabilities
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')

  """
regression(consulta_sql_3, 'Vulnerabilities')


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
    CAST(securityRating as INTEGER) as SecurityRating
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')

  """
regression(consulta_sql_4, 'SecurityRating')


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
    CAST(codeSmells as INTEGER) as CodeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
  """
regression(consulta_sql_5, 'CodeSmells')

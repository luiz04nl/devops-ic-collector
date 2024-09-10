import pandas as pd
import statsmodels.api as sm
from statsmodels.formula.api import ols
import sqlite3
from statsmodels.stats.multicomp import pairwise_tukeyhsd

def anova(consulta_sql, value):
    title=f"ANOVA results for {value} related with use of Agile methodologies and DevOps"

    conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')
    query = consulta_sql
    df = pd.read_sql_query(query, conn)
    conn.close()

    print(f"### {title} ###")

    try:
        model = ols(f'{value} ~ C(CategoryGroup)', data=df).fit()
        anovaTable = sm.stats.anova_lm(model, typ=2)

        # print(anovaTable)

        # print(f"PR(>F) (Valor p): Probabilidade de observar um valor da estatística F tão extremo quanto o observado, assumindo que a hipótese nula é verdadeira")
        prfValue = anovaTable["PR(>F)"]["C(CategoryGroup)"]
        print(f"PR(>F) = {prfValue}")

        alpha = 0.05
        if prfValue < alpha:
            print(f"Para {value} há uma diferença estatisticamente significativa entre os grupos.")

            # post-hoc
            tukey_result = pairwise_tukeyhsd(endog=df[value], groups=df['CategoryGroup'], alpha=0.05)
            print(tukey_result)

        else:
            print(f"Para {value} não há uma diferença estatisticamente significativa entre os grupos.")

    except Exception as e:
        print(f"Error: {e}")
    
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
anova(consulta_sql_2, 'Bugs')


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
anova(consulta_sql_3, 'Vulnerabilities')


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
anova(consulta_sql_4, 'SecurityRating')


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
anova(consulta_sql_5, 'CodeSmells')

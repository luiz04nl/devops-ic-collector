import pandas as pd
import statsmodels.api as sm
from statsmodels.formula.api import ols
import sqlite3
from statsmodels.stats.multicomp import pairwise_tukeyhsd

def chiSquareTest(consulta_sql_agile, consulta_sql_devops):
    title=f"Chi-Square results for use of Agile methodologies and DevOps"

    conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')

    dfAgile = pd.read_sql_query(consulta_sql_agile, conn)

    dfDevops = pd.read_sql_query(consulta_sql_devops, conn)

    conn.close()

    # print(dfAgile)

    print(f"### {title} ###")

    # Criar um DataFrame com dados sobre o uso das práticas X e Y
    data = {
        'Agile': dfAgile['Agile'],
        'Devops': dfDevops['Devops']
    }
    df = pd.DataFrame(data)

    # Criar a tabela de contingência
    contingency_table = pd.crosstab(df['Agile'], df['Devops'])

    # Exibir a tabela de contingência
    print("Tabela de Contingência:")
    print(contingency_table)

    # Realizar o Teste Qui-Quadrado usando statsmodels
    table = sm.stats.Table(contingency_table)
    chi2_result = table.test_nominal_association()

    # Calcular as frequências esperadas
    expected = table.fittedvalues

    # Exibir os resultados
    print("\nResultado do Teste Qui-Quadrado:")
    print(f"Chi-Quadrado: {chi2_result.statistic}")
    print(f"Valor p: {chi2_result.pvalue}")
    print(f"Graus de Liberdade: {chi2_result.df}")
    print("Frequências Esperadas:")
    print(expected)

    alpha = 0.05
    if chi2_result.pvalue < alpha:
      print("Há uma associação significativa entre Agile e Devops.")
    else:
      print("Não há uma associação significativa entre Agile e Devops.")

    print("###############")

consulta_sql_agile = """
SELECT
    (
    	CASE WHEN useAgile = 1  THEN 'Yes'
    	ELSE
    		'No'
    	END
    ) as Agile
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
  """

consulta_sql_devops = """
SELECT
    (
    	CASE WHEN useDevops = 1  THEN 'Yes'
    	ELSE
    		'No'
    	END
    ) as Devops
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
  """

chiSquareTest(consulta_sql_agile, consulta_sql_devops)

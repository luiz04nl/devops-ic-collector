import pandas as pd
import statsmodels.api as sm
from statsmodels.formula.api import ols
import sqlite3

def contributorsRegression(consulta_sql, value):
	title=f"Relationship between the number of contributors and {value}"

	conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')
	query = consulta_sql
	df = pd.read_sql_query(query, conn)
	conn.close()

	print(df)

	print(f"### {title} ###")

	modelo = ols(f"{value} ~ Contributors", data=df).fit()

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

consulta_sql = """
SELECT
  CAST(projectContributors as INTEGER) as Contributors,
  CAST(codeSmells as INTEGER) as CodeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
  """
contributorsRegression(consulta_sql, 'CodeSmells')

consulta_sql = """
SELECT
  CAST(projectContributors as INTEGER) as Contributors,
  CAST(linesOfCodesFromSonar / codeSmells as INTEGER) as CodeSmellsBylinesOfCode
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
  """
contributorsRegression(consulta_sql, 'CodeSmellsBylinesOfCode')
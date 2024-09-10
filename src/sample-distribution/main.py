import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def sample_distribution(value1, value2, consulta_sql):
  conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')

  query = consulta_sql
  df = pd.read_sql_query(query, conn)

  conn.close()

  sns.scatterplot(x=value1, y=value2, data=df)

  plt.xlabel(value1)
  plt.ylabel(value2)
  plt.title('Sample Code Line Scatterplot (distribution)')

  plt.savefig(f"../../out/sample-distribution/{value1}-{value2}.png")

consulta_sql_3 = """
 SELECT
    id,
    codeSmells
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
sample_distribution('id', 'groups', consulta_sql_3)


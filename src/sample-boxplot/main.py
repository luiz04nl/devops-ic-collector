import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def sample_boxplot(value1, value2, consulta_sql):
  # Conectar ao banco de dados SQLite
  conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')

  # Consultar os dados
  query = consulta_sql
  df = pd.read_sql_query(query, conn)

  # Fechar a conexão com o banco de dados
  conn.close()

  # Criar o boxplot utilizando Seaborn
  plt.figure(figsize=(10, 6))
  sns.boxplot(y=value2, data=df)

  # Adicionar títulos e rótulos
  plt.ylabel(value2)
  plt.title('Boxplot dos ' + value2)

  # Mostrar o gráfico
  # plt.show()
  plt.savefig(f"../../out/sample-boxplot/{value1}-{value2}.png")

consulta_sql_1 = """
 SELECT
    id,
    linesOfCodesFromSonar
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
sample_boxplot('id', 'linesOfCodesFromSonar', consulta_sql_1)

consulta_sql_2 = """
 SELECT
    id,
    codeSmells
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
sample_boxplot('id', 'codeSmells', consulta_sql_2)

consulta_sql_3 = """
 SELECT
    id,
    (linesOfCodesFromSonar / codeSmells) CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
sample_boxplot('id', 'CodeSmellsBylinesOfCodes', consulta_sql_3)


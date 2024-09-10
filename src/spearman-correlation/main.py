import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def spearman_gen(value1, value2, title, consulta_sql):
  # Conectar ao banco de dados
  conexao = sqlite3.connect('../../database/sqlite/repository-dataset.db')
  cursor = conexao.cursor()

  # Executar a consulta
  cursor.execute(consulta_sql)

  # Obter todos os resultados
  dados = cursor.fetchall()

  # Fechar a conexão
  conexao.close()

  # Converter os dados em um DataFrame do Pandas
  df = pd.DataFrame(dados, columns=[value1, value2])

  df[value1] = pd.to_numeric(df[value1], errors='coerce')
  df[value2] = pd.to_numeric(df[value2], errors='coerce')

  # Calcular a correlação de Spearman
  correlacao_spearman = df.corr(method='spearman')

  # Exibir a matriz de correlação
  print("#############################")
  print(correlacao_spearman)

  # Criar o gráfico de dispersão
  plt.figure(figsize=(10, 6))
  sns.scatterplot(x=value1, y=value2, data=df)
  sns.regplot(x=value1, y=value2, data=df, scatter=False, color='red')

  # Adicionar títulos e rótulos
  plt.title(title)
  plt.xlabel(f"Uso de {value1}")
  plt.ylabel(f"{value2}")

  # Mostrar o gráfico
  # plt.show()
  plt.savefig(f"../../out/spearman-correlation/spearman-{value1}-{value2}.png")

# Correlação positiva moderada
consulta_sql_4 = """
  SELECT
    useDevops as UsaDevops,
    useAgile as UsaAgile
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('UsaDevops', 'UsaAgile', 'Correlação do Uso de Devops e Agile', consulta_sql_4)

consulta_sql_1 = """
  SELECT
    useAgile as UsaAgile,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('UsaAgile', 'CodeSmellsBylinesOfCodes', 'Correlação do Uso de Agile e CodeSmells por Linhas de Código', consulta_sql_1)

consulta_sql_2 = """
  SELECT
    useDevops as UsaDevops,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('UsaDevops', 'CodeSmellsBylinesOfCodes', 'Correlação do Uso de Devops e CodeSmells por Linhas de Código', consulta_sql_2)

consulta_sql_3 = """
  SELECT
    CASE WHEN useDevops = 1 AND useAgile = 1 THEN 1 ELSE 0 END AS UsaAgileDevops,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('UsaAgileDevops', 'CodeSmellsBylinesOfCodes', 'Correlação do Uso de Agile Devops e CodeSmells por Linhas de Código', consulta_sql_3)

consulta_sql_5 = """
  SELECT
    projectContributors as Contribuidores,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('Contribuidores', 'CodeSmellsBylinesOfCodes', 'Correlação de Quantidade de Contribuidores e CodeSmells por Linhas de Código', consulta_sql_5)

consulta_sql_6 = """
  SELECT
    projectCommits as Commits,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('Commits', 'CodeSmellsBylinesOfCodes', 'Correlação de Quantidade de Commits e CodeSmells por Linhas de Código', consulta_sql_6)

consulta_sql_6 = """
  SELECT
    commitsIntervalInDays as IntevaloCommits,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('IntevaloCommits', 'CodeSmellsBylinesOfCodes', 'Correlação de Intevalo de Commits e CodeSmells por Linhas de Código', consulta_sql_6)

consulta_sql_6 = """
  SELECT
    linesOfCodesFromSonar as QuantidadeLinhas,
    (linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
spearman_gen('QuantidadeLinhas', 'CodeSmellsBylinesOfCodes', 'Correlação de Quantidade de Linhas escritas e CodeSmells por Linhas de Código', consulta_sql_6)









# 1: Correlação positiva perfeita
# 0.5 a 1: Correlação positiva forte
####### 0.3 a 0.5: Correlação positiva moderada
####### 0 a 0.3: Correlação positiva fraca
# 0: Nenhuma correlação
# -0.3 a 0: Correlação negativa fraca
# -0.5 a -0.3: Correlação negativa moderada
# -1 a -0.5: Correlação negativa forte
# -1: Correlação negativa perfeita

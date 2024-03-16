import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def sample_distribution(value1, value2, consulta_sql):
  # Conectar ao banco de dados SQLite
  conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')

  # Consultar os dados
  query = consulta_sql
  df = pd.read_sql_query(query, conn)

  # Fechar a conexão com o banco de dados
  conn.close()

  # Criar o gráfico de dispersão utilizando Seaborn
  sns.scatterplot(x=value1, y=value2, data=df)

  # Adicionar títulos e rótulos
  plt.xlabel(value1)
  plt.ylabel(value2)
  plt.title('Gráfico de Dispersão de Linhas de Código da amostra')

  # Mostrar o gráfico
  # # plt.show()
  plt.savefig(f"../../out/sample-distribution/{value1}-{value2}.png")

# consulta_sql_1 = """
#  SELECT
#     id,
#     linesOfCodesFromSonar
#   from repositories
#   where wasCloned = 1 and linesOfCodesFromSonar > 0
#   and projectType in ('maven', 'gradle', 'ant')
#   """
# sample_distribution('id', 'linesOfCodesFromSonar', consulta_sql_1)

# consulta_sql_2 = """
#  SELECT
#     id,
#     codeSmells
#   from repositories
#   where wasCloned = 1 and linesOfCodesFromSonar > 0
#   and projectType in ('maven', 'gradle', 'ant')
#   """
# sample_distribution('id', 'codeSmells', consulta_sql_2)

consulta_sql_3 = """
 SELECT
    id,
    codeSmells
  from repositories
  where wasCloned = 1 and linesOfCodesFromSonar > 0
  and projectType in ('maven', 'gradle', 'ant')
  """
sample_distribution('id', 'groups', consulta_sql_3)


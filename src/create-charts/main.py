import sqlite3
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

def create_chart(consulta_sql, title, descX, descY):
  # Conectar ao banco de dados SQLite
  conn = sqlite3.connect('../../database/sqlite/repository-dataset.db')

  # Consultar os dados
  query = consulta_sql
  df = pd.read_sql_query(query, conn)

  # Fechar a conexão com o banco de dados
  conn.close()

  # Criar o gráfico de linha
  # sns.lineplot(data=df, x=descX, y=descY)
  sns.barplot(data=df, x=descX, y=descY)

  # Adicionar títulos e rótulos
  plt.title(title)
  plt.xlabel(descX)
  plt.ylabel(descY)

  # Mostrar o gráfico
  # plt.show()
  plt.savefig(f"../../out/create-chart/{descX}-{descY}.png")

consulta_sql_1 = """
SELECT
    (
    	CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Grupo1'
    	ELSE
    		CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Grupo2'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Grupo3'
    			ELSE
    				CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Grupo4'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) AS Grupo,
    count(id) as Quantidade
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by Grupo
  """
create_chart(consulta_sql_1, 'Classificações', 'Grupo', 'Quantidade')


consulta_sql_2 = """
SELECT
    (
    	CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Grupo1'
    	ELSE
    		CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Grupo2'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Grupo3'
    			ELSE
    				CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Grupo4'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) AS Grupo,
    sum(codeSmells) as CodeSmells
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by Grupo
  """
create_chart(consulta_sql_2, 'CodeSmells por Grupo', 'Grupo', 'CodeSmells')

consulta_sql_3 = """
SELECT
    (
    	CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Grupo1'
    	ELSE
    		CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Grupo2'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Grupo3'
    			ELSE
    				CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Grupo4'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) AS Grupo,
    sum(linesOfCodesFromSonar) as LinesOfCode
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by Grupo
  """
create_chart(consulta_sql_3, 'Linhas de Código por Grupo', 'Grupo', 'LinesOfCode')


consulta_sql_4 = """
SELECT
    (
    	CASE WHEN useDevops = 0 AND useAgile = 0 THEN 'Grupo1'
    	ELSE
    		CASE WHEN useAgile = 1 AND useDevops = 0  THEN 'Grupo2'
    		ELSE
    			CASE WHEN useDevops = 1 AND useAgile = 1 THEN 'Grupo3'
    			ELSE
    				CASE WHEN useDevops = 1 AND useAgile = 0 THEN 'Grupo4'
    				ELSE
    					null
    				END
    			END
    		END
    	END
    ) AS Grupo,
    sum(linesOfCodesFromSonar / codeSmells) as CodeSmellsBylinesOfCodes
from repositories
where wasCloned = 1 and linesOfCodesFromSonar > 0
and projectType in ('maven', 'gradle', 'ant')
GROUP by Grupo
  """
create_chart(consulta_sql_4, 'CodeSmells por Linhas de Código por Grupo', 'Grupo', 'CodeSmellsBylinesOfCodes')
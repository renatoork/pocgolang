#!/bin/bash 

echo "Configurando serviço Pedido" 
export API_PEDIDOVENDA_HOST=127.0.0.1
export API_PEDIDOVENDA_PORT=8080
export API_PEDIDOVENDA_ORACLE="MGVEN/MEGAVEN@10.0.3.161:1521/DESENV"
echo "Configurando serviço AnaliseCredito" 
export API_ANALISECREDITO_PATH='/api/v0'; 
export API_ANALISECREDITO_NOME='credito'; 
export API_ANALISECREDITO_HOST=127.0.0.1; 
export API_ANALISECREDITO_PORT=8081
echo "Executando Serviço" 
./ms-pedidovenda
 



$env:API_PEDIDOVENDA_HOST=127.0.0.1; 
$env:API_PEDIDOVENDA_PORT=8080; 
$env:API_PEDIDOVENDA_ORACLE='MGVEN/MEGAVEN@PC_LEANDROA:1521/DESENV'; 

$env:API_ANALISECREDITO_PATH='/api/v0'; 
$env:API_ANALISECREDITO_NOME='credito'; 
$env:API_ANALISECREDITO_HOST=127.0.0.1; 
$env:API_ANALISECREDITO_PORT=8081

.\ms-pedidovenda.exe -porta 8080
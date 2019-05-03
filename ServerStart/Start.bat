echo Build serviço.
cd \DESENV\git\go\src\mega\ms-consul
go build .
cd \DESENV\git\go\src\mega\ms-analisecredito
go build .
cd \DESENV\git\go\src\mega\ms-pedidovenda
go build .
echo iniciando o Consul...
start \DESENV\git\go\src\mega\ms-consul\ms-consul carregar
echo iniciando o analisecreditogo...
start \DESENV\git\go\src\mega\ms-analisecredito\ms-analisecredito
echo iniciando o analisecreditodelphi...
start \DESENV\git\go\src\mega\ms-analisecreditodelphi\ms-analisecreditodelphi.bat
echo iniciando o pedidovenda...
start \DESENV\git\go\src\mega\ms-pedidovenda\ms-pedidovenda

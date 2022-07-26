# REVAMP-CMI-APPS
 PHP to GO (ERP)

   Library Required
    go get -u github.com/nicholasjackson/env
    
    go get -u github.com/go-sql-driver/mysql [Choose if SQL] 
    
    go get go.mongodb.org/mongo-driver [Choose if Mongo]
    
    go get -u github.com/gin-gonic/gin [API]
    
    go get github.com/joho/godotenv [env]
    

   Database 
    MonggoDB

   Architecture
    Hexagonal ref = https://github.com/matiasvarela/minesweeper-hex-arch-sample/blob/master/internal/handlers/gamehdl/http.go

  [Path]
    root
    _Cmd
    
    _Internal
    
    _ _Core 
    
    _ _ _Domain -> Bisnis Logic -> Entitiy
    
    _ _ _Port/Interface -> Constructor / Interface
    
    _ _ _Service -> Controller
    
    _ _Handler -> Adapter and usually for Json Maker
    
    _ _DataSource/Reporsitories -> Data source just like database
    
    _mocks -> Data Source for Testing
    
    _package -> Error only

  [Step]
  1. Set the domain 
  2. Set the port/interface
  3. Create the service to send out data to endpoint
  4. Create the adapter to get data from Data Source
  5. Create the testing using mocks

  [Use Case]
  1. RetreiveData =  REST API -> Handler [Api (GET)] -> Port Service ->  Domain -> Port Repo -> Adapter -> Data Resource
  
  [Testing]
  - Test per File
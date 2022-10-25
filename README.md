Use Tool

    $ go get -u github.com/ory/dockertest/v3 -> For SQL Test
    $ go get github.com/jmoiron/sqlx [Lib SQL]
    $ go get -u github.com/gin-gonic/gin [API]
    $ go mod init REVAMP-PHP-GO


Testing

    $ go get github.com/smartystreets/goconvey

Run With docker

    $ docker-compose up

Hexagonal Architecture  

The hexagonal architecture, or ports and adapters architecture, is an architectural pattern used in software design. It aims at creating loosely coupled application components that can be easily connected to their software environment by means of ports and adapters. This makes components exchangeable at any level and facilitates test automation

Path

    |── database.
        |── mysql.go
    |── dist.
        |── main.go
    |── docker.
        |── docker-compose.yml
        |── init
            |── db.sql
    |── modules (core).
        |── account (domain module)
            |── domain
                |── domain.go
            |── handler
                |── http.go
            |── service
                |── service.go
        |── modules.go
    |── repo (port database).
        |── repo.go
    |── seed.
    |── go.mod
    |── go.sum
    |── README.md


Features :

    |── Create, Read, Use, Delete
        ── Global
            - Account (DONE)
            - User (DONE)
            - Customer
            - Supplier
            - Estimation
            - Autorizations
            - Items
            - Bank Deposit
            - Bank Account
            - Bank Account Reconciliation
            - Dashboard Report
        ── Sales
            - Sales Order
            - Surat Jalan / Surat Pengiriman Barang
            - Sales Invoice
            - Sales Return
        ── Service
            - Service Order
            - Service Invoice
        ── Reimbursement
            - AP Note (Biaya Sales)
            - Expense Note (Biaya Service)
            - Reimburse Invoice
        ── Kasbon/DP
            - Pending Cash
            - Nota Jaminan
        ── Payment/Receive
            - Payment Voucher
            - Receive Voucher
                Case :  - Cash (Penagihan Cash)
                        - Transfer : Bank Account
                        - Cheque/Giro : Bank Account, Due Date
                        - Settlement  
        ── Purchase
            - Purchase Request
            - Purchase Order
            - Purchase Invoice
            - Receipt Purchase
        ── Asset
            - Asset
            - Asset Depresiasi
            - Asset Rental
            - Asset Purchase
            - Asset Sales
        ── Jurnal
            - GL Journal
            - Ledger
        ── Inventory
            - Good Transafer (Transfer Barang)
            - Inventory Usage
            - Inventory Adjustment